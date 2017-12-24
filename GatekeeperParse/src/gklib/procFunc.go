package gklib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func SetMaxCpuCore() {
	coreCount := runtime.NumCPU()

	if coreCount > 1 {
		coreCount -= 1
	}

	runtime.GOMAXPROCS(coreCount)
}

func GetBasicParam() map[string]string {
	if len(os.Args) != 4 {
		fmt.Println("Please, Check the required value.")
		os.Exit(1)
	}

	basicParam := make(map[string]string)

	basicParam["path"] = os.Args[1] + os.Args[2]

	if len(os.Args) == 4 {
		targetServer := strings.Split(os.Args[3], ",")
		if len(targetServer) != 2 {
			fmt.Println("You must enter 2 server for comparsion.")
			os.Exit(1)
		}

		basicParam["targetServer"] = os.Args[3]
	}

	if len(os.Args) == 5 {
		basicParam["addIgnoteUrl"] = os.Args[4]
	}

	return basicParam
}

func setReturnData(get map[string]map[string]string, urlStatus map[string]string, method string, urlKey string) (map[string]map[string]string, map[string]string) {
	urlKey = strings.TrimSpace(urlKey)

	switch method {
	case "GET":
		if val, ok := get[urlKey]; ok {
			nowCount, err := strconv.Atoi(val["count"])
			if err != nil {
				panic(err)
			}
			val["count"] = strconv.Itoa(nowCount + 1)
			get[urlKey] = val
		} else {
			urlStatus["count"] = "1"
			get[urlKey] = urlStatus
		}
		//case "POST":
		//	if val, ok := post[urlKey]; ok {
		//		nowCount, err := strconv.Atoi(val["count"])
		//		if err != nil {
		//			panic(err)
		//		}
		//		val["count"] = strconv.Itoa(nowCount + 1)
		//		post[urlKey] = val
		//	} else {
		//		urlStatus["count"] = "1"
		//		post[urlKey] = urlStatus
		//	}
	}

	return get, urlStatus
}

func AnalysisLog(lines []string, addIgnoteUrl string) map[string]map[string]string {
	if len(lines) > 0 {
		fmt.Println("")
		fmt.Println("")
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		fmt.Println("@@")
		fmt.Println("@@ toral " + strconv.Itoa(len(lines)) + "lines are analyzing the log file.")
		fmt.Println("@@")
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		fmt.Println("")
		fmt.Println("")
	} else {
		fmt.Println("There is no log file info.")
		os.Exit(1)
	}

	get := make(map[string]map[string]string)

	ignoreUrl := [8]string{"refundToReceiver", "completedProcess", "order", "adminUser", "buyer", "seller", "Order", "voucher"}
	regExp := regexp.MustCompile("([0-9]{2}/[a-zA-Z]+/[0-9]{4}:[0-9]{2}:[0-9]{2}:[0-9]{2}).*(GET|POST)\\s(/.*)\\sHTTP.*\"\\s([0-9]{3})")
	regExpByUrl := regexp.MustCompile("(/[0-9]+|/[0-9]+/)")

	showIndex := len(lines) / 100
	nowShowIndex := 0

	for index, _ := range lines {
		if (nowShowIndex * showIndex) == index {
			fmt.Println("->(", nowShowIndex, "%)")
			nowShowIndex++
		}

		line := lines[index]
		validateUrl := true

		for _, element := range ignoreUrl {
			if strings.Index(line, element) > -1 {
				validateUrl = false
			}
		}

		if validateUrl && len(addIgnoteUrl) > 0 {
			tmpAddIgnore := strings.Split(addIgnoteUrl, ",")

			for _, element := range tmpAddIgnore {
				if strings.Index(line, strings.TrimSpace(element)) > -1 {
					validateUrl = false
				}
			}
		}

		if !validateUrl {
			continue
		}

		result := regExp.FindStringSubmatch(line)

		// Only successful cases will be handled separately for error logs
		if result[4] != "200" {
			continue
		}

		urlKey := AddSuffix(result[3], "/")

		urlStatus := make(map[string]string)
		urlStatus["code"] = result[4]

		u, err := url.Parse(urlKey)
		if err != nil {
			panic(err)
		}

		if len(u.RawQuery) > 0 {
			urlStatus["param"] = TrimSuffix(u.RawQuery, "/")
			urlKey = u.Path
		}

		// Remove data from key
		resultUrl := regExpByUrl.FindAllStringSubmatch(urlKey, -1)
		if len(resultUrl) > 0 {
			var stringBuffer bytes.Buffer

			for index, _ := range resultUrl {
				addString := strings.Replace(resultUrl[index][1], "/", "", -1)

				re := regexp.MustCompile("/(" + addString + ")")

				if re.MatchString(urlKey) {
					urlKey = strings.Replace(urlKey, "/"+addString, "/{data"+strconv.Itoa(index)+"}", 1)
				}

				var delimiter string = ""

				if stringBuffer.Len() > 0 {
					delimiter = ","
				}

				stringBuffer.WriteString(delimiter + addString)
			}

			urlStatus["data"] = stringBuffer.String()
		}

		urlKey = TrimSuffix(urlKey, "/")

		setReturnData(get, urlStatus, result[2], urlKey)

		//tmpDateString := strings.Split(result[1], "/")
		//tmpYearAndTimeString := strings.Split(tmpDateString[2], ":")
		//var tmpMonth string
		//
		//switch tmpDateString[1] {
		//	case "Jan":
		//		tmpMonth = "01"
		//	case "Feb":
		//		tmpMonth = "02"
		//	case "Mar":
		//		tmpMonth = "03"
		//	case "Apr":
		//		tmpMonth = "04"
		//	case "May":
		//		tmpMonth = "05"
		//	case "Jun":
		//		tmpMonth = "06"
		//	case "Jul":
		//		tmpMonth = "07"
		//	case "Aug":
		//		tmpMonth = "08"
		//	case "Sep":
		//		tmpMonth = "09"
		//	case "Oct":
		//		tmpMonth = "10"
		//	case "Nov":
		//		tmpMonth = "11"
		//	case "Dec":
		//		tmpMonth = "12"
		//}

		//logtime, err := time.Parse(time.RFC3339, (tmpYearAndTimeString[0]+"-"+tmpMonth+"-"+tmpDateString[0]+"T"+tmpYearAndTimeString[1]+":"+tmpYearAndTimeString[2]+":"+tmpYearAndTimeString[3]+"+00:00"))
		//
		//if err == nil && (t.Unix() - logtime.Unix()) <= 3600 {
		//	fmt.Println(logtime.Unix())
		//}
	}

	fmt.Println("")

	return get
}

// JSONEqual compares the JSON from two Readers.
func JSONEqual(a, b io.Reader) (bool, error) {
	var j, j2 interface{}
	d := json.NewDecoder(a)
	if err := d.Decode(&j); err != nil {
		return false, err
	}
	d = json.NewDecoder(b)
	if err := d.Decode(&j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}

// JSONBytesEqual compares the JSON in two byte slices.
func JSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}

func inArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}
