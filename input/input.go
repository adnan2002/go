package input

import (
	f "fmt"
	"os"
	"strconv"
	"strings"
)


func GetPricesArray(filePath string) []int {

	text_bytes, err := os.ReadFile(filePath) 

	if err != nil {
		f.Println("Error: ", err)
		panic(err) 
	}
	text := string(text_bytes)
	textArr := strings.Split(text, "\r\n")
	integerArr := make([]int, len(textArr)-1, cap(textArr))

	for i:=1; i<len(textArr); i++ {
		result, err := handleStringToIntegerConvertion(textArr[i])
		if err != nil {
			panic(err)
		}
		integerArr[i-1] = result


		

	}

	return integerArr
	
}


func handleStringToIntegerConvertion(text string) (int, error) {
	var result int 
	var err error
	if text != "" {
	result, err = strconv.Atoi(text)
	}

	if err != nil {
		f.Println("Error: ", err)
		return -1, err
	}

	return result, nil
}
