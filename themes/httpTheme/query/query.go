package query

import (
	"net/http"

	"github.com/k0kubun/pp"
)

func TestQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Encode()

	if _, err := w.Write([]byte(query)); err != nil {
		pp.Errorf("Error with write query: ", err)
	}
}
