package web

import (
	"encoding/json"
	"log"
	"net/http"
	"renatonasc/multiservice-weathercep/internal/usecase"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type Error struct {
	Message string `json:"message"`
}

type CepHandler struct {
	OTELTracer trace.Tracer
}

func NewCepHandler(OTELTracer trace.Tracer) *CepHandler {
	return &CepHandler{
		OTELTracer: OTELTracer,
	}
}

func (h *CepHandler) GetWeatherByCep(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	_, span := h.OTELTracer.Start(ctx, "GetWeatherByCep")
	defer span.End()

	cep := chi.URLParam(r, "cep")
	if cep == "" {
		http.Error(w, "File cep is required", http.StatusBadRequest)
		return
	}

	log.Println("Service B - CEP informado: ", cep)
	getWeatherByCepUsecase := usecase.GetWeatherByCepUseCase{
		Ctx:        ctx,
		OTELTracer: h.OTELTracer,
	}
	weather, err := getWeatherByCepUsecase.Execute(cep)
	if err != nil {
		if err.Error() == "CEP deve conter 8 digitos" {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		if err.Error() == "CEP não encontrado" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
