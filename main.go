package main

import (
	"crawlscore/src"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	// ASCII ART
	fmt.Println(`

	████████╗██╗  ██╗██████╗ ████████╗ ██████╗  ██████╗               ██████╗  ██████╗ ██████╗ ██╗  ██╗
	╚══██╔══╝██║  ██║██╔══██╗╚══██╔══╝██╔═══██╗██╔════╝               ╚════██╗██╔═████╗╚════██╗██║  ██║
	   ██║   ███████║██████╔╝   ██║   ██║   ██║██║  ███╗    █████╗     █████╔╝██║██╔██║ █████╔╝███████║
	   ██║   ██╔══██║██╔═══╝    ██║   ██║▄▄ ██║██║   ██║    ╚════╝    ██╔═══╝ ████╔╝██║██╔═══╝ ╚════██║
	   ██║   ██║  ██║██║        ██║   ╚██████╔╝╚██████╔╝              ███████╗╚██████╔╝███████╗     ██║
	   ╚═╝   ╚═╝  ╚═╝╚═╝        ╚═╝    ╚══▀▀═╝  ╚═════╝               ╚══════╝ ╚═════╝ ╚══════╝     ╚═╝
	`)

	// Load .env data
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	maxThread, _ := strconv.Atoi(os.Getenv("MAX_THREAD"))
	fmt.Print("Configurations:\n\n")
	fmt.Printf(" - Data Source: https://vietnamnet.vn/giao-duc/diem-thi/tra-cuu-diem-thi-tot-nghiep-thpt\n")
	fmt.Printf(" - Max Thread: %d\n", maxThread)
	fmt.Printf(" - Output Folder: %s/\n", os.Getenv("OUTPUT_FOLDER"))

	// Load area range
	areaCodeRangeMap := src.LoadAreaRangeFile()
	if areaCodeRangeMap == nil {
		areaCodeRangeMap = src.SearchAreaRange()
		src.SaveAreaRangeFile(areaCodeRangeMap)
	}

	// Fetch scores
	src.Run(areaCodeRangeMap, maxThread)

	fmt.Print("\nFinished!\n\n")
}
