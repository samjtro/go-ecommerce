package api

import "net/http"

// AddProductToShoppingCartResponse is a response for AddProductToShoppingCartHandler
type GetShoppingCartHandler struct {
	Products []Product `json:"products"`
	Subtotal float64   `json:"price"`
}

func (s *Server) GetShoppingCartHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) AddProductToShoppingCartHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) DeleteProductFromShoppingCartHandler(w http.ResponseWriter, r *http.Request) {

}
