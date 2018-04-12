// # testable Examples for  rnv Start.Info API (1.12) Parameter.
package api
import (
	"fmt"
	"time"
)

var hafasIdValid=[]string{"1236","1648","125"}
var hafasIdInValid=[]string{"12A6","","1234567"}

// ## stations.go Parameter
// TODO implement validtion ...
// func ExampleUpdateParams(){}

// ## news.go Parameter

func ExampleTickerParams(){
	params:=NewTickerParams()
	err:=params.Validate()
	fmt.Println(err)
	fmt.Println(params)
	params,err=params.AddLine("9")
	fmt.Println(err)
	fmt.Println(params)
	params,err=params.AddLine("969325")
	fmt.Println(err)
	fmt.Println(params)
	// Output:
	// <nil>
	// []
	// <nil>
	// [9]
	// Invalid value for lines or lineId
	// [9]
}


// ## lines.go Parameter
// Requests to Fahrtverläufe requrire Data from an Departure Object

func ExampleLineJourneyParams(){
	// Requests to Fahrtverläufe requires Data from an Departure Object
	// use this method for getting LineJourneyParams
	dep:=Departure{"17:03+ 0","OK","Neckarstadt West","C","STRAB","11836-10327-1","454AUS","15","","RNV002_RNV002A","2","0","false","false"}
	params,err:=NewLineJourneyParamsFromDepartureAndStationId(dep,hafasIdValid[0]) // Note: just an example, the hafasId may not fit to Departure Object
	fmt.Println(err)
	y := 2017
	m := time.March
	d := 17
	h := 13
	min := 24
	s := 0
	t := time.Date(y, m, d, h, min, s, 0, time.UTC)
	params.SetTime(t) // needs to be set to testable time
	params.Validate()
	fmt.Println(err)
	fmt.Println(params)
	err=params.SetStopIndex(dep.PositionInTour)
	fmt.Println(err)
	fmt.Println(params)
	// Output:
	// <nil>
	// <nil>
	// {1236 2 0 454 11836-10327-1 2017-03-17+13:24}
	// <nil>
	// {1236 2 15 454 11836-10327-1 2017-03-17+13:24}

}

func ExampleCancelledLineParams(){
	params:=CancelledLineParams{nil,nil}
	err:=params.Validate()
	fmt.Println(err)
	fmt.Println(params)
	err=params.SetLineId("23")
	fmt.Println(err)
	fmt.Println(params)
	err=params.SetLineId("4 / 4A")
	fmt.Println(err)
	fmt.Println(params)
	err=params.SetLineId("23B")
	fmt.Println(err)
	fmt.Println(params)
	y := 2017
	m := time.March
	d := 17
	h := 13
	min := 24
	s := 0
	t := time.Date(y, m, d, h, min, s, 0, time.UTC)
	params.SetDepartureTime(t)
	err=params.Validate()
	fmt.Println(err)
	fmt.Println(params)
	ivT:="2017-03-17+11:00"
	params.DepartureTime=&ivT
	err=params.Validate()
	fmt.Println(err)
	fmt.Println(params)
	// Output:
	// <nil>
	// {<nil> <nil>}
	// <nil>
	// {23 <nil>}
	// Invalid value for lines or lineId
	// {23 <nil>}
	// Invalid value for lines or lineId
	// {23 <nil>}
	// <nil>
	// {23 1489757040}
	// Invalid value for time or departureTime
	// {23 <nil>}


}

// ## info.go Parameter

func ExampleJourneyInfoTranserParams(){
	params:=JourneyInfoTranserParams{nil,nil,nil}
	fmt.Println(params)
	// Setting Poles
	err:=params.AddPole("3")
	fmt.Println(err)
	fmt.Println(params)
	err=params.AddPole("A")
	fmt.Println(err)
	fmt.Println(params)
	err=params.AddPole("356")
	fmt.Println(err)
	fmt.Println(params)
	err=params.AddPole("03")
	fmt.Println(err)
	fmt.Println(params)
	err=params.Validate()
	fmt.Println(err)
	// Setting departureTime
	y := 2017
	m := time.March
	d := 17
	h := 13
	min := 24
	s := 0
	t := time.Date(y, m, d, h, min, s, 0, time.UTC)
	// set to a time
	params.SetDepartureTime(t)
	err=params.Validate()
	fmt.Println(err)
	fmt.Println(params)
	tString:=" "
	params.DepartureTime=&tString
	err=params.Validate()
	fmt.Println(err)
	fmt.Println(params)
	params.DepartureTime=nil
	// Set hafasId
	params.HafasId=&hafasIdValid[2]
	err=params.Validate()
	fmt.Println(err)
	fmt.Println(params)
	params.HafasId=&hafasIdInValid[2]
	err=params.Validate()
	fmt.Println(err)
	fmt.Println(params)

	// Output:
	// {<nil> <nil> <nil>}
	// <nil>
	// {<nil> [3] <nil>}
	// Invalid value for poles
	// {<nil> [3] <nil>}
	// Invalid value for poles
	// {<nil> [3] <nil>}
	// <nil>
	// {<nil> [3 03] <nil>}
	// Not valid Combination of JourneyInfoTranser. Valid Combinations are:  HafasId | HafasId + Poles | HafasId + DepartureTime | HafasId + DepartureTime + Poles | DepartureTime
	// <nil>
	// {<nil> [3 03] 1489757040}
	// Invalid value for time or departureTime
	// {<nil> [3 03]  }
	// <nil>
	// {125 [3 03] <nil>}
	// Invalid value for stationId or hafasId
	// {1234567 [3 03] <nil>}
}

//
//
func ExampleStationInfoTranserParams(){
	params:=StationInfoTransferParams{nil,nil,nil}
	err:=params.Validate()
	fmt.Println(err)
	fmt.Println(params)
	err=params.AddLine("1")
	fmt.Println(err)
	fmt.Println(params)
	err=params.AddLine("4A")
	fmt.Println(err)
	fmt.Println(params)
	err=params.AddLine("368")
	fmt.Println(err)
	fmt.Println(params)
	err=params.AddLine("LA")
	fmt.Println(err)
	fmt.Println(params)
	err=params.AddLine("5896")
	fmt.Println(err)
	fmt.Println(params)
	// Setting departureTime
	y := 2017
	m := time.March
	d := 17
	h := 13
	min := 24
	s := 0
	t := time.Date(y, m, d, h, min, s, 0, time.UTC)
	// set to a time
	params.SetDepartureTime(t)
	err=params.Validate()
	fmt.Println(err)
	fmt.Println(params)
	tString:=" "
	params.DepartureTime=&tString
	err=params.Validate()
	fmt.Println(err)
	fmt.Println(params)
	params.DepartureTime=GetTimeUnix(t)
	// Set hafasId
	params.HafasId=&hafasIdValid[0]
	err=params.Validate()
	fmt.Println(err)
	fmt.Println(params)
	params.HafasId=&hafasIdInValid[0]
	err=params.Validate()
	fmt.Println(err)
	fmt.Println(params)
	// Output:
	// <nil>
	// {<nil> <nil> <nil>}
	// <nil>
	// {[1] <nil> <nil>}
	// <nil>
	// {[1 4A] <nil> <nil>}
	// <nil>
	// {[1 4A 368] <nil> <nil>}
	// Invalid value for lines or lineId
	// {[1 4A 368] <nil> <nil>}
	// Invalid value for lines or lineId
	// {[1 4A 368] <nil> <nil>}
	// <nil>
	// {[1 4A 368] <nil> 1489757040}
	// Invalid value for time or departureTime
	// {[1 4A 368] <nil>  }
	// <nil>
	// {[1 4A 368] 1236 1489757040}
	// Invalid value for stationId or hafasId
	// {[1 4A 368] 12A6 1489757040}

}


// ## maps.go Parameter

func ExampleMapEntityParams(){
	params:=MapEntityParams{nil,nil}
	err:=params.Validate()
	fmt.Println(params)
	fmt.Println(err)
	pTS:=GetSupportedThumbnailSizes()
	pF:=GetSupportedFormats()
	params.ThumbnailSize=&pTS[3]
	err=params.Validate()
	fmt.Println(params)
	fmt.Println(err)
	ivTS:="111"
	params.ThumbnailSize=&ivTS
	err=params.Validate()
	fmt.Println(params)
	fmt.Println(err)
	params.Format=&pF[0]
	err=params.Validate()
	fmt.Println(params)
	fmt.Println(err)
	ivF:="jpg"
	params.Format=&ivF
	err=params.Validate()
	fmt.Println(params)
	fmt.Println(err)
	// Output:
	// {<nil> <nil>}
	// <nil>
	// {256 <nil>}
	// <nil>
	// {<nil> <nil>}
	// Invalid value for MapEntityParams.ThumbnailSize
	// {<nil> pdf}
	// <nil>
	// {<nil> <nil>}
	// Invalid value for MapEntityParams.Format
}

// ## stationMonitor.go Parameter

func ExampleJourneyParams(){
	params,err:=NewJourneyParams(hafasIdInValid[1])
	fmt.Println(params)
	fmt.Println(err)
	params,err=NewJourneyParams(hafasIdValid[1])
	fmt.Println(params)
	fmt.Println(err)
	y := 2017
	m := time.March
	d := 17
	h := 13
	min := 24
	s := 0
	t := time.Date(y, m, d, h, min, s, 0, time.UTC)
	params.SetTime(t)
	err=params.Validate()
	fmt.Println(params)
	fmt.Println(err)
	// for more examples for poles see ExampleJourneyInfoTranserParams
	err=params.SetPoles([]string{"5","13","1235"})
	fmt.Println(err)
	err=params.Validate()
	fmt.Println(params)
	fmt.Println(err)
	pM:=GetSupportedModes()
	params.Mode=pM[1]
	err=params.Validate()
	fmt.Println(params)
	fmt.Println(err)
	params.Mode="departure"
	fmt.Println(params)
	err=params.Validate()
	fmt.Println(params)
	fmt.Println(err)
	params.SetPlatformDetailNeeded(true)
	err=params.Validate()
	fmt.Println(params)
	fmt.Println(err)
	params.NeedPlatformDetail="ja"
	fmt.Println(params)
	err=params.Validate()
	fmt.Println(params)
	fmt.Println(err)
	err=params.SetTransportFilter("2365")
	fmt.Println(params)
	fmt.Println(err)
	err=params.SetTransportFilter("22")
	fmt.Println(params)
	fmt.Println(err)
	_,err,_=params.GetAsUrlParams()
	fmt.Println(err)
	isNotValidTimeParameter:="17.03.2017 14:2"
	params.Time=&isNotValidTimeParameter
	err=params.Validate()
	fmt.Println(params) // the invalid value will not be resetted. But it isn't possible to make a request.
	fmt.Println(err)
	_,err,_=params.GetAsUrlParams()
	fmt.Println(err)
	// Output:
	// {1125 null <nil> DEP false <nil>}
	// Invalid value for stationId or hafasId
	// {1648 null <nil> DEP false <nil>}
	// <nil>
	// {1648 2017-03-17+13:24 <nil> DEP false <nil>}
	// <nil>
	// Invalid value for poles
	// {1648 2017-03-17+13:24 [5 13] DEP false <nil>}
	// <nil>
	// {1648 2017-03-17+13:24 [5 13] ARR false <nil>}
	// <nil>
	// {1648 2017-03-17+13:24 [5 13] departure false <nil>}
	// {1648 2017-03-17+13:24 [5 13] DEP false <nil>}
	// Invalid value for JourneyParams.Mode
	// {1648 2017-03-17+13:24 [5 13] DEP true <nil>}
	// <nil>
	// {1648 2017-03-17+13:24 [5 13] DEP ja <nil>}
	// {1648 2017-03-17+13:24 [5 13] DEP false <nil>}
	// Invalid value for JourneyParams.NeedPlatformDetail
	// {1648 2017-03-17+13:24 [5 13] DEP false <nil>}
	// Invalid value for lines or lineId
	// {1648 2017-03-17+13:24 [5 13] DEP false 22}
	// <nil>
	// <nil>
	// {1648 17.03.2017 14:2 [5 13] DEP false 22}
	// Invalid value for time or departureTime
	// Invalid value for time or departureTime
}
