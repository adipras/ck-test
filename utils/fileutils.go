package utils

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"sort"
)

// CountryCount represents the count of cities in a country
type CountryCount struct {
	CountryName string `json:"country_name"`
	CityCount   int    `json:"city_count"`
}

func ReadCSVFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	reader.LazyQuotes = true

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func SaveToJSON(data []CountryCount, filename string) error {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}

func SortCountryCounts(counts []CountryCount, ascending bool) []CountryCount {
	sorted := make([]CountryCount, len(counts))
	copy(sorted, counts)

	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].CityCount == sorted[j].CityCount {
			return sorted[i].CountryName < sorted[j].CountryName
		}
		if ascending {
			return sorted[i].CityCount < sorted[j].CityCount
		}
		return sorted[i].CityCount > sorted[j].CityCount
	})

	return sorted
}
