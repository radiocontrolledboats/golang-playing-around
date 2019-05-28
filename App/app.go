package App

import (
	"../model"
	"net/http"
	"../cache"
	"../handler"
	"sync"
)

type App struct {
	serveMux http.ServeMux
	cache Cache
}

func (a *App) GetCache() Cache {
	return a.cache
}

func init() {
	a.cache = NewCahe()
}

func (a *App) Initialize(db *map[model.ID]*model.Book) {
	a.cache = cache{db: db}
}

func (a *App) getRequestHandler(handler RequestHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.GetCache(), w, r)
	}
}