package account

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/liborio7/accountingo/response"
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
	r.Get("/{id}", r.getId)
	r.Get("/", r.get)

	return r
}

func (h *Handler) insert(resp http.ResponseWriter, req *http.Request) {
	body := &Request{}
	if err := render.Bind(req, body); err != nil {
		_ = render.Render(resp, req, response.ErrBadRequest(err))
		return
	}
	model := body.Model
	if err := h.repo.Insert(model); err != nil {
		_ = render.Render(resp, req, response.ErrBadRequest(err))
		return
	}
	render.JSON(resp, req, NewResponse(model))
}

func (h *Handler) getId(resp http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	uid, uidErr := uuid.FromString(id)
	if uidErr != nil {
		_ = render.Render(resp, req, response.ErrNotFound(uidErr))
		return
	}
	model := &Model{Id: uid}
	if err := h.repo.LoadById(model); err != nil {
		_ = render.Render(resp, req, response.ErrNotFound(err))
		return
	}
	render.JSON(resp, req, NewResponse(model))
}

func (h *Handler) get(resp http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	sa := params.Get("sa")
	limit := params.Get("limit")
	parsedSa, err := strconv.ParseUint(sa, 10, 64)
	if err != nil {
		_ = render.Render(resp, req, response.ErrBadRequest(err))
		return
	}
	parsedLimit, err := strconv.ParseUint(limit, 10, 8)
	if err != nil {
		_ = render.Render(resp, req, response.ErrBadRequest(err))
		return
	}
	models, err := h.repo.Load(parsedSa, uint(parsedLimit)+1)
	if err != nil {
		_ = render.Render(resp, req, response.ErrBadRequest(err))
		return
	}
	render.JSON(resp, req, NewPaginatedResponse(models, int(parsedLimit)))
}
