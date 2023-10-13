package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Weather struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {

	rand.Seed(time.Now().Unix())

	for {
		time.Sleep(2 * time.Second)

		RandomvalueforWater := rand.Intn(100)
		RandomValueForWind := rand.Intn(100)

		var weather Weather = Weather{
			Water: RandomvalueforWater,
			Wind:  RandomValueForWind,
		}

		JSONmarshal, err := json.Marshal(weather)
		if err != nil {
			fmt.Println("gagal membuat jadi json")
		} else {
			fmt.Println(string(JSONmarshal))
		}

		if RandomvalueforWater < 5 {
			fmt.Println("status water: aman")
		} else if RandomvalueforWater >= 6 && RandomvalueforWater <= 8 {
			fmt.Println("status water: siaga")
		} else {
			fmt.Println("status water: bahaya")
		}

		if RandomValueForWind < 6 {
			fmt.Println("status wind: aman")
		} else if RandomValueForWind >= 7 && RandomValueForWind <= 15 {
			fmt.Println("status wind: siaga")
		} else {
			fmt.Println("status wind: bahaya")
		}
	}
}
