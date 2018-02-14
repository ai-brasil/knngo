package main

import (
	"encoding/csv"
	"log"
	"os"
	"fmt"
	"github.com/italojs/knn-usage/algorithm"
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
	records := readFile("../datas/ocr-train.csv")

	// Use the PrepareDataset method is optional
	percentToTrain := float32(0.6)
	train, test := knn.PrepareDataset(percentToTrain, records)

	// You can choose some not classificated data to classify
	// the line 5 of test dataset was choosed randomicly by programmer just to example
	result := knn.Classify(train, test[5], 10)
	
	fmt.Println(result)
}