package entity

type WeaterRespose struct {
	Location              string  `json:"city"`
	Temparatue_celcius    float32 `json:"temp_C"`
	Temparatue_fahrenheit float32 `json:"temp_F"`
	Temperature_kelvin    float32 `json:"temp_K"`
}
