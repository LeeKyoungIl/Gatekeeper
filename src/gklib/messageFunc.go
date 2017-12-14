package gklib

import (
	"fmt"
	"strconv"
	"time"
)

func PrintgklibLogo() {
	fmt.Println("")
	fmt.Println("     _____           _           _               _____       _       _                                        ")
	fmt.Println("    |  __ \\         (_)         | |             / ____|     | |     | |                                      ")
	fmt.Println("    | |__) | __ ___  _  ___  ___| |_   ______  | |  __  __ _| |_ ___| | _____  ___ _ __   ___ _ __            ")
	fmt.Println("    |  ___/ '__/ _ \\| |/ _ \\/ __| __| |______| | | |_ |/ _` | __/ _ \\ |/ / _ \\/ _ \\ '_ \\ / _ \\ '__|    ")
	fmt.Println("    | |   | | | (_) | |  __/ (__| |_           | |__| | (_| | ||  __/   <  __/  __/ |_) |  __/ |              ")
	fmt.Println("    |_|   |_|  \\___/| |\\___|\\___|\\__|           \\_____|\\__,_|\\__\\___|_|\\_\\___|\\___| .__/ \\___|_|  ")
	fmt.Println("                   _/ |                                                           | |                         ")
	fmt.Println("                  |__/                                                            |_|              	          ")
	fmt.Println("")
}

func TimeTrack(start time.Time, description string) {
	elapsed := time.Since(start)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("@@")
	fmt.Println("@@ "+description+" ", elapsed)
	fmt.Println("@@")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("")
	fmt.Println("")
}

func TimeTrackOnlyTime(start time.Time) time.Duration {
	return time.Since(start)
}

func GetLogPattern(get map[string]map[string]string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("@@")
	fmt.Println("@@ Extract log pattern total : ", strconv.Itoa(len(get)))
	fmt.Println("@@")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("")
	fmt.Println("")

	for k, _ := range get {
		fmt.Println("api pattern : ", k, " 총 실행수 : ", get[k]["count"])
	}
}

func GetStartToCompareServersInfoMessage(get map[string]map[string]string, targetServer string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("@@")
	fmt.Println("@@ 총 " + strconv.Itoa(len(get)) + "회의 " + targetServer + " Start the API comparison.")
	fmt.Println("@@")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("")
	fmt.Println("")
}

func PrintApiComparedResult(get map[string]map[string]string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("@@")
	fmt.Println("@@ List of API : different result or error")
	fmt.Println("@@")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("")
	fmt.Println("")

	errorResultCheck := false

	for _, v := range get {
		if _, ok := v["error"]; ok {
			if !errorResultCheck {
				errorResultCheck = true
			}

			fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
			fmt.Println("++ url 1 : ", v["url0"])
			fmt.Println("++ url 2 : ", v["url1"])
			fmt.Println("++ statusCode : ", v["httpStatusCode"])
			fmt.Println("++ errorMsg : ", v["error"])
			fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
			fmt.Println("")
		}
	}

	if !errorResultCheck {
		fmt.Println("All values are fine. You can also deploy.")
	}
}
