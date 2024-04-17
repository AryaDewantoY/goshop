package controllers

import (
	"net/http"
	"strconv"

	"github.com/aryadewantoy/goshop/app/models"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func (routes *Server) Products(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout:     "layout",
		Extensions: []string{".html", ".tmpl"},
	})

	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	if page <= 0 {
		page = 1
	}

	perPage := 9

	productModel := models.Product{}
	products, totalRows, err := productModel.GetProduct(routes.DB, perPage, page)
	if err != nil {
		return
	}

	pagination, _ := GetPaginationLinks(routes.AppConfig, PaginationParams{
		Path:        "Products",
		TotalRows:   int32(totalRows),
		PerPage:     int32(perPage),
		CurrentPage: int32(page),
	})

	// fmt.Println("===", pagination)

	_ = render.HTML(w, http.StatusOK, "Products", map[string]interface{}{
		"products":   products,
		"pagination": pagination,
	})
}

func (routes *Server) GetProductBySlug(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout: "layout",
	})

	vars := mux.Vars(r)

	if vars["slug"] == "" {
		return
	}

	productModel := models.Product{}
	product, err := productModel.FindBySlug(routes.DB, vars["slug"])
	if err != nil {
		return
	}

	_ = render.HTML(w, http.StatusOK, "product", map[string]interface{}{
		"product": product,
	})
}
