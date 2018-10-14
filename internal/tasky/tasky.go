package tasky

import (
	"net/http"

	asecurity "github.com/PGo-Projects/authboss-security"
	"github.com/PGo-Projects/tplmgr"
)

var ProtectedRoutes = []asecurity.ProtectedRoute{
	asecurity.ProtectedRoute{http.MethodGet, "/tasky", taskyPage},
}

func taskyPage(w http.ResponseWriter, r *http.Request) {
	tplmgr.AuthbossSAHTMLRenderer(w, r, "tasky", ".tmpl", tplmgr.HTMLData{})
}
