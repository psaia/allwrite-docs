package gdrive

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/LevInteractive/allwrite-docs/model"
	"github.com/LevInteractive/allwrite-docs/util"
	"github.com/russross/blackfriday"
)

const parentDir string = "0B4pmjFk2yyz2NFcwZzQwVHlCRWc"

type titleParts struct {
	Title string
	Order int
}

type pages struct {
	sync.Mutex
	collection model.Pages
	wg         *sync.WaitGroup
}

// Obtain the contents of a google doc by its ID. This essentially pulls the
// nasty html.
func (client *Client) getContents(id string, mimeType string) ([]byte, error) {
	res, err := client.Service.Files.Export(id, mimeType).Download()
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

// Atomically and safely writes to collection when a page is retrieved.
func (s *pages) atomicAppendPage(page *model.Page) {
	s.Lock()
	defer s.Unlock()
	s.collection = append(s.collection, page)
}

// Responsible for splitting apart the allwrite title format.
// e.g. |n| The Title
func getPartsFromTitle(title string) (*titleParts, error) {
	re := regexp.MustCompile("^\\|(\\d+)\\|\\W+?(.+)$")
	result := re.FindStringSubmatch(strings.Trim(title, " "))

	if len(result) == 3 {
		order, err := strconv.Atoi(result[1])
		if err != nil {
			return &titleParts{}, err
		}

		return &titleParts{
			Order: order,
			Title: result[2],
		}, nil
	}
	return &titleParts{}, nil
}

// Determines if the slice already contains a page that isn't a directory with
// the same slug.
func hasLandingPage(collection model.Pages, dir *model.Page) bool {
	hasLanding := false
	for _, page := range collection {
		if page.Type == "file" && page.Slug == dir.Slug {
			hasLanding = true
			break
		}
	}
	return hasLanding
}

// Will take in an existing list and remove any directory pages which have the
// same slug as a "file" page.
func consolidate(collection model.Pages) model.Pages {
	newSlice := make(model.Pages, 0, len(collection))
	for _, page := range collection {
		if page.Type == "dir" {
			if hasLandingPage(collection, page) == false {
				newSlice = append(newSlice, page)
			}
		} else {
			newSlice = append(newSlice, page)
		}
	}
	return newSlice
}

// Query google and walk its directory structure pulling out files.
func (client *Client) processDriveFiles(env *util.Env, baseSlug string, parentID string, pages *pages) {
	defer pages.wg.Done()

	r, err := client.Service.Files.List().
		PageSize(1000). // OK for now.
		Q("'" + parentID + "' in parents").
		Do()

	if err != nil {
		fmt.Printf("Unable to retrieve files from google drive: %v\n", err)
		return
	}

	if len(r.Files) > 0 {
		for _, i := range r.Files {
			// Grab the sort order and title from formatted title names.
			parts, err := getPartsFromTitle(i.Name)
			if err != nil {
				fmt.Printf("Skipping document. There was an issue getting parts from title: %s\n", err.Error())
				continue
			}
			// If the format was incorrect an empty struct will be returned.
			if parts.Title == "" {
				fmt.Printf("Skipping document because of format: %s\n", i.Name)
				continue
			}

			// Define the page that will be saved.
			newPage := &model.Page{}
			newPage.Name = parts.Title
			newPage.DocID = i.Id
			newPage.Order = parts.Order
			newPage.Created = i.CreatedTime
			newPage.Updated = i.ModifiedTime

			// Switch depending on type of ducment.
			switch mime := i.MimeType; mime {
			case "application/vnd.google-apps.document":
				htmlBytes, err := client.getContents(i.Id, "text/html")
				if err != nil {
					fmt.Printf("Skipping. There was an error grabbing the contents for a document: %s", err.Error())
					continue
				}
				md, err := MarshalMarkdownFromHTML(bytes.NewReader(htmlBytes))
				if err != nil {
					fmt.Printf("There was a problem parsing html to markdown: %s", err.Error())
					continue
				}

				newPage.Md = md
				newPage.HTML = string(blackfriday.Run([]byte(md)))
				newPage.Type = "file"

				if parts.Order == 0 {
					// If the order is 0, always take on the same path as the directory.
					newPage.Slug = baseSlug
				} else {
					if baseSlug != "" {
						newPage.Slug = baseSlug + "/" + util.MarshalSlug(parts.Title)
					} else {
						newPage.Slug = baseSlug + util.MarshalSlug(parts.Title)
					}
				}

				fmt.Printf("Saving page \"%s\" with slug \"%s\".\n", newPage.Name, newPage.Slug)
				pages.atomicAppendPage(newPage)

			case "application/vnd.google-apps.folder":
				var dirBaseSlug string

				if baseSlug != "" {
					dirBaseSlug = baseSlug + "/" + util.MarshalSlug(parts.Title)
				} else {
					dirBaseSlug = util.MarshalSlug(parts.Title)
				}
				newPage.Type = "dir"
				newPage.Slug = dirBaseSlug
				fmt.Printf("Saving directory \"%s\" with slug \"%s\".\n", newPage.Name, newPage.Slug)
				pages.atomicAppendPage(newPage)

				pages.wg.Add(1)
				fmt.Printf("Submerging deeper into %s\n", i.Name)
				go client.processDriveFiles(env, dirBaseSlug, i.Id, pages)
			default:
				fmt.Printf("Unknown filetype in drive directory: %s\n", mime)
			}
		}
	} else {
		fmt.Println("No files found.")
	}
}

// UpdateMenu triggers the database to sync with the content.
func UpdateMenu(env *util.Env) error {
	client := DriveClient()
	var wg sync.WaitGroup
	p := &pages{wg: &wg}
	p.wg.Add(1)

	go client.processDriveFiles(
		env,
		"",
		env.CFG.ActiveDir,
		p,
	)

	p.wg.Wait()

	// Loop through and remove any directories that have parent pages.
	p.collection = consolidate(p.collection)

	if err := env.DB.RemoveAll(); err != nil {
		return err
	}
	if _, err := env.DB.SavePages(p.collection); err != nil {
		return err
	}
	return nil
}
