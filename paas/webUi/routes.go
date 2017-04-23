package webUi

import (
	"net/http"
)

var routes = map[string]func(w http.ResponseWriter,r *http.Request){
	"/index.html": indexHtml,
}
