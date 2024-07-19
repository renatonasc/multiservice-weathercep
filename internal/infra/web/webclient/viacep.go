package webclient

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type ViaCepDTO struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Location   string `json:"localidade"`
	Error      string `json:"erro"`
}

type ViaCepClient struct {
}

func NewViaCepClient() *ViaCepClient {
	return &ViaCepClient{}
}

func (v *ViaCepClient) GetLocationByCep(cep string) (string, error) {
	url := "https://viacep.com.br/ws/" + cep + "/json/"
	context := context.Background()
	req, err := http.NewRequestWithContext(context, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var viaCepDTO ViaCepDTO
	err = json.Unmarshal(body, &viaCepDTO)
	if err != nil {
		return "", err
	}
	if viaCepDTO.Error == "true" {
		return "", errors.New("CEP não encontrado")
	}
	return viaCepDTO.Location, nil

}
