package account

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/liborio7/accountingo/api"
	"github.com/rs/zerolog/log"
	"github.com/satori/go.uuid"
	"net/http"
	"strconv"
)

type Handler struct {
	chi.Router
	repo *Repo
}

func NewHandler(repo *Repo) *Handler {
	r := &Handler{chi.NewRouter(), repo}

	r.Post("/", r.insert)
	r.Get("/{id}", r.getById)
	r.Get("/", r.get)

	return r
}

func (h *Handler) insert(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	body := &Request{}
	if err := render.Bind(req, body); err != nil {
		_ = render.Render(resp, req, api.ErrBadRequest(err))
		return
	}
	log.Ctx(ctx).Info().Msgf("handle insert with body: %+v", body)
	model := body.Model
	if err := h.repo.Insert(ctx, model); err != nil {
		_ = render.Render(resp, req, api.ErrBadRequest(err))
		return
	}
	render.Status(req, http.StatusOK)
	render.JSON(resp, req, NewResponse(model))
}

func (h *Handler) getById(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	id := chi.URLParam(req, "id")
	uid, err := uuid.FromString(id)
	if err != nil {
		_ = render.Render(resp, req, api.ErrNotFound(err))
		return
	}
	log.Ctx(ctx).Info().Msgf("handle get by id: %+v", id)
	model := &Model{}
	if err := h.repo.LoadById(ctx, model, &uid); err != nil {
		_ = render.Render(resp, req, api.ErrNotFound(err))
		return
	}
	render.Status(req, http.StatusOK)
	render.JSON(resp, req, NewResponse(model))
}

func (h *Handler) get(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	params := req.URL.Query()
	sa, _ := strconv.ParseUint(params.Get("sa"), 10, 0)
	limit, _ := strconv.ParseUint(params.Get("limit"), 10, 0)
	if limit <= 0 || limit > 20 {
		limit = 20
	}
	log.Ctx(ctx).Info().Msgf("handle get starting after %+v limit %+v", sa, limit)
	var models []Model
	if err := h.repo.Load(ctx, &models, &sa, limit+1); err != nil {
		_ = render.Render(resp, req, api.ErrBadRequest(err))
		return
	}
	render.Status(req, http.StatusOK)
	render.JSON(resp, req, NewPaginatedResponse(models, int(limit)))
}

func (h *Handler) update(resp http.ResponseWriter, req *http.Request) {

}
