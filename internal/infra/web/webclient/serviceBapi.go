package webclient

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"renatonasc/multiservice-weathercep/internal/entity"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type SeviceBClient struct {
}

func NewSeviceBClient() *SeviceBClient {
	return &SeviceBClient{}
}

func (v *SeviceBClient) GetWeatherByCep(cep string, ctx context.Context) (*entity.WeaterRespose, error) {
	url := "http://serviceB-app:8080/weather/" + cep
	context := context.Background()
	req, err := http.NewRequestWithContext(context, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weaterRespose entity.WeaterRespose
	err = json.Unmarshal(body, &weaterRespose)
	if err != nil {
		return nil, err
	}

	return &weaterRespose, nil
}
