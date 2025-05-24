package saveJson

import (
	"encoding/json"
	"fmt"
	"os"
)

func Save[T any](data []T) {
	// Use MarshalIndent instead of Marshal for pretty output
	dataJson, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		fmt.Println("Error", err)
		panic(err)
	}

	err = os.WriteFile("results.json", dataJson, 0644)

	if err != nil {
		fmt.Println("Error", err)
		panic(err)
	}
}
