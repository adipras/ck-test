package programs

import (
	"ck-test/utils"
	"fmt"
)

func ProcessCities() {
	// Read CSV file
	records, err := utils.ReadCSVFile("cities.csv")
	if err != nil {
		fmt.Printf("Error reading CSV file: %v\n", err)
		return
	}

	if len(records) == 0 {
		fmt.Println("CSV file is empty")
		return
	}

	// Find required column indexes from header
	header := records[0]
	countryIdIndex := -1
	stateIdIndex := -1
	countryNameIndex := -1

	for i, column := range header {
		switch column {
		case "country_id":
			countryIdIndex = i
		case "state_id":
			stateIdIndex = i
		case "country_name":
			countryNameIndex = i
		}
	}

	// Validate required columns exist
	if countryIdIndex == -1 || stateIdIndex == -1 || countryNameIndex == -1 {
		fmt.Println("Error: Required columns not found in CSV header")
		fmt.Printf("Available columns: %v\n", header)
		return
	}

	// Map to store unique states for each country
	countryStates := make(map[string]map[string]bool)
	countryNames := make(map[string]string) // Map country_id to country_name
	skippedRows := 0

	// Process data
	for _, record := range records[1:] {
		if len(record) <= countryIdIndex || len(record) <= stateIdIndex || len(record) <= countryNameIndex {
			skippedRows++
			continue
		}

		countryId := record[countryIdIndex]
		stateId := record[stateIdIndex]
		countryName := record[countryNameIndex]

		if countryId == "" || stateId == "" || countryName == "" {
			skippedRows++
			continue
		}

		// Initialize map for new country
		if _, exists := countryStates[countryId]; !exists {
			countryStates[countryId] = make(map[string]bool)
			countryNames[countryId] = countryName
		}

		// Add state to country's set of states
		countryStates[countryId][stateId] = true
	}

	// Convert to CountryCount slice
	var countryCounts []utils.CountryCount
	for countryId, states := range countryStates {
		countryCounts = append(countryCounts, utils.CountryCount{
			CountryName: countryNames[countryId],
			CityCount:   len(states), // Number of unique states
		})
	}

	// Sort the data
	ascCountries := utils.SortCountryCounts(countryCounts, true)
	descCountries := utils.SortCountryCounts(countryCounts, false)

	// Save results to JSON files
	err = utils.SaveToJSON(ascCountries, "output_asc.json")
	if err != nil {
		fmt.Printf("Error saving ascending order JSON: %v\n", err)
		return
	}

	err = utils.SaveToJSON(descCountries, "output_desc.json")
	if err != nil {
		fmt.Printf("Error saving descending order JSON: %v\n", err)
		return
	}

	// Display summary
	fmt.Printf("\nProcessing Summary:\n")
	fmt.Printf("Total rows in CSV: %d\n", len(records)-1)
	fmt.Printf("Skipped rows: %d\n", skippedRows)
	fmt.Printf("Unique countries: %d\n", len(countryStates))

	// Display results
	fmt.Println("\nTop 10 countries by number of states/cities:")
	fmt.Println("------------------------------------------------")
	fmt.Printf("%-30s %s\n", "Country", "Number of States/Cities")
	fmt.Println("------------------------------------------------")
	for i, cc := range descCountries {
		if i >= 10 {
			break
		}
		fmt.Printf("%-30s %d\n", cc.CountryName, cc.CityCount)
	}

	fmt.Println("\nBottom 10 countries by number of states/cities:")
	fmt.Println("------------------------------------------------")
	fmt.Printf("%-30s %s\n", "Country", "Number of States/Cities")
	fmt.Println("------------------------------------------------")
	for i, cc := range ascCountries {
		if i >= 10 {
			break
		}
		fmt.Printf("%-30s %d\n", cc.CountryName, cc.CityCount)
	}

	fmt.Println("\nResults have been saved to 'asc_output.json' and 'desc_output.json'")
}
