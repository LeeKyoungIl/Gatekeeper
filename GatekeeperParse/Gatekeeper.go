package main

import (
	"gklib"
	"fmt"
	"time"
	"os"
	"strings"
	"strconv"
)

func main() {
	gklib.PrintgklibLogo()
	gklib.SetMaxCpuCore()
	// basic param from Args
	basicParam := gklib.GetBasicParam()
	// validation log file
	gklib.ValidationFile(basicParam["path"])

	fileReadTime := time.Now()

	lines, err := gklib.ReadLines(basicParam["path"])
	if err != nil {
		fmt.Println("Can't read file.")
		os.Exit(1)
	}

	gklib.TimeTrack(fileReadTime, "The time. It took to read file.")

	analysisTime := time.Now()
	// start analysis a log file
	get := gklib.AnalysisLog(lines, basicParam["addIgnoteUrl"])

	if len(get) == 0 {
		fmt.Println(basicParam["path"], "There is no extracted log file.")
		os.Exit(1)
	}

	gklib.TimeTrack(analysisTime, "The time. it took to analysis pattern of log file.")
	// end analysis a log file

	// get extracted pattern
	gklib.GetLogPattern(get)
	// compare of servers start message
	gklib.GetStartToCompareServersInfoMessage(get, basicParam["targetServer"])

	// GET
	for k, v := range get {
		var urlPath = k
		// make url
		urlPath = gklib.GetHttpUrl(urlPath, v)

		// start to compare servers
		targetServer := strings.Split(basicParam["targetServer"], ",")

		if len(targetServer) != 2 {
			fmt.Println("You must enter 2 server for comparsion.")
			os.Exit(1)
		}

		result := make(map[int]string)

		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Println("++")

		for serverIndex, v := range targetServer {
			//url string, timeoutSec int
			url := "http://" + v + urlPath
			apiCallTime := time.Now()

			fmt.Print("++ start api request -> url : ", url)

			get[k]["url"+strconv.Itoa(serverIndex)] = url

			result[serverIndex], get[k]["httpStatusCode"], err = gklib.GetDataByHttpGet(url, 60)

			if err != nil || get[k]["httpStatusCode"] != "200" {
				get[k]["server"] = v
				get[k]["error"] = "Invalid URL or error."

				if err != nil {
					get[k]["error"] += " (" + err.Error() + ")"
				}
			} else if get[k]["httpStatusCode"] == "200" && strings.Index(result[serverIndex], "status: 404") > -1 {
				get[k]["httpStatusCode"] = "404"
				get[k]["server"] = v
				get[k]["error"] = "There is no API result. You need to check."
			}

			fmt.Println(" 응답시간 : ", gklib.TimeTrackOnlyTime(apiCallTime))
		}

		fmt.Println("++")
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Println("")

		if eq, _ := gklib.JSONBytesEqual([]byte(result[0]), []byte(result[1])); !eq && len(result) == 2 {
			get[k]["error"] = "Please The API result value check before distribution."
		}
	}

	gklib.PrintApiComparedResult(get)
	os.Exit(0)
}
