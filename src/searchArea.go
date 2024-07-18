package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/schollz/progressbar/v3"
)

func upperBoundAreaCode(areaCode int, left int, right int, chBound chan BoundChannel) {
	if right > left {
		mid := (right + left) / 2
		std := FetchScore(SBDFormat(areaCode, mid))

		if std != nil {
			upperBoundAreaCode(areaCode, mid+1, right, chBound)
		} else {
			upperBoundAreaCode(areaCode, left, mid-1, chBound)
		}
	}

	std := FetchScore(SBDFormat(areaCode, right))

	if std != nil {
		chBound <- BoundChannel{
			areaCode: areaCode,
			bound:    right,
		}
	} else {
		chBound <- BoundChannel{
			areaCode: areaCode,
			bound:    right - 1,
		}
	}
}

func SearchAreaRange() map[int]int {
	fmt.Println("\nSearching SBD range of each area...")
	fmt.Println()
	bar := progressbar.Default(64)

	res := make(map[int]int)
	chBound := make(chan BoundChannel)
	defer close(chBound)

	for areaCode := 1; areaCode <= 64; areaCode++ {
		bound := 90000
		if areaCode == 1 {
			bound = 200000
		}
		go upperBoundAreaCode(areaCode, 1, bound, chBound)
	}

	for i := 1; i <= 64; i++ {
		bound := <-chBound
		bar.Add(1)
		res[bound.areaCode] = bound.bound
	}

	return res
}

func LoadAreaRangeFile() (data map[int]int) {
	filename := "area_range.json"
	file, err := os.Open(filename)

	if err != nil {
		return nil
	}

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &data)

	file.Close()
	return data
}

func SaveAreaRangeFile(data map[int]int) {
	filename := "area_range.json"
	file, _ := os.Create(filename)

	byteValue, _ := json.MarshalIndent(data, "", "\t")
	file.Write(byteValue)
	file.Close()
}
