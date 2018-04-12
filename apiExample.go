//
// # Examples for  rnv Start.Info API (1.12).
package api
//
import (
//
    "log"
    "time"
	"math/rand"
	"strconv"
)
//
var example_linePackage []LineJourneyPackage
var example_stationPackage *StationPackage
var nullTime string
//
func loadPackages()error{
	var err error
	nullTime="null"
	example_linePackage,err=GetLinesPackage()
    if err!=nil{
        log.Println(err)
        return err
    }
//
	example_stationPackage,err=GetStationsPackage("1")
	return err
}
//
// ### Zufallsfunktionen
// Zufällige Linien oder Haltestellen Ids erhalten aus dem aktuellen Linien bzw. Haltestellenpaket.
//
func getRandomLine()string{
	index:=rand.Intn(len(example_linePackage)-1)
	return  example_linePackage[index].LineId
}
//
func getRandomLinesObject()LineJourneyPackage{
	index:=rand.Intn(len(example_linePackage)-1)
	return  example_linePackage[index]
}
//
func getRandomHafasId()string{
	index:=rand.Intn(len(example_stationPackage.Stations)-1)
	return example_stationPackage.Stations[index].HafasId
}
//
func getRandomStationsObject() *Station {
index:=rand.Intn(len(example_stationPackage.Stations)-1)
	return example_stationPackage.Stations[index]
}
//
func  Example(){
//
// Für Tests im RNV Netz müssen diese Proyeinstellungen genutzt werden
//
//	os.Setenv("HTTP_PROXY", "192.168.190.145:8080")
//	os.Setenv("HTTPS_PROXY", "192.168.190.145:8080")
//
	err:=loadPackages()
	if (err!=nil){
		log.Println("Failed to load Haltestellen- and Linienpaket")
		log.Println(err)
		return
	}
// #### Haltestellenmonitor, Haltepunkte from StationPackage
// Request Haltestellenmonitor
//
	station:=getRandomStationsObject()
	monitor,_:=station.GetStationMonitorTimeNow()
//
// Request Fahrtverläufe for first Departure Object
//
	line:=monitor.ListOfDepartures[0]
	lines,_:=line.GetLineJourney(station.HafasId)
	log.Println(lines)
//
// Request News if available for this Journey
//
	news,_:=line.GetNewsIfAvailable()
	log.Println(news)
//
// Request active Haltepunkte
//
	details,_:=station.GetActiveStationDetails()
//
// Request Haltestellenmonitor for all Haltepunkte for this Station
//
	for _,val:= range details {
		log.Println("Abfahrtsmonitor für diesen Haltepunkt" + val.GetPoleAsString())
		monitor,_=val.GetStationMonitorTimeNow(station.HafasId)
		PrintHaltestellenmonitor(monitor)
	}
//
// ####  Haltestellenmonitor, Haltepunkte from Parameter hafasId/stationId
//
	HafasId:=getRandomHafasId()
//
// Request Haltestellenmonitor
//
	stationMonitorQueryParameter:=JourneyParams{HafasId,&nullTime,nil,GetSupportedModes()[0],GetSupportedNeedPlatformDetail()[0],nil}
	monitor,_=GetStationMonitor(stationMonitorQueryParameter)
	PrintHaltestellenmonitor(monitor)
//
// Request Fahrtverläufe for first Departure Object
//
	departureObject:=monitor.ListOfDepartures[0]
	tt:=GetTimeFormatNow()
	lines,_=GetLines(LineJourneyParams{HafasId,departureObject.LineLabel,departureObject.PositionInTour, departureObject.KindOfTour,departureObject.TourId,*tt})
//
// Request Haltepunkte
//
	stationDetails,_:=GetStationDetail(HafasId)
//
// Request Haltestellenmonitor several Haltepunkte for the next day
// Setting Parameters for Haltestellenmonitor
//
	listOfPoles:=[]string{}
	for i:=0;i<len(stationDetails)-1;i++ {
		p:=strconv.Itoa(stationDetails[i].Pole)
		listOfPoles=append(listOfPoles,p)
	}
	stationMonitorQueryParameter.SetPlatformDetailNeeded(true)
	stationMonitorQueryParameter.SetPoles(listOfPoles)
//
// apply other time value then "null"
//
	t:=time.Now()
	t.Add(time.Hour * 24)
	stationMonitorQueryParameter.SetTime(t)
	monitor,_=GetStationMonitor(stationMonitorQueryParameter)
	PrintHaltestellenmonitor(monitor)
//
// #### Liniennetzpläne
//
	format:=GetSupportedFormats()[0]
	thumbSize:=GetSupportedThumbnailSizes()[0]
	maps,_:=GetMaps(MapEntityParams{&thumbSize,&format})
	log.Println(maps)
//
// #### News
//
	news,_=GetNews()
	log.Println(news)
//
    for _,n:= range news{
		log.Println(n.AffectsGivenLine(getRandomLine()))
    }
//
// #### Ticker
//
	tick,_:=GetTickerForLines([]string{getRandomLine(),getRandomLine(),getRandomLine(),getRandomLine(),getRandomLine()})
	log.Println(tick)
//
// #### Haltestelleninformationen
	var Lines []string
    Lines=append(Lines,getRandomLine())
    HafasId =getRandomHafasId()
//
	stationInfo,_:=GetStationInfo(StationInfoTransferParams{&Lines,&HafasId,nil})
	// staitonInfo,_:=GetStationInfo(StationInfoTransferParams{&Lines,nil,nil})
	for _,in := range *stationInfo{
        log.Println(in)
		log.Println(in.IsCurrentlyValid())
		log.Println(in.IsValidInFuture())
		log.Println(in.IsCurrentlyDisplayed())
		log.Println(in.AffectsGivenStation(HafasId))
    }
// #### Linieninformation
//
	lineInfo,_:=GetLineInfo(JourneyInfoTranserParams{nil,nil,GetTimeUnixNow()})
//
    for _,in := range *lineInfo{
        log.Println(in)
		log.Println(in.IsCurrentlyValid())
		log.Println(in.IsValidInFuture())
    }
//
// #### Entfallene Linien
//
	lineObjectFromLinesPackage:=getRandomLinesObject()
	cLine,_:=GetCancelledLine(CancelledLineParams{&lineObjectFromLinesPackage.LineId,GetTimeUnixNow()})
//
    for _,cl:= range cLine {
		log.Println(cl)
		log.Println(cl.IsCurrentlyValid())
		log.Println(cl.IsValidInFuture())
	}
//
// #### Linieninformation
//
	lineInfos,err:=GetLineInfos()
	log.Println(lineInfos)
//
// #### Aktualisierung
//
	time:=GetTimeFormatNow()
    updatesPackage,_:=GetUpdates(UpdateParams{RNV_REGION_ID_PARAMS,time,time,time})
	log.Println(updatesPackage)
}
//
//
// ## Ausgabe des Haltestellenmonitors
//
func PrintHaltestellenmonitor(monitor *Journey){
	log.Println(""+monitor.Label + " (" + monitor.ShortLabel+") Abfragezeit:"+monitor.Time)
	log.Println(monitor.Ticker)
//
	log.Println("Linie\t|\t ist Bus\t|\t ist Bahn\t|\t Ziel\t\t|\t  Abfahrt (in/um)\t\t|\t mit Prognose?\t|\t Steig/Gleis")
	for _,v:= range monitor.ListOfDepartures{
		log.Println(v.LineLabel +"\t\t|\t"+strconv.FormatBool(v.IsBus())+"\t\t|\t"+strconv.FormatBool(v.IsTram())+"\t\t|\t"+v.Direction+"\t|\t"+v.DifferenceTime+" min / "+v.Time+"\t\t\t|\t"+strconv.FormatBool(v.IsRealtimeData())+"\t\t|\t"+v.Platform)
	}
}
//
//
//
