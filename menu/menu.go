package menu

import (
	"fmt"
	"os"

	"ck-test/programs"
)

func ShowMenu() {
	for {
		fmt.Println("\nAvailable Programs:")
		fmt.Println("1. Group Cities by Country")
		fmt.Println("2. Root Word Replacement")
		fmt.Println("0. Exit")
		fmt.Print("\nEnter program number to run (0 to exit): ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			programs.ProcessCities()
		case "2":
			programs.ProcessRootWords()
		case "0":
			fmt.Println("Exiting program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
