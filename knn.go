package knn

import (
	"sync"
	"errors"
	"math"
	"sort"
	"strconv"
)

// ColumnClassIndex is the index that is the class of your data(line) in dataset. By default it's be the first column of your dataset
var ColumnClassIndex int

func distinct(elements *[]string) (result []string) {
	encountered := map[string]bool{}

	for i := range *elements {
		if encountered[(*elements)[i]] == true {
		} else {
			encountered[(*elements)[i]] = true
			result = append(result, (*elements)[i])
		}
	}
	return result
}
func getCollum(elements *[][]string, ColumnIndex int) (column []string) {
	for i := range *elements {
		column = append(column, (*elements)[i][ColumnIndex])
	}
	return
}
func getValuesByClass(records *[][]string, class string) (newRecords [][]string) {
	newRecords = make([][]string, 0)
	for l := range *records {
		if (*records)[l][len((*records)[l])-1] == class {
			newRecords = append(newRecords, (*records)[l])
		}
	}
	return
}
func divideInPercent(records *[][]string, percent float32) (newRecords [][]string, residue [][]string) {
	i := float32(len(*records)) * percent
	for i >= 0 {
		index := int(i - 1)
		newRecords = append(newRecords, (*records)[index])
		i--
	}

	i = float32(len(*records)) * percent
	j := float32(len(*records))
	for j > i {
		index := int(j - 1)
		residue = append(residue, (*records)[index])
		j--
	}
	return
}
func euclideanDist(pi *[]string, qi *[]string) (result float64, err error) {
	i := len(*pi) - 1

	for i >= 0 {
		pif, err := strconv.ParseFloat((*pi)[i], 32)
		if err != nil {
			return 0, err
		}
		qif, err := strconv.ParseFloat((*qi)[i], 32)
		if err != nil {
			return 0, err
		}
		result += math.Pow(pif-qif, 2)
		i--
	}
	result = math.Sqrt(result)

	return
}
func getMapValues(m *map[float64]string) (values []string) {
	for _, v := range *m {
		values = append(values, v)
	}
	return
}
func getMapKeys(m *map[float64]string) (keys []float64) {
	for k := range *m {
		keys = append(keys, k)
	}
	return
}
func getKnn(list *map[float64]string, k int) (sortedMap map[float64]string) {
	keys := getMapKeys(&(*list))

	sort.Float64s(keys)

	sortedMap = make(map[float64]string)

	for i, key := range keys {
		if i < k {
			sortedMap[key] = (*list)[key]
		} else {
			break
		}
	}
	return
}
func getPredominantClass(knn *map[float64]string) (class string) {
	mapValues := getMapValues(&(*knn))
	classes := distinct(&mapValues)
	var predominantClass int
	for c := range classes {
		var countClass int
		for i := range *knn {
			if (*knn)[i] == classes[c] {
				countClass++
			}
		}
		if predominantClass < countClass {
			predominantClass = countClass
			class = classes[c]
		}
	}
	return
}

// PrepareDataset divide a dataset(records) in x percet to train and the rest data get to test
func PrepareDataset(percent float32, records [][]string) (train [][]string, test [][]string, err error) {
	column := getCollum(&records, len(records[0])-1)
	classes := distinct(&column)
	for i := range classes {
		values := getValuesByClass(&records, classes[i])
		trainRecords, testRecords := divideInPercent(&values, percent)
		train = append(train, trainRecords[0:]...)
		test = append(test, testRecords[0:]...)
	}
	return
}

// Classify will predict a new data(not classificated yet) based in all you train data(already classificated)
func Classify(trainData [][]string, dataToPredict []string, k int) (result string, err error) {
	dists := make(map[float64]string)
	for i := range trainData {
		class := trainData[i][ColumnClassIndex]
		d, err := euclideanDist(&trainData[i], &dataToPredict)
		if err != nil {
			return "Error: no class predicted", err
		}
		dists[d] = class
	}
	knn := getKnn(&dists, k)
	result = getPredominantClass(&knn)

	return
}

