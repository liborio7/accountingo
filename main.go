package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/liborio7/accountingo/account"
	"github.com/liborio7/accountingo/api"
	"github.com/liborio7/accountingo/db"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.PartsOrder = []string{"time", "level", "rid", "caller", "message"}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s >", i)
	}
	log.Logger = zerolog.New(output).With().Timestamp().Caller().Logger()

	dbService := db.Postgres(&db.PostresOpt{
		ConnStr:  "postgres://@localhost:5432/postgres?sslmode=disable",
		MaxConns: 10,
	})
	// cacheService := cache.Redis(&cache.RedisOpt{
	//     Addr:     "localhost:6379",
	//     PoolSize: 10,
	// })

	repo := account.NewRepo(dbService, nil)

	r := chi.NewRouter()
	r.Use(requestId)
	r.Use(tracingRequest)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/users", account.NewHandler(repo))

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Panic().Msgf("error during router server startup: %+v", err)
	}
}

func requestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "rid", fmt.Sprintf("%06d", rand.Intn(999999)))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func tracingRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		rid := ctx.Value("rid")
		logger := log.With().
			Str("rid", rid.(string)).
			Logger()

		req := r.WithContext(logger.WithContext(ctx))
		t1 := time.Now()
		logger.Info().
			Str("method", req.Method).
			Str("uri", req.RequestURI).
			Msg("--- START ---")

		resp := api.NewResponse(w)
		next.ServeHTTP(resp, req)
		logger.Info().
			Str("status", fmt.Sprintf("%d", resp.Status())).
			Str("response_time", time.Since(t1).String()).
			Msg("--- END ---")
	})
}
