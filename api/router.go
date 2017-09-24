package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/LevInteractive/allwrite-docs/model"
	"github.com/LevInteractive/allwrite-docs/util"
)

type jsonResponse struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Error  string      `json:"error"`
}

// SortPages sorts the pages in the proper heirarchy.
func SortPages(pages []*model.Page) []*model.Page {
	return pages
}

func getPage(env *util.Env, uri string, w http.ResponseWriter, req *http.Request) {
	page, err := env.DB.GetPage(uri)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&jsonResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(jsonResponse{
			Code:   http.StatusOK,
			Result: page,
		})
	}
}

// func getMenu(uri string, w http.ResponseWriter, req *http.Request) {
// 	rows, err := db.Query("SELECT name FROM users WHERE age=?", age)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var name string
// 		if err := rows.Scan(&name); err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Printf("%s is %d\n", name, age)
// 	}
// 	if err := rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// Listen to incoming requests and serve.
func Listen(env *util.Env) {
	stripSlashes := regexp.MustCompile("^\\/|\\/$")
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		uri := stripSlashes.ReplaceAllString(req.RequestURI, "")
		if uri == "/menu" {
			// getMenu(uri, w, req)
		} else {
			getPage(env, uri, w, req)
		}
	})

	fmt.Printf("\nListening on %s\n", env.CFG.Port)
	log.Fatal(http.ListenAndServe(env.CFG.Port, nil))
}
