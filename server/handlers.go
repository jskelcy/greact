package server

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path"
)

// Handlers contains function handlers
type Handlers struct {
	buildPath string
}

// NewHandlers returns router
func NewHandlers(buildPath string) Handlers {
	handlers := Handlers{
		buildPath: buildPath,
	}
	return handlers
}

// HelloWorld returns hello world for debugging
func (h *Handlers) HelloWorld(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	body, err := json.Marshal(map[string]string{
		"data": "Hello, world",
	})
	if err != nil {
		res.WriteHeader(500)
		return
	}
	res.WriteHeader(200)
	res.Write(body)
}

// Home returns tempeled html
func (h *Handlers) Home(res http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("templates", "index.html"))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	data, err := NewViewData(h.buildPath)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	if err := tmpl.Execute(res, data); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
