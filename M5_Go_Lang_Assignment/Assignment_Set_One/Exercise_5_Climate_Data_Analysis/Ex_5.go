/*
Exercise 5: Climate Data Analysis

Topics Covered: Go Arrays, Go Strings, Go Type Casting, Go Functions, Go Conditions, Go Loops

Case Study:

You are tasked with analyzing climate data from multiple cities. The data includes the city name, average temperature (°C), and rainfall (mm).

1. Data Input: Create a slice of structs to store data for each city. Input data can be hardcoded or taken from the user.

2. Highest and Lowest Temperature: Write functions to find the city with the highest and lowest average temperatures. Use conditions for comparison.

3. Average Rainfall: Calculate the average rainfall across all cities using loops. Use type casting if necessary.

4. Filter Cities by Rainfall: Use loops to display cities with rainfall above a certain threshold. Prompt the user to enter the threshold value.

5. Search by City Name: Allow users to search for a city by name and display its data.

Bonus:

• Add error handling for invalid city names and invalid input for thresholds.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ClimateData struct {
	City    string
	AvgTemp float64
	AvgRain float64
}

func displayMenu() {
	fmt.Println("\n=== Climate Data Analysis Menu ===")
	fmt.Println("1. Display all cities")
	fmt.Println("2. Show highest and lowest temperatures")
	fmt.Println("3. Calculate average rainfall")
	fmt.Println("4. Filter cities by rainfall threshold")
	fmt.Println("5. Search for a city")
	fmt.Println("6. Exit")
	fmt.Print("\nEnter your choice (1-6): ")
}

func displayAllCities(cities []ClimateData) {
	fmt.Println("\n=== All Cities Data ===")
	for _, city := range cities {
		fmt.Printf("City: %s\nTemperature: %.2f °C\nRainfall: %.2f mm\n\n",
			city.City, city.AvgTemp, city.AvgRain)
	}
}

func findTemperatureExtremes(cities []ClimateData) (highestCity ClimateData, lowestCity ClimateData) {
	highestCity = cities[0]
	lowestCity = cities[0]

	for _, city := range cities {
		if city.AvgTemp > highestCity.AvgTemp {
			highestCity = city
		}
		if city.AvgTemp < lowestCity.AvgTemp {
			lowestCity = city
		}
	}

	return highestCity, lowestCity
}

func calculateAverageRainfall(cities []ClimateData) float64 {
	if len(cities) == 0 {
		return 0
	}

	var totalRainfall float64
	for _, city := range cities {
		totalRainfall += city.AvgRain
	}

	return totalRainfall / float64(len(cities))
}

func filterCitiesByRainfall(cities []ClimateData, threshold float64) []ClimateData {
	var filteredCities []ClimateData
	for _, city := range cities {
		if city.AvgRain > threshold {
			filteredCities = append(filteredCities, city)
		}
	}
	return filteredCities
}

func promptThreshold() (float64, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nEnter the rainfall threshold (mm): ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("error reading input: %v", err)
	}

	text = strings.TrimSpace(text)
	threshold, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid threshold value: %v", err)
	}

	if threshold < 0 {
		return 0, fmt.Errorf("threshold cannot be negative")
	}

	return threshold, nil
}

func searchCity(cities []ClimateData, name string) (*ClimateData, error) {
	if name == "" {
		return nil, fmt.Errorf("city name cannot be empty")
	}

	for _, city := range cities {
		if strings.EqualFold(city.City, name) {
			return &city, nil
		}
	}

	return nil, fmt.Errorf("city '%s' not found", name)
}

func getMenuChoice() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("error reading input: %v", err)
	}

	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid choice: please enter a number between 1 and 6")
	}

	if choice < 1 || choice > 6 {
		return 0, fmt.Errorf("invalid choice: please enter a number between 1 and 6")
	}

	return choice, nil
}

func main() {
	cities := []ClimateData{
		{City: "New York", AvgTemp: 22.5, AvgRain: 50.0},
		{City: "London", AvgTemp: 15.0, AvgRain: 30.0},
		{City: "Tokyo", AvgTemp: 25.0, AvgRain: 40.0},
		{City: "Sydney", AvgTemp: 28.0, AvgRain: 60.0},
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		displayMenu()
		
		choice, err := getMenuChoice()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		switch choice {
		case 1:
			displayAllCities(cities)

		case 2:
			highestTempCity, lowestTempCity := findTemperatureExtremes(cities)
			fmt.Printf("\nHighest Average Temperature: %s (%.2f °C)\n",
				highestTempCity.City, highestTempCity.AvgTemp)
			fmt.Printf("Lowest Average Temperature: %s (%.2f °C)\n",
				lowestTempCity.City, lowestTempCity.AvgTemp)

		case 3:
			averageRainfall := calculateAverageRainfall(cities)
			fmt.Printf("\nAverage Rainfall Across All Cities: %.2f mm\n", averageRainfall)

		case 4:
			threshold, err := promptThreshold()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}

			filteredCities := filterCitiesByRainfall(cities, threshold)
			if len(filteredCities) > 0 {
				fmt.Printf("\nCities with rainfall above %.2f mm:\n", threshold)
				for _, city := range filteredCities {
					fmt.Printf("%s: %.2f mm\n", city.City, city.AvgRain)
				}
			} else {
				fmt.Printf("\nNo cities found with rainfall above %.2f mm\n", threshold)
			}

		case 5:
			fmt.Print("\nEnter the name of the city to search: ")
			cityName, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Error reading input: %v\n", err)
				continue
			}

			cityName = strings.TrimSpace(cityName)
			foundCity, err := searchCity(cities, cityName)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}

			fmt.Printf("\nCity: %s\nAverage Temperature: %.2f °C\nAverage Rainfall: %.2f mm\n",
				foundCity.City, foundCity.AvgTemp, foundCity.AvgRain)

		case 6:
			fmt.Println("\nThank you for using the Climate Data Analysis Program!")
			return
		}
	}
}