package gdrive

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/LevInteractive/allwrite-docs/model"
	"github.com/LevInteractive/allwrite-docs/util"
	drive "google.golang.org/api/drive/v3"
)

const parentDir string = "0B4pmjFk2yyz2NFcwZzQwVHlCRWc"

type titleParts struct {
	Title string
	Order int64
}

func getContents(service *drive.Service, id string, mimeType string) ([]byte, error) {
	res, err := service.Files.Export(id, mimeType).Download()
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func getPartsFromTitle(title string) (*titleParts, error) {
	re := regexp.MustCompile("^\\|(\\d+)\\|\\W+?(.+)$")
	result := re.FindStringSubmatch(strings.Trim(title, " "))

	if len(result) == 3 {
		order, err := strconv.ParseInt(result[1], 10, 64)
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

func getFiles(service *drive.Service, baseSlug string, parentID string) []model.Page {
	r, err := service.Files.List().
		PageSize(1000).
		Q("'" + parentID + "' in parents").
		Do()

	if err != nil {
		log.Fatalf("Unable to retrieve files from google drive: %v\n", err)
	}

	var pages []model.Page

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

			// Switch depending on type of ducment.
			switch mime := i.MimeType; mime {
			case "application/vnd.google-apps.document":
				htmlBytes, err := getContents(service, i.Id, "text/html")
				if err != nil {
					fmt.Printf("Skipping. There was an error grabbing the contents for a document: %s", err.Error())
					continue
				}

				htmlStr := string(htmlBytes)
				fmt.Println(htmlStr)

				newPage := model.Page{
					Name:    parts.Title,
					Created: i.CreatedTime,
					Md:      htmlStr,
					HTML:    htmlStr,
					Updated: i.ModifiedTime,
				}

				if parts.Order == 0 {
					newPage.Slug = baseSlug
				} else {
					newPage.Slug = baseSlug + "/" + util.MarshalSlug(parts.Title)
				}

				fmt.Printf("Saving page \"%s\" with slug \"%s\".\n", newPage.Name, newPage.Slug)
				pages = append(pages, newPage)

			case "application/vnd.google-apps.folder":
				var dirBaseSlug string

				if baseSlug != "/" {
					dirBaseSlug = "/" + baseSlug + "/" + util.MarshalSlug(parts.Title)
				} else {
					dirBaseSlug = util.MarshalSlug(parts.Title)
				}
				fmt.Printf("Submerging deeper into %s\n", i.Name)
				getFiles(service, dirBaseSlug, i.Id)
				break
			default:
				fmt.Printf("Unknown filetype in drive directory: %s\n", mime)
			}
		}
	} else {
		fmt.Println("No files found.")
	}
	return pages
}

// UpdateMenu is
func UpdateMenu(cfg *util.Conf, service *drive.Service) {
	_ = getFiles(
		service,
		"/",
		cfg.ActiveDir,
	)
}
