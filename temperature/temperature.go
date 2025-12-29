package temperature

import "fmt"

type Temperature struct {
	celsius float64
}

func NewCelsius(c float64) Temperature {

	if c < -273.15 {
		c = -273.15
	}

	return Temperature{celsius: c}
}

func NewFahrenheit(f float64) Temperature {
	if f < -459.67 {
		f = -459.67
	}

	c := (f - 32) * 5 / 9
	return Temperature{celsius: c}
}

func NewKelvin(k float64) Temperature {
	if k < 0 {
		k = 0
	}

	c := k - 273.15
	return Temperature{celsius: c}
}

func (t Temperature) Celsius() float64 {
	return t.celsius
}

func (t Temperature) Fahrenheit() float64 {
	return t.celsius*9/5 + 32
}

func (t Temperature) Kelvin() float64 {
	return t.celsius + 273.15
}

func (t Temperature) String() string {
	return fmt.Sprintf("%.1f°C", t.celsius)
}
