package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)
type problem struct{
	q string
	a string
}
func dataParser(fileName string) (map[string]string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error while opening csv file:", err)
		os.Exit(1)
	}
	
	reader := csv.NewReader(file)
	//reader.Comma = ','

	// Add your code for processing the file here

	record, panic := reader.ReadAll()
	if(panic != nil){
		fmt.Println("ReadAll error", panic)
		os.Exit(1)
	}

	file.Close()
	parsedData := make(map[string]string)
	for _, value := range(record){
		var quiz problem 
		quiz.q = value[0]
		quiz.a = value[1]
		parsedData[quiz.q] = quiz.a
	}
	//fmt.Println(parsedData)
	return parsedData
}

func startQuiz(m map[string]string, c chan int){
	score := 0
	
	for value, data := range(m){
    fmt.Print(value," :")
	
	var usrinput string
	fmt.Scanln(&usrinput)
	 
	if usrinput == data {
	  score++
	}else{
		fmt.Println("Wrong! ")
		c <- score
	}
	}
	c <- score
}

// func initQ()(map[string]string){
// 	csvFilename := flag.String("csv", "Quiz.csv", "File name where csv data is stored")
// 	flag.Parse()
// 	mapdata := dataParser(*csvFilename)
// 	return mapdata
// }

func main(){
	t := flag.Int("time", 100, "Give time limit to your quiz for added fun!")
	flag.Parse()
	timedQuiz(t)
	// mapdata := initQ()
	// d := startQuiz(mapdata)
	// fmt.Println("Score is: ", d)
}

func timer(t int,c chan int){
	newTimer := time.NewTimer(time.Duration(t) * time.Second)
	<- newTimer.C
	c <- -1
}

func timedQuiz(t *int){
	csvFilename := flag.String("csv", "Quiz.csv", "File name where csv data is stored")
	flag.Parse()
	mapdata := dataParser(*csvFilename)
	tm := make(chan int)
	q := make(chan int)
	go startQuiz(mapdata, q) 
	go timer(*t, tm)
	
	select {
	case msg := <- q:
		fmt.Println(msg)
	case msg := <- tm:
		if msg == -1{
			fmt.Println("Ran out of time!")
			os.Exit(1)
		}
		

	}
}