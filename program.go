package main

import (
	"api/Models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

// -----------------------------------------------------
// Main
// -----------------------------------------------------

func main() {
	for {
		var timeFrame int = SelectTimeFrame()
		GetApiStockData(timeFrame)
	}
}

// -----------------------------------------------------
// Functions
// -----------------------------------------------------

func SelectTimeFrame() int {
	var timeFrameInput string
	var timeFrameResult int
	fmt.Println("=======================================================================")
	fmt.Println("	                     SCANNER SIGNAL								    ")
	fmt.Println("=======================================================================")
	fmt.Println("Type 1 :                                                             1D")
	fmt.Println("Type 2 :                                                             1W")
	fmt.Println("Type 3 :                                                             1M")
	fmt.Println("Type 4 :                                                             1Y")
	fmt.Println("=======================================================================")
	fmt.Printf("Please select type to scanner 1, 2, 3, 4 : ")
	n, err := fmt.Scanf("%s\n", &timeFrameInput)
	if err != nil || n == 0 {
		fmt.Println(n, err)
	}
	if len(timeFrameInput) != 1 {
		return 0
	}

	// regex
	numberRegexp := regexp.MustCompile(`[1-4]{1}`)
	timeFrameRegex := numberRegexp.FindString(timeFrameInput)
	timeFrameResult, err = strconv.Atoi(timeFrameRegex)
	if err != nil {
		fmt.Println(n, err)
	}
	fmt.Println("=======================================================================")

	return timeFrameResult
}

func GetApiStockData(timeFrame int) {
	if timeFrame == 1 || timeFrame == 2 || timeFrame == 3 || timeFrame == 4 {
		response, err := http.Get("https://storage.googleapis.com/finno-ex-re-v2-static-staging/recruitment-test/fund-ranking-" + GetTimeFrameString(timeFrame) + ".json")
		if err != nil {
			fmt.Print(err.Error())
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		if response.StatusCode == http.StatusNotFound {
			fmt.Println("Not found")
		}

		var responseObject coin.Response
		json.Unmarshal(responseData, &responseObject)
		if len(responseObject.Data) > 0 {
			// Format time
			current_time := time.Now()

			// Display
			fmt.Println("===============================================================================================================================================================================================")
			fmt.Printf("%20v	:                 %10v         :                 %10v         :                 %10v             :                 %10v           \n", "Symbol", "Nav_Return", "Nav_Date", "Avg_Return", "Nav")
			fmt.Println("===============================================================================================================================================================================================")
			for _, res := range responseObject.Data {
				fmt.Printf("%20v	:                 %10v         :                 %10v         :                 %10v             :                 %10v           \n", res.ThailandFundCode, res.NavReturn, current_time.Format("2006-01-02"), res.AvgReturn, res.Nav)
			}
			fmt.Println("===============================================================================================================================================================================================")
		}
	}
}

func GetTimeFrameString(timeFrame int) string {
	switch timeFrame {
	case 1:
		return "1D"
	case 2:
		return "1W"
	case 3:
		return "1M"
	case 4:
		return "1Y"
	}

	return "1Y"
}
