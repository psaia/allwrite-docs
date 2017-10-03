package api

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/LevInteractive/allwrite-docs/util"
)

type jsonResponse struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Error  string      `json:"error,omitempty"`
}

func getPage(env *util.Env, uri string, w http.ResponseWriter, req *http.Request) {
	page, err := env.DB.GetPage(uri)
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

func getMenu(env *util.Env, uri string, w http.ResponseWriter, req *http.Request) {
	menu, err := env.DB.GetMenu()
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
			Result: menu,
		})
	}
}

func search(env *util.Env, search string, uri string, w http.ResponseWriter, req *http.Request) {
	menu, err := env.DB.Search(search)
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
			Result: menu,
		})
	}
}

// Listen to incoming requests and serve.
func Listen(env *util.Env) {
	stripSlashes := regexp.MustCompile("^\\/|\\/$|\\?.*$")
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Cache-Control", "public, max-age=3600")
		s := req.URL.Query().Get("q")
		uri := stripSlashes.ReplaceAllString(req.RequestURI, "")
		if uri == "menu" {
			log.Printf("Request: Menu\n")
			getMenu(env, uri, w, req)
		} else if len(s) > 0 && uri == "" {
			log.Printf("Request: Search - %s\n", s)
			search(env, s, uri, w, req)
		} else {
			log.Printf("Request: Get Page - %s\n", uri)
			getPage(env, uri, w, req)
		}
	})

	log.Printf("\nListening on %s\n", env.CFG.Port)
	log.Fatal(http.ListenAndServe(env.CFG.Port, nil))
}
