package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/liborio7/accountingo/account"
	"github.com/liborio7/accountingo/cache"
	"github.com/liborio7/accountingo/db"
	"log"
	"net/http"
)

func main() {
	db := db.Postgres(&db.PostresOpt{
		ConnStr:  "postgres://@localhost:5432/postgres?sslmode=disable",
		MaxConns: 10,
	})
	cache := cache.Redis(&cache.Opt{
		Addr:     "localhost:6379",
		PoolSize: 10,
	})

	repo := account.NewRepo(db, cache)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/users", account.NewHandler(repo))

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Println("error during router server startup:", err)
		panic("application stopped")
	}
}
