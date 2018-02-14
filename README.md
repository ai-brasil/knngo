# knngo

Go library to machine learning classification problems

## Getting Started

Run at terminal``` go get -u github.com/ai-brasil/knngo  ```

Import the package```go import "github.com/italojs/knngo"   ```

### Prerequisites

1º: Only the class column can be a NaN (Not a Number)

2º: The column class must be the last column.

3º: When you read your file (e.g. your .CSV file) don't worry to transform the type of your data to number, you can use the values as string.

## Methods

```go 
knn.PrepareDataset(percent float32, records [][]string) (train [][]string, test [][]string, err error)  
```

```go
func Classify(trainData [][]string, dataToPredict []string, k int) (result string, err error)
```

## Example

```go
package main

import (
	"fmt"
	"github.com/italojs/knngo"
)

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
```

## Authors

* **Italo Jose* - *Initial work* - [italojs](https://github.com/italojs)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
