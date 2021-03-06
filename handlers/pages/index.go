package pages

import (
	"net/http"

	"github.com/vision8tech/goings/common"
	"github.com/vision8tech/goings/shared/templatedata"
	"go.isomorphicgo.org/go/isokit"
)

// IndexPageHandler is the handler for the GET "/" requests.
func IndexPageHandler(env *common.Env) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templateData := templatedata.Index{PageTitle: "Goings"}
		env.TemplateSet.Render("index_page", &isokit.RenderParams{Writer: w, Data: templateData})
	})

}
