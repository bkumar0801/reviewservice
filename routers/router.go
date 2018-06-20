package routers

import (
	"github.com/gorilla/mux"
	"github.com/reviewservice/handlers"
)

/*
Init ...
*/
func Init() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/health", handlers.HandleHealthCheck).Methods("GET")
	router.HandleFunc("/apps/{appId}/products/{id}/reviews", handlers.HandleGetReview).Methods("GET")
	router.HandleFunc("/apps/{appId}/products/{id}/reviews", handlers.HandleCreateReview).Methods("POST")
	router.HandleFunc("/apps/{appId}/products/{id}", handlers.HandleDeleteReview).Methods("DELETE")
	return router
}
