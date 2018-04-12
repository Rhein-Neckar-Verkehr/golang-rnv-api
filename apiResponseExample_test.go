// # Unit Tests for apiUtility.
package api

// Imports needed for this test:
import (
	"fmt"
)

// ## stationMonitor.go Response Objects


func ExampleJourneyResponseObject(){	
	var ResponseObjecct Journey
	var errorMessage =[]byte(`<errordescription>Could not parse parameter "2563" as date</errordescription>`)
	var responseObjectFromStationMonitor =[]byte(`{
			"ticker": "*** A ticker for this Response Object ****** another ticker ***",
			"listOfDepartures": [{
				"time": "14:58+1",
				"status": "OK",
				"direction": "Weinheim - Mannheim",
				"platform": "C",
				"transportation": "STRAB",
				"tourId": "41282-33671-1",
				"kindOfTour": "454AUS",
				"positionInTour": "38",
				"statusNote": "",
				"lineId": "RNV005_RNV005A",
				"lineLabel": "5",
				"differenceTime": "0",
				"foreignLine": "false",
				"newsAvailable": "false"			
			}],
			"pastRequestText": "n/a",
			"updateTime": "0",
			"updateIterations": "0",
			"stationInfos": [],
			"time": "14:57",
			"icon": "icon_tram.png",
			"label": "Richtung ",
			"color": "002a4e",
			"shortLabel": "n/a",
			"projectedTime": "n/a",
			"otherTimes": [],
			"otherProjectedTimes": []
		}`)
	err:= parseResponseObject(errorMessage, &ResponseObjecct)
	fmt.Println(err)
	err= parseResponseObject(responseObjectFromStationMonitor, &ResponseObjecct)
	fmt.Println(err)
	//fmt.Println(ResponseObjecct) todo implement to string method
	//fmt.Println(ResponseObjecct.SplitMultipleTicker()) // todo fix this [ " A ticker for this Response Object " " another ticker "]
	fmt.Println(ResponseObjecct.ListOfDepartures[0].IsRealtimeData())
	fmt.Println(ResponseObjecct.ListOfDepartures[0].IsForeignLine())
	fmt.Println(ResponseObjecct.ListOfDepartures[0].IsBus())
	fmt.Println(ResponseObjecct.ListOfDepartures[0].IsTram())
	fmt.Println(ResponseObjecct.ListOfDepartures[0].GetGTFSExtendedRouteType())
	fmt.Println(ResponseObjecct.ListOfDepartures[0].IsCancelled())
	fmt.Println(ResponseObjecct.ListOfDepartures[0].AreNewsAvailable())
	fmt.Println(*ResponseObjecct.ListOfDepartures[0].GetDelayIfRealtime())
	fmt.Println(ResponseObjecct.ListOfDepartures[0].GetRNVGTFSRouteId())
	// Output:
	// <errordescription>Could not parse parameter "2563" as date</errordescription>
	// <nil>
	// true
	// false
	// false
	// true
	// 900
	// false
	// false
	// 1
	// 5
	
}