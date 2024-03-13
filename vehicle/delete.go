package vehicle

import (
	"net/http"
	"strconv"

	"github.com/rom2049/vehicle-server/storage"
	"go.uber.org/zap"
)

type DeleteHandler struct {
	store  storage.Store
	logger *zap.Logger
}

func NewDeleteHandler(store storage.Store, logger *zap.Logger) *DeleteHandler {
	return &DeleteHandler{
		store:  store,
		logger: logger.With(zap.String("handler", "delete_vehicles")),
	}
}

func (d *DeleteHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//http.Error(rw, "Not Implemented", http.StatusInternalServerError)
	id := r.PathValue("id")
	i, _ := strconv.ParseInt(id, 10, 64)
	cont := r.Context()
	rep, _ := d.store.Vehicle().Delete(cont, i)
	if rep {
		rw.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(rw, "404", http.StatusNotFound)
	}
}
