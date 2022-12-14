package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

const idQueryKey = "id"

// Root is a struct for root handler
type Root struct{}

// ServeHTTP is a root handler
func (r *Root) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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

// SnippetGetter is a struct for snippet getter
type SnippetGetter struct{}

// ServeHTTP is a snippet getter by id
func (sg *SnippetGetter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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

// SnippetAdder is a struct for adding snippets
type SnippetAdder struct{}

// AddSnippet adds a new snippet
func (ss *SnippetAdder) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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
