package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/qdm12/gluetun/internal/publicip"
	"github.com/qdm12/golibs/logging"
)

func newPublicIPHandler(
	looper publicip.Looper,
	logger logging.Logger) http.Handler {
	return &publicIPHandler{
		looper: looper,
		logger: logger,
	}
}

type publicIPHandler struct {
	looper publicip.Looper
	logger logging.Logger
}

func (h *publicIPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = strings.TrimPrefix(r.RequestURI, "/publicip")
	switch r.RequestURI {
	case "/ip":
		switch r.Method {
		case http.MethodGet:
			h.getPublicIP(w)
		default:
			http.Error(w, "", http.StatusNotFound)
		}
	default:
		http.Error(w, "", http.StatusNotFound)
	}
}

func (h *publicIPHandler) getPublicIP(w http.ResponseWriter) {
	data := h.looper.GetData()
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(data); err != nil {
		h.logger.Warn(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
