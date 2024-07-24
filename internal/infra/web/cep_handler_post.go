package web

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"renatonasc/multiservice-weathercep/internal/infra/web/webclient"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type CepHandlerPost struct {
	OTELTracer trace.Tracer
}

func NewCepHandlerPost(OTELTracer trace.Tracer) *CepHandlerPost {
	return &CepHandlerPost{
		OTELTracer: OTELTracer,
	}
}

type CepInputDTO struct {
	Cep string `json:"cep"`
}

func (h *CepHandlerPost) GetWeatherByCep(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, span := h.OTELTracer.Start(ctx, "GetWeatherByCep")
	defer span.End()

	var dto CepInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cep := dto.Cep
	if cep == "" {
		http.Error(w, "File cep is required", http.StatusBadRequest)
		return
	}
	log.Println("Service A - CEP informado: ", cep)
	var expReg = regexp.MustCompile(`^\d{5}-?\d{3}$`)
	if !expReg.MatchString(cep) {
		http.Error(w, "CEP deve conter 8 digitos", http.StatusUnprocessableEntity)
		return
	}

	serviceBClient := webclient.NewSeviceBClient()
	weather, err := serviceBClient.GetWeatherByCep(cep, ctx)
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
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)

}
