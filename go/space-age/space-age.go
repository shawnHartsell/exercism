package space

//Planet is the name of a planet
type Planet string

const earthYearInSeconds float64 = 31557600

var planets = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

//Age determines how old someone would be on a particular planet
func Age(secs float64, planet Planet) float64 {
	return (secs / (earthYearInSeconds * planets[planet]))
}
