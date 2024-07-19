package usecase

import (
	"context"
	"log"
	"renatonasc/multiservice-weathercep/internal/entity"

	"go.opentelemetry.io/otel/trace"
)

type GetWeatherByCepUseCase struct {
	Ctx        context.Context
	OTELTracer trace.Tracer
}

func (u *GetWeatherByCepUseCase) Execute(cep string) (*entity.WeaterRespose, error) {

	_, span := u.OTELTracer.Start(u.Ctx, "GetWeatherByCepUseCase  - GetLocationByCep")
	getLocationUseCase := GetLocationByCepUseCase{}
	location, err := getLocationUseCase.Execute(cep)
	log.Println("Service B - Location: ", location)
	log.Println("Service B - Error: ", err)
	if err != nil {
		return nil, err
	}
	span.End()

	_, span = u.OTELTracer.Start(u.Ctx, "GetWeatherByCepUseCase  - GetWeatherByLocation")
	getWeatherByLocationUseCase := GetWeatherByLocationUsecase{}
	weather, err := getWeatherByLocationUseCase.Execute(location)
	if err != nil {
		return nil, err
	}
	span.End()

	return weather, nil
}
