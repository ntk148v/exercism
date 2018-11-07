/*Package space calculate how old someone would be in age*/
package space

// Planet is the location where age will be calculated
type Planet string

// Possible planets
const (
	PlanetEarth        Planet = "Earth"
	PlanetJupiter      Planet = "Jupiter"
	PlanetMars         Planet = "Mars"
	PlanetMercury      Planet = "Mercury"
	PlanetNeptune      Planet = "Neptune"
	PlanetSaturn       Planet = "Saturn"
	PlanetUranus       Planet = "Uranus"
	PlanetVenus        Planet = "Venus"
	yearInSecondsEarth        = 31557600.0
)

var orbitalPeriods = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	yearInSecondsPlanet := yearInSecondsEarth * orbitalPeriods[planet]
	years := seconds / yearInSecondsPlanet
	return years
}
