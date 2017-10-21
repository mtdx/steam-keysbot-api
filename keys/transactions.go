package keys

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/mtdx/keyc/common"
)

func (rd *TransactionsResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// TransactionsHandler rest route handler
func TransactionsHandler(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	dbconn := r.Context().Value("DBCONN").(*sql.DB)
	transactionsresp, err := findAllTransactions(dbconn, claims["id"])
	if err != nil {
		render.Render(w, r, common.ErrInternalServer(err))
		return
	}
	common.RenderResults(w, r, transactionsresp, err)
}