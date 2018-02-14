package main

import (
	"encoding/csv"
	"log"
	"os"
	"fmt"
	"github.com/italojs/knngo"
)

func readFile(path string)(record [][]string){
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("Error: ", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	record, err = reader.ReadAll()
	if err != nil {
		log.Fatalln("Error:", err)
		return
	}
	return
}

func main (){
	records := readFile("../datas/wdbc.csv")
	
	// Use the PrepareDataset method is optional
	percentToTrain := float32(0.6)
	train, test, err := knn.PrepareDataset(percentToTrain, records)
	if err != nil{
		log.Fatalln(err)
	}

	// You can choose some not classificated data to classify
	// The line 5 of test dataset was choosed randomicly by programmer just to example
	result, err := knn.Classify(train, test[5], 10)
	if err != nil{
		log.Fatalln(err)
	}

	fmt.Println(result)
}