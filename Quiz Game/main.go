package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func dataParser(fileName string) (map[string]string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error while opening csv file:", err)
		os.Exit(1)
	}
	
	reader := csv.NewReader(file)
	reader.Comma = ','

	// Add your code for processing the file here

	record, panic := reader.ReadAll()
	if(panic != nil){
		fmt.Println("ReadAll error", panic)
		os.Exit(1)
	}

	file.Close()
	parsedData := make(map[string]string)
	for _, value := range(record){
		parsedData[value[0]] = value[1]
	}
	//fmt.Println(parsedData)
	return parsedData
}

func startQuiz(m map[string]string)(int){
	score := 0
	for value, data := range(m){
    fmt.Print(value," :")
	var usrinput string
	fmt.Scanln(&usrinput)
	if usrinput == data {
		score++
	}else{
		fmt.Println("Wrong! score is : ")
		return score
	}
	}
	return score
}

func main(){
	mapdata := dataParser("Quiz.csv")
	d := startQuiz(mapdata)
	fmt.Println("Score is: ", d)
}