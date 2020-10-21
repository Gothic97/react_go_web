package main

import "fmt"

func main() {

	// var a map[string]int // var map_name map[key_type]value_type

	// var a map[string]int = make(map[string]int)
	// var b = make(map[string]int)
	// c := make(map[string]int)

	// a := map[string]int{"Hello": 10, "world": 20}
	//
	// b := map[string]int{
	// 	"Hello": 10,
	// 	"world": 20, // when we enumerate variables as use '{}', we must affix the ',' in last.
	// }
	//
	// fmt.Println(a["Hello"])
	// fmt.Println(b["world"])

	solarSystem := make(map[string]float32)

	solarSystem["Mercury"] = 87.969
	solarSystem["Venus"] = 224.70069
	solarSystem["Earth"] = 365.25641
	solarSystem["Mars"] = 686.9600
	solarSystem["Jupiter"] = 4333.2867
	solarSystem["Saturn"] = 10756.1995
	solarSystem["Uranus"] = 30707.4896
	solarSystem["Neptune"] = 60223.3528
	//
	// fmt.Println(solarSystem["Earth"])
	// fmt.Println(solarSystem["Pluto"]) // if we make an inquiry about a value that does'nt exist, print empty value not nil.
	//
	// value, ok := solarSystem["Pluto"] // if you want to know about whether value is empty or not, you can appoint two variables. the first one is value, and second is bool type about whether it is or not.
	// fmt.Println(value, ok)
	//
	// if value, ok := solarSystem["Saturn"]; ok {
	// 	fmt.Println(value) // 10756.1995
	// }

	// for 'key', 'value' := range 'map' {}
	// for key, value := range solarSystem { // when the repetitive statement execute, key and value go into variables automatically.
	//
	// 	fmt.Println(key, value)
	//
	// }

	// a := map[string]int{"Hello": 10, "world": 20}
	//
	// delete(a, "world") // delete "world" key in map.
	//
	// fmt.Println(a)

	// map[key_type]map[key_type]value_type

	terrestrialPlanet := map[string]map[string]float32{

		"Mercury": map[string]float32{
			"meanRadius":    2439.7,
			"mass":          3.3022E+23,
			"orbitalPeriod": 87.969,
		},
		"Venus": map[string]float32{
			"meanRadius":    6051.8,
			"mass":          4.8676E+24,
			"orbitalPeriod": 22.70069,
		},
		"Earth": map[string]float32{
			"meanRadius":    6371.0,
			"mass":          5.97219E+24,
			"orbitalPeriod": 365.25641,
		},
		"Mars": map[string]float32{
			"meanRadius":    3389.5,
			"mass":          6.4185E+23,
			"orbitalPeriod": 686.9600,
		},
	}

	fmt.Println(terrestrialPlanet["Mars"]["mass"])
	// if you affix map keyword at value_type of map keyword, you can create map in the map. if you want to receive juniority item, you must designate keys consecutively.

}
