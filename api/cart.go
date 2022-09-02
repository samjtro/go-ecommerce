package api

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
)

// AddProductToShoppingCartResponse is a response for AddProductToShoppingCartHandler
type GetShoppingCartHandler struct {
	Products []Product `json:"products"`
	Subtotal float64   `json:"price"`
}

func (s *Server) GetShoppingCartHandler(w http.ResponseWriter, r *http.Request) {
	products := ReturnProductSlice(r)
	subtotal, err := strconv.ParseFloat(chi.URLParam(r, "subtotal"), 64)

	if err != nil {
		log.Fatal(err)
	}

	cart := NewShoppingCart(products, subtotal)

	rnd.JSON(w, http.StatusOK, map[string]interface{}{
		"products": cart,
		"subtotal": subtotal,
	})
}

func (s *Server) AddProductToShoppingCartHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) DeleteProductFromShoppingCartHandler(w http.ResponseWriter, r *http.Request) {

}

func ReturnProductSlice(r *http.Request) []Product {
	var products []Product
	unformattedProducts := chi.URLParam(r, "products")

	// need to figure this out, how will cart get passed and reformatted?
	for _, x := range strings.Split(unformattedProducts, "") {
		price, err := strconv.ParseFloat(strings.Split(x, ",")[1], 64)

		if err != nil {
			log.Fatal(err)
		}

		product := NewProduct(strings.Split(x, ",")[0], price)
		products = append(products, product)
	}

	return products
}
