package app

import (
	"github.com/go-chi/chi/v5"
	"gocloud/handlers"
	"gocloud/services/auth"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Post("/login", auth.Authorize)
	r.Use(auth.Authenticate)

	r.Get("/create", handlers.CreateStorage)
	r.Get("/read", handlers.Read)
	r.Post("/write", handlers.Write)
	r.Put("/update", handlers.Update)
	r.Delete("/delete", handlers.Delete)

	http.ListenAndServe(":8080", r)
}
