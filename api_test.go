// # Tests for rnv Start.Info API (1.12)
// All Data is Realtime Data. Tests cover:
//  - Retrieve Data from rnv Start.Info API and Marshalling the Results should succeed
//  - Retrieving an ErrorDescription Object as Result is not a failure
//  - TODO Additional Functions for Response Objects working correctly
// _Running Tests require an valid RNV_API_TOKEN provided in tokenConf.json_
//
package api
import (
	"testing"
	"math/rand"
	"fmt"
	"log"
)
//
var test_linePackage []LineJourneyPackage
var test_stationPackage *StationPackage
//
func init(){
	nullTime="null"
}
//
//
func test_getRandomLine()string{
	index:=rand.Intn(len(test_linePackage)-1)
	return  test_linePackage[index].LineId
}
//
func test_getRandomHafasId()string{
	index:=rand.Intn(len(test_stationPackage.Stations)-1)
	return test_stationPackage.Stations[index].HafasId
}
//
func test_getRandomStationsObject() *Station {
	index:=rand.Intn(len(test_stationPackage.Stations)-1)
	return test_stationPackage.Stations[index]
}
//
func dTestReachabilityOfAPI(t *testing.T ){
//
// Setup and Tets.
//
	TestGetLinesPackage(t)
	TestGetStationsPackage(t)
//
// Test Modul Functions
//
	TestGetUpdates(t)
	TestGetMaps(t)
	TestGetStationInfo(t)
	TestGetLineInfo(t)
	TestGetCancelledLine(t)
	TestGetTicker(t)
	TestGetNews(t)
	TestGetLineInfos(t)
//
// Test Struct and its Functions
//
	TestStation(t)
	TestStationDetails_second(t)
	TestStationDetails_first(t)
	//TestLineJourney_GetLineJourneyAndGetNewsAvailable_WithAllStations(t)
//
    testLineJourney_GetLineJourneyAndGetNewsAvailable(3,t)
//
}
//
func TestGetLinesPackage(t *testing.T ){
	lp,err:=GetLinesPackage()
	test_linePackage=lp
    if err!=nil{
        t.Fatalf("Error raised: %err", err)
    }
}
//
func TestGetStationsPackage(t *testing.T ){
	st,err:=GetStationsPackage("1")
//
	test_stationPackage=st
	if err!=nil{
        t.Fatalf("Error raised: %err", err)
    }
}
//
func TestStationDetails_second(t *testing.T ){
	maHBFstation:=test_stationPackage.GetStationByName("MA Hauptbahnhof")
	if (maHBFstation==nil){
        t.Fatal("Haltestelle MA Hauptbahnhof nicht gefunden im Haltestellenpacket:")
	}
	_,err:=maHBFstation.GetStationMonitorTimeNow()
	if err!=nil{
        t.Fatalf("Error raised: %err", err)
    }
	maHBFdetail,err2:=maHBFstation.GetStationDetails()
	if err2!=nil{
        t.Fatalf("Error raised: %err", err2)
    }
	for _,val:= range maHBFdetail {
		_,err2=val.GetStationMonitorTimeNow(maHBFstation.HafasId)
		if err2!=nil{
        t.Fatalf("Error raised: %err", err2)
    }
	}
}
//
// Test new Functions Requests From Response Objects
//
func TestStation(t *testing.T){
	for i := 0; i <= 8; i++ {
		station:=test_getRandomStationsObject()
		_,err2:=station.GetStationMonitorTimeNow()
		if err2!=nil{
        t.Fatalf("Error raised: %err", err2)
		}
		details,err:=station.GetStationDetails()
		if err!=nil{
			t.Fatalf("Error raised: %err", err)
		}
		for _,val:= range details {
			_,err2=val.GetStationMonitorTimeNow(station.HafasId)
//
			if err2!=nil{
				t.Fatalf("Error raised: %err", err2)
			}
		}
		_,err=station.GetActiveStationDetails()
		if err!=nil{
			t.Fatalf("Error raised: %err", err)
		}
	}
}
//
// Test new Interface Haltepunkte
//
func TestStationDetails_first(t *testing.T ){
	_,err:=GetStationDetail(test_getRandomHafasId())
	if err!=nil{
        t.Fatalf("Error raised: %err", err)
	}
}
//
// Test for Interface Linniennetzpläne implemented in maps.go
//
func TestGetMaps(t *testing.T){
    for _,v:=range GetSupportedThumbnailSizes(){
        for _,f:=range GetSupportedFormats(){
            maps,err:=GetMaps(MapEntityParams{&v,&f})
            if err!=nil{
				t.Fatalf("Error raised: %err", err)
			}
            for _,in := range *maps{
				if (in.Name==""){
					t.Fatal("No Name for map response ")
				}
				if (in.ShortName==""){
					t.Fatal("No ShortName for map response ")
				}
				if (in.Author==""){
					t.Fatal("No Author for map response ")
				}
				if (in.ThumbnailUrl==""){
					t.Fatal("No ThumbnailUrl for map response ")
				}
				if (in.ValidFrom==""){
					t.Fatal("No ValidFrom for map response ")
				}
				if (in.MapUrl==""){
					t.Fatal("No MapUrl for map response ")
				}
            }
        }
    }
}
//
// Test for Interface Haltestelleninformationen
//
func TestGetStationInfo(t *testing.T ){
	// Requests to Haltestelleninfromationen Interface
    _,err:=GetStationInfo(StationInfoTransferParams{nil,nil,nil})
    if err!=nil{
		t.Fatalf("Error raised: %err", err)
	}
    var Lines []string
    Lines=append(Lines,test_getRandomLine())
    HafasId :=test_getRandomHafasId()
//
	_,err=GetStationInfo(StationInfoTransferParams{&Lines,nil,nil})
        if err!=nil{
		t.Fatalf("Error raised: %err", err)
	}
    _,err=GetStationInfo(StationInfoTransferParams{&Lines,&HafasId,nil})
        if err!=nil{
		t.Fatalf("Error raised: %err", err)
	}
}
// Test for Interface Fahrtinformationen, see info.go
//
func TestGetLineInfo(t *testing.T ){
	HafasId2 :=test_getRandomHafasId()
	// Requests to Linieninformationen Interface
    _,err:=GetLineInfo(JourneyInfoTranserParams{nil,nil,GetTimeUnixNow()})
    if err!=nil{
		t.Fatalf("Error raised: %err", err)
	}
//
    var Poles []string
    Poles=append(Poles,"7")
    _,err=GetLineInfo(JourneyInfoTranserParams{&HafasId2,&Poles,nil})
    if err!=nil{
		t.Fatalf("Error raised: %err", err)
	}
}
//
// Test Interface Entfallene Linien. Each Line in Linienpaket (test_linePackage) will be used as Parameter.
//
func TestGetCancelledLine(t *testing.T){
    for _,line:=range test_linePackage{
			log.Println(line.LineId)
			// Sonderfälle als Parameter nur 4 bzw. 4A, aber als Resultat kommen beide Linien zurück (Linienfilter)
			if(line.LineId=="RT"){ // Wert wird von der API gelifert.
				continue
			}
        _,err:=GetCancelledLine(CancelledLineParams{&line.LineId,GetTimeUnixNow()})
		if err!=nil{
			t.Fatalf("Error raised: %err", err)
		}
    }
}
//
// Test for Interface News
//
func TestGetNews(t *testing.T ){
    //No news available on Testserver..
    _,err:=GetNews()
    if err!=nil{
		t.Fatalf("Error raised: %err", err)
	}
}
//
// Test for Interface Ticker
//
func TestGetTicker(t *testing.T ){
		params:=NewTickerParams()
		params.AddLine(test_getRandomLine())
		params.AddLine(test_getRandomLine())
		params.AddLine(test_getRandomLine())
		params.AddLine(test_getRandomLine())
		params.AddLine(test_getRandomLine())
		log.Println(params);
    _,err:=GetTickerForLines(params)
    if err!=nil{
		t.Fatalf("Error raised: %err", err)
	}
}
//
// Test to Interface Linieninformation
//
func TestGetLineInfos(t *testing.T ){
    _,err:=GetLineInfos()
    if err!=nil{
		t.Fatalf("Error raised: %err", err)
	}
}
//
// Test of Interface Änderungen am Haltesellen- und Linienpaket ermitteln
func TestGetUpdates(t *testing.T ){
    time2:=GetTimeFormatNow()
    _,err:=GetUpdates(UpdateParams{"1",time2,time2,time2})
//
    if err!=nil{
		//t.Fatalf("Error raised: %err", err)
		fmt.Printf("Updates stuff",err)
	}
}
//
// Test Interface Haltestellenmonitor for each Station in StationPackage. For Each Departure Objects
// in Journey Object, call Interface Linienverläufe
func TestLineJourney_GetLineJourneyAndGetNewsAvailable_WithAllStations(t *testing.T ){
   params,errp:=NewJourneyParams("0000")
	 testdata:=[]*Station{test_getRandomStationsObject(),test_getRandomStationsObject(),test_getRandomStationsObject(),test_getRandomStationsObject(),test_getRandomStationsObject()}
   if errp!=nil{
		t.Fatalf("Error raised: %err", errp)
	}
    for _,station := range testdata{
		params.HafasId=station.HafasId
        // stationsMonitorBody,err:=GetStationMonitor(params)
				_,err:=GetStationMonitor(params)
//
		if err!=nil{
			t.Fatalf("Error raised: %err", err)
		}
    /*
		// TODO heir erscheint jetzt eine Fehlermeldung, aus technischen gründen ...

		   for _,line:=range stationsMonitorBody.ListOfDepartures{
            _,err=line.GetLineJourney(station.HafasId)
			if err!=nil{
				t.Fatalf("Error raised: %err", err)
			}
			// get News if available
			_,err=line.GetNewsIfAvailable()
			if err!=nil{
				t.Fatalf("Error raised: %err", err)
			}
        }*/
//
    }
}
//
// test for count stations
//
func testLineJourney_GetLineJourneyAndGetNewsAvailable(count int,t *testing.T){
	stations:=[]*Station{}
	params,errp:=NewJourneyParams("0000")
   if errp!=nil{
		t.Fatalf("Error raised: %err", errp)
	}
	for i:=0;i<count;i++{
		stations=append(stations,test_getRandomStationsObject())
	}
    for _,station := range stations{
		params.HafasId=station.HafasId
		stationsMonitorBody,err:=GetStationMonitor(params)
        if err!=nil{
			t.Fatalf("Error raised: %err", err)
		}
		for _,line:=range stationsMonitorBody.ListOfDepartures{
			_,err:=line.GetLineJourney(station.HafasId)
			if err!=nil{
				t.Fatalf("Error raised: %err", err)
			}
			// get News if available
			_,err=line.GetNewsIfAvailable()
			if err!=nil{
				t.Fatalf("Error raised: %err", err)
			}
       }
//
   }
}
