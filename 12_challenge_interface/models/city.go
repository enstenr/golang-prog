package models

var(
	beachVacationThreshold  float64=22
	skiVacationThreshold float64=-2
)

type city struct {

	name string
	tempC float64
}
type CityTemp interface{
	Name() string
	Tempc() float64
	TempF() float64

	BeachVacationReady() bool
	SkiVacationReady() bool
}
func NewCity(n string,t float64,hasBeach bool , hasMountain bool) CityTemp{
	return &city{
		name:n,
		tempC:t,
	}
}

func (c city) Name() string{
	return c.name
}
func (c city) Tempc() float64{
	return c.tempC
}
func (c city) TempF() float64{
	return (c.tempC*9/5)+32
}

func (c city)BeachVacationReady()bool{
	if c.Tempc()>70{
		return true
	}
	return false
}
func (c city)SkiVacationReady()bool{
	if (c.TempF()< -2) {
		return true
	}
	return false
}