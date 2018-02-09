# knngo

Go library to machine learning classification problems

## Getting Started

Run at terminal``` go get -u github.com/ai-brasil/knngo  ```

### Prerequisites

1º: Yours class must be at last column at right side

2º: current oly work with numbers

3º: when you read your file (e.g. your .CSV file) don't worry to transform the type of your data to number, you can use the values as string.

## Methods

```go 
knn.PrepareDataset(percent float32, records [][]string) (train [][]string, test [][]string)  
```

```go
knn.Classify(train [][]string, dataToPredict []string, k int) (result string) 
```

Docs is in progress yet

## Example

```go
package main

import (
	"fmt"
	"github.com/italojs/knn-usage/algorithm"
)

func main (){
	records := readFile("datas/wdbc.csv")

	train, test := knn.PrepareDataset(0.6, records)

	index := 5
	result := knn.Classify(train,test[index],10)
	
	fmt.Println(result)
} 
```

## Authors

* **Italo Jose* - *Initial work* - [italojs](https://github.com/italojs)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
