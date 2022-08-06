package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

const idQueryKey = "id"

// Root is a root handler
func Root(w http.ResponseWriter, req *http.Request) {
	rootFiles := []string{
		"ui/html/root.page.html",
		"ui/html/base.layout.html",
		"ui/html/footer.partial.html",
	}

	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	ts, err := template.ParseFiles(rootFiles...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// GetSnippet is a snippet getter by id
func GetSnippet(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(req.URL.Query().Get(idQueryKey), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, req)
		log.Println(err)
		return
	}

	_, err = fmt.Fprintf(w, "This is a snippet with id: %d", id)
	if err != nil {
		log.Println(err)
		return
	}
}

// AddSnippet adds a new snippet
func AddSnippet(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.Header().Add("Allow", http.MethodPost)
		http.Error(w, "Method "+req.Method+" is not allowed", http.StatusMethodNotAllowed)
		return
	}
	_, err := w.Write([]byte("This is a snippet creator"))
	if err != nil {
		log.Println(err)
		return
	}
}
