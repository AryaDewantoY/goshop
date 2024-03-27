package controllers

import (
	"net/http"

	"github.com/aryadewantoy/goshop/app/models"
	"github.com/unrolled/render"
)

func (routes *Server) Products(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout: "layout",
		Extensions: []string{".html", ".tmpl"},
	})

	productModel := models.Product{}
	products, err := productModel.GetProduct(routes.DB)
	if err !=  nil {
		return
	}

	_ = render.HTML(w, http.StatusOK, "Products", map[string]interface{}{
		"products":  products,
	})
}