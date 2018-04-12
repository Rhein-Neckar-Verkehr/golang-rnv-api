// # Interfaces from rnv Start.Info API
//  - Entfallene Linien
//  	 - Fahrtverläufe
//  	 - Linienpaket
//  		- Linieninformation (sub Interface of Linienpaket)
//

package api
//
import (
	"time"
	"regexp"
	"strings"
)

// ## Parameter
//
//var UI_SOURCE_LINE="LINE"
// Parameter for Interface Fahrtverläufe
//
type LineJourneyParams struct{
    HafasId string
    LineId string
    StopIndex string
    TourType string
    TourId string
    Time string
}
//
func NewLineJourneyParamsFromDepartureAndStationId(d Departure,stationId string)(LineJourneyParams,error){
	time:=GetTimeFormatNow()
	params:=LineJourneyParams{stationId,d.LineLabel,"0", d.GetTourType(),d.TourId,*time}
	err:=params.Validate()
	return params,err
}
//
func (params *LineJourneyParams)SetStopIndex(index string)error{
	err:=validateIndex(index)
	if err==nil{
		params.StopIndex=index
	}
	return err
}
//
func (params *LineJourneyParams)SetTime(t time.Time){
	time:=GetTimeFormat(t)
	params.Time=*time
}
//
func (params *LineJourneyParams)GetTime()(*time.Time,error){
	err:=validateTimeAsString(params.Time)
	t:=time.Now() // TODO implement utility for this
	return &t,err
}
//
func (params *LineJourneyParams)Validate()error{
	err:=validateStationId(params.HafasId)
	if err!=nil {return err}
	params.LineId=checkLineId(params.LineId)
	err=validateLine(params.LineId)
	if err!=nil {return err}
	err=validateIndex(params.StopIndex)
	if err!=nil {return err}
	err=validateTourType(params.TourType)
	if err!=nil {
		tmp:=parseKindOfTourToTourType(params.TourType)
		err=validateTourType(tmp)
		if err!=nil {return err}else {params.TourType=tmp}
	}
	// validate tour id?? how?
	err=validateTimeAsString(params.Time)
	if err!=nil {return err}
	return nil
}
//
func (params *LineJourneyParams)GetAsUrlParams()(map[string][]string,error,bool){
//
	par :=make(map[string][]string)
	addTimeManually:=false
	err:=params.Validate()

	if err!=nil {return par,err,addTimeManually}

	par["hafasID"]=[]string{params.HafasId}
    par["lineID"]=[]string{params.LineId}
    par["stopIndex"]=[]string{params.StopIndex}
    par["tourType"]=[]string{params.TourType}
    par["tourID"]=[]string{params.TourId}

    if(params.Time=="null"){
       par["time"]=[]string{params.Time}
        addTimeManually=false
    }  else{
        addTimeManually=true
    }
	return par,err,addTimeManually
}

// Parameter for Interface Entfallene Linien
type CancelledLineParams struct{
    LineId *string
	DepartureTime *string
}
//
func (params *CancelledLineParams)SetDepartureTime(t time.Time){
	params.DepartureTime=GetTimeUnix(t)
}
//
// Getter for time?
//
func (params *CancelledLineParams)SetLineId(lineId string)error{
	err:=validateLine(lineId)
	if err==nil {
		params.LineId=&lineId
	}
	return err
}
// Manche Werte sind als Parameter, wie zur Filterung nicht zulässig. Diese korrigieren und ersetzen
func checkLineId(lineId string)string{
	// M [12345] Nachtbusse mit Leerzeichen
	valid:=regexp.MustCompile(`^M [12345]$`)
	res:=valid.MatchString(lineId)
		if res {
			return strings.Replace(lineId, " ", "", -1);
		}
	//[456]A Routen im Soll
	if(lineId=="4 / 4A"){
		return "4"
	}
	if(lineId=="6 / 6A"){
		return "6"
	}
	if(lineId=="5 / 5A"){
		return "5"
	}
	return lineId
}

//
func (params *CancelledLineParams)Validate()error{
	if params.LineId!=nil {
		temp:=checkLineId(*params.LineId)
		params.LineId=&temp
		err:=validateLine(*params.LineId)
		if err!=nil {
			params.LineId=nil
			return err
		}
	}
	if params.DepartureTime!=nil {
		err:=validateTimeAsUnixTimestamp(*params.DepartureTime)
		if err!=nil {
			params.DepartureTime=nil
			return err
		}
	}
	return nil
}
//
func (params CancelledLineParams)String()string{
	result:="{"
	if params.LineId==nil {result=result+"<nil> "}else {result=result+*params.LineId+" "}
	if params.DepartureTime==nil {result=result+"<nil>"}else {result=result+*params.DepartureTime}
	return result+"}"
}
//
func (params *CancelledLineParams)GetAsUrlParams()(map[string][]string,error){
	err:=params.Validate()
	par:=map[string][]string{}
	if err!=nil {return par,err}
	if params.DepartureTime!=nil{
		par["departureTime"]=[]string{*params.DepartureTime}
	}
	if params.LineId!=nil{
	par["line"]=[]string{*params.LineId}
	}
	return par,err
}

// ## Response Objects
//
// Response Object for Interface Linienpaket
//
type LineJourneyPackage struct {
	LineId string `json:"lineId"` // nicht vorhanden, laut doku schon
    //StopListIds []string `json:"stopListIds,omitempty"` // nicht vorhanden laut doku schon
	Ticker string `json:"ticker"`
    PredictedTimeList []string `json:"predictedTimeList"`//2014-10-23 20:04
    StationIds []string `json:"stationIDs"`
    LineIds []string `json:"lineIDs"`//Bad Dürkheim Bahnhof lineIDs großes d
    TimeList []string `json:"timeList"`//2014-10-23 20:04+0
    ValidFromIndex int `json:"validFromIndex"`//-1
	Directions []string  `json:"directionList"` // heißt anders, tatsächlich directionList
}
//
//
// Response Object for Interface Fahrtverläufe. Not all fieldsd are set. Need to use different struct, even it should be the same object
//
type LineJourney struct{
    Ticker string `json:"ticker"`
    PredictedTimeList []string `json:"predictedTimeList"`//2014-10-23 20:04
    StationIds []string `json:"stationsIDs"`
    LineIds []string `json:"lineIDs"`//Bad Dürkheim Bahnhof
    TimeList []string `json:"timeList"`//2014-10-23 20:04+0
    ValIdFromIndex int `json:"valIdFromIndex"`//-1
}
//
//
//
// Response Line Object for Interface Linieninformation
//
type Line struct{
    LineId string `json:"lineID"`
    LineType string `json:"lineType"`
    Hexcolor string `json:"hexcolor"`
    IconName string `json:"iconName"`
	ElementId int  `json:"elementID"`
	Icon	string `json:"icon"` // Kompatiblitätsgründe
}
// Response Object for Interface Entfallene Linien

type CanceledLineTransfer struct {
	// vorher LineID !!!, created gabs nicht.., author ist auch neu
	Id string `json:"Id"`
    Line string `json:"line"`
    ValidFrom uint64 `json:"validFrom"`
    ValidTo uint64 `json:"validTo"`
	Created uint64 `json:"created"`
	Author string `json:"author"`
    OperationDay string `json:"operationDay"` // Wo ist das hin??
    Canceled bool `json:"canceled"`
    Deleted bool `json:"deleted"`
}
//
func (c *CanceledLineTransfer)IsCurrentlyValid()bool{
	return IsInCurrentlyInTimeRange(c.ValidFrom,c.ValidTo)
}
//
func (c *CanceledLineTransfer)IsValidInFuture()bool{
	return IsTimeRangeInTheFuture(c.ValidFrom)
}

// ## Requests
//
// Request for Interface Entfallene Linien
//
func GetCancelledLine(params CancelledLineParams)([]CanceledLineTransfer,error){
    par,err:=params.GetAsUrlParams()
	var record []CanceledLineTransfer

	if err!=nil {return record,err}

    values:=getValues(par)
    urlV:=getUrlWithValues(SITE_URL+MODUL_URL+CANCELED_LINE_URL,values)

	err=makeRequest(urlV.String(),&record)

	return record, err
}

// Request for Interface Linienpaket
//
func GetLinesPackage()([]LineJourneyPackage,error){
    urlV:=getUrl(SITE_URL+MODUL_URL+LINE_PACKAGE_URL)
	var record []LineJourneyPackage
	err:=makeRequest(urlV.String(),&record)
	return record, err
}

// Request for  Interface Fahrtverläufe
//
func GetLines(params LineJourneyParams) (*LineJourney,error){
//
	par,err,addTimeManually:=params.GetAsUrlParams()
    if err!=nil {return nil,err}
    values:=getValues(par)
    urlV:=getUrlWithValues(SITE_URL+MODUL_URL+LINE_URL,values)
    var urlString string
    if addTimeManually {
        urlString=urlV.String()+"&time="+params.Time
    } else {
        urlString=urlV.String()
    }

	var record LineJourney
    errr:=makeRequest(urlString,&record)

	return &record, errr
}
// Request for Interface Linieninformation
//
func GetLineInfos()([]Line,error){
    urlV:=getUrl(SITE_URL+MODUL_URL+LINE_ALL_URL)
	var record []Line
	err:=makeRequest(urlV.String(),&record)
	return record, err
}
// ## Other
//
// Ergebnis Linienpaket. Ist nicht das gleiche. Liefert Fehler..
// <lineJourney>
// <lineId>1</lineId>
// <lineIDs><stopListIds>2525</stopListIds><stopListIds>2532</stopListIds><stopListIds>2553</stopListIds><stopListIds>2543</stopListIds><stopListIds>2544</stopListIds><stopListIds>2550</stopListIds><stopListIds>2483</stopListIds><stopListIds>2492</stopListIds><stopListIds>2487</stopListIds><stopListIds>2481</stopListIds><stopListIds>2506</stopListIds><stopListIds>2494</stopListIds><stopListIds>2499</stopListIds><stopListIds>2564</stopListIds><stopListIds>2516</stopListIds><stopListIds>2466</stopListIds><stopListIds>2417</stopListIds><stopListIds>2471</stopListIds><stopListIds>2462</stopListIds><stopListIds>2451</stopListIds><stopListIds>2438</stopListIds><stopListIds>2447</stopListIds><stopListIds>5548</stopListIds><stopListIds>2410</stopListIds><stopListIds>2449</stopListIds><stopListIds>2421</stopListIds><stopListIds>2472</stopListIds><stopListIds>2386</stopListIds><stopListIds>2615</stopListIds><stopListIds>2626</stopListIds><stopListIds>5554</stopListIds><stopListIds>2584</stopListIds><stopListIds>5553</stopListIds></lineIDs>
// <validFromIndex>0</validFromIndex>
// </lineJourney>
