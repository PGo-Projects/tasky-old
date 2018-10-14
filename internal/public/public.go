package public

import (
	"net/http"

	"github.com/PGo-Projects/tplmgr"
	"github.com/go-chi/chi"
)

func RegisterRoutes(mux chi.Router) {
	mux.Get("/", indexPage)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	tplmgr.AuthbossSAHTMLRenderer(w, r, "index", ".tmpl", tplmgr.HTMLData{})
}
