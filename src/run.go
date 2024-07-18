package src

import (
	"fmt"
	"os"

	"github.com/schollz/progressbar/v3"
)

func WriteStudent(std *Student, areaCode int) {
	filename := fmt.Sprintf("%s/%02d.csv", os.Getenv("OUTPUT_FOLDER"), areaCode)
	AppendToFile(filename, std)
	AppendToFile(fmt.Sprintf("%s/total.csv", os.Getenv("OUTPUT_FOLDER")), std)
}

func RunThread(id int, areaCode int, guard chan struct{}) {
	sbd := SBDFormat(areaCode, id)
	student := FetchScore(sbd)

	if student != nil {
		WriteStudent(student, areaCode)
	}
	<-guard
}

func Run(areaCodeRangeMap map[int]int, maxThread int) {
	fmt.Println("\nFetching scores...")
	fmt.Println()
	bar := progressbar.Default(int64(NumberOfStudent(areaCodeRangeMap)))

	fileHeader := "SBD,Toán,Văn,Lý,Hoá,Sinh,Lịch Sử,Địa Lý,GDCD,Ngoại Ngữ"

	if _, err := os.Stat("data"); os.IsNotExist(err) {
		os.Mkdir(os.Getenv("OUTPUT_FOLDER"), 0755)
	}
	// Add header to total file
	totalFileName := os.Getenv("TOTAL_FILENAME")
	AppendToFile(fmt.Sprintf("%s/%s", os.Getenv("OUTPUT_FOLDER"), totalFileName), fileHeader)

	guard := make(chan struct{}, maxThread)
	// Loop through all areas
	for areaCode, areaCodeRange := range areaCodeRangeMap {
		filename := fmt.Sprintf("%s/%02d.csv", os.Getenv("OUTPUT_FOLDER"), areaCode)
		file, _ := os.Create(filename)
		file.Close()
		AppendToFile(filename, fileHeader)

		for i := 1; i <= areaCodeRange; i += 1 {
			guard <- struct{}{} // will block if max threads

			go RunThread(i, areaCode, guard)
			bar.Add(1)
		}
	}
	bar.Finish()
}
