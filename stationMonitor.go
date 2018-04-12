// # Interfaces from rnv Start.Info API
//  	- Haltestellenmonitor
// 
package api
import (
	"io/ioutil"
	"time"
	"strings"
	"strconv"
	"errors"
	"log"
)
// 

const TICKER_START_AND_END_IDENTIFIER string="***"
// 
// 
// Enumerations for Response Object Attributes for this Interface

const KIND_OF_TOUR_PLANNED string="452"
const KIND_OF_TOUR_REALTIME_DATA string="454AUS"
const KIND_OF_TOUR_PLANNED_DAILY string="454REFAUS"
// 
const TRANSPORTATION_STRAB string="STRAB"
const TRANSPORTATION_KOM string="KOM"
const TRANSPORTATION_WEBU string="WEBU"
// 
const STATUS_OK string="OK"
const STATUS_CANCELLED string="CANCELLED"

// ## Parameter
// 
// TODO move this stuff values for Fahrtverlauf

var index string="0"	
// 
var time_default string="null"
// 
var supportedModes =[2]string{"DEP","ARR"} //default is DEP
// 
func GetSupportedModes()[2]string{return supportedModes}
const ERROR_MODE string=ERROR_START+" JourneyParams.Mode"
const CATCH_ERROR_MSG_MODE=CATCH_ERROR_MSG_START+" JourneyParams.Mode: DEP"
// 
var supportedNeedPlatformDetail =booleanA // default is false
func GetSupportedNeedPlatformDetail()[2]string{return supportedNeedPlatformDetail}
const ERROR_NEED_PLATFORM_DETAIL string=ERROR_START+" JourneyParams.NeedPlatformDetail"
const CATCH_ERROR_MSG_NEED_PLATFORM_DETAIL=CATCH_ERROR_MSG_START+" JourneyParams.NeedPlatformDetail: false"

// 

type JourneyParams struct{
    // NOTE hafasId=116$-1&transportFilter=4&time=null&uiSource=LINE illegal api access
    HafasId string // required
    Time *string    
    // NOTE UiSource *string field from illegal api access
    Poles *[]string
    Mode string // use default values
    NeedPlatformDetail string // use default values
    TransportFilter *string 
}
// 
// 
// 
func NewJourneyParams(hafasId string)(JourneyParams,error){
	err:=validateStationId(hafasId)
	if err!=nil {hafasId="1125"}
	return JourneyParams{hafasId,&time_default,nil,supportedModes[0],supportedNeedPlatformDetail[1],nil},err
}
// 

func (params *JourneyParams)SetTransportFilter(lineLabel string)error{
	err:=validateLine(lineLabel)
	if err==nil {params.TransportFilter=&lineLabel}
	return err
}
// 

func (params JourneyParams)String()string{
	result:="{"+params.HafasId+" "
	if params.Time!=nil {result=result+*params.Time+" "} else {result=result+"<nil> "}
	if params.Poles!=nil {result=result+arrayToString(*params.Poles)+" "} else {result=result+"<nil> "}
	result=result+params.Mode+" "+params.NeedPlatformDetail+" " 
	if params.TransportFilter!=nil {result=result+*params.TransportFilter} else {result=result+"<nil>"}
	return result+"}"
}
// 

func (params *JourneyParams)Validate()error{
	err:=validateStationId(params.HafasId)
	if err!=nil{
		return err
	}
	if params.Mode != supportedModes[0] && params.Mode != supportedModes[1]{
		params.Mode=supportedModes[0]
		return errors.New(ERROR_MODE)
	} 
	if params.NeedPlatformDetail != supportedNeedPlatformDetail[0] && params.NeedPlatformDetail != supportedNeedPlatformDetail[1]{
		params.NeedPlatformDetail=supportedNeedPlatformDetail[1]
		return errors.New(ERROR_NEED_PLATFORM_DETAIL)
	} 
	if params.Time==nil {
		return nil
	}
	if *params.Time==time_default {
		return nil
	}
	return validateTimeAsString(*params.Time)
}
// 

func (params *JourneyParams)GetAsUrlParams()(map[string][]string,error,bool){
	err:=params.Validate()
	par :=make(map[string][]string)
	// catch Mode and needPlatformDetail errors
	if err==errors.New(ERROR_MODE) {
		log.Println(CATCH_ERROR_MSG_MODE)
		params.Mode=supportedModes[0]
		err=nil
	}
	if err==errors.New(ERROR_NEED_PLATFORM_DETAIL) {
		log.Println(CATCH_ERROR_MSG_NEED_PLATFORM_DETAIL)
		params.NeedPlatformDetail=supportedNeedPlatformDetail[1]
		err=nil
	}
	if (err!=nil){
		return par,err,false
	}
	var addTimeManually bool
    par["hafasID"]=[]string{params.HafasId}
    if(*params.Time==time_default){
       par["time"]=[]string{*params.Time}  
        addTimeManually=false
    } else if(params.Time==nil){
       par["time"]=[]string{time_default}  
        addTimeManually=false
    } else{
        
        addTimeManually=true
    }
    par["mode"]=[]string{params.Mode}
    
    if(params.Poles!=nil){
        par["poles"]=[]string{GetLinesSeparated(*params.Poles)}
    }
    par["needPlatformDetail"]=[]string{params.NeedPlatformDetail}
    
    if (params.TransportFilter!=nil){
        par["transportFilter"]=[]string{*params.TransportFilter}
    }
	return par,err,addTimeManually
}

// Function to set time correctly

func (params *JourneyParams)SetTime(t time.Time){
	params.Time=GetTimeFormat(t)
}

// If needed sets it to correct string value. Default is false

func (params *JourneyParams)SetPlatformDetailNeeded(platformDetail bool){
	if platformDetail {
		params.NeedPlatformDetail=supportedNeedPlatformDetail[0]
		return
	}
	params.NeedPlatformDetail=supportedNeedPlatformDetail[1]
}

// Sets Poles from multiple poles (as string) array to correct string parameter

func (params *JourneyParams)SetPoles(poles []string)error{
	var err error
	for _,v:=range poles {
		err=validatePole(v)	
		if err!=nil{return err}
		params.AddPole(v)
	}
	//temp:=GetLinesSeparated(poles)// TODO rename GetLinesSeparated, to something else.
	//params.Poles=&temp
	return nil
}
// 
// 
func (params *JourneyParams)AddPole(pole string)error{
	err:=validatePole(pole)	
	if err!=nil{return err}
	if params.Poles==nil{
		ne:=[]string{pole}
		params.Poles=&ne
	} else {
		ne:=append(*params.Poles,pole)
		params.Poles=&ne
	}
	return nil
}
// 
// 
// 
// ## Response Objects
// 
// Response Object Interface Haltestellenmonitor

type Journey struct {
	ShortLabel string `json:"shortLabel"`
	ProjectedTime string `json:"projectedTime"`
	Label	string `json:"label"`
	Icon	string `json:"icon"`
	Color	string `json:"color"`
	OtherProjectedTimes []string `json:"otherProjectedTimes"` // Kompatibilitätsgründe
	OtherTimes	[]string `json:"otherTimes"` // Kompatibilitätsgründe
	PastRequestText	string `json:"pastRequestText"`
	UpdateIterations	string `json:"updateIterations"` // Kompatibilitätsgründe
	UpdateTime	string `json:"updateTime"` // Kompatibilitätsgründe
    Ticker string `json:"ticker"`
    Time string `json:"time"`
    ListOfDepartures []*Departure `json:"listOfDepartures"`
}
// 

func (j *Journey)SplitMultipleTicker()[]string{
	result:=[]string{}
	if j.Ticker!=""{
		result:= strings.Split(j.Ticker,TICKER_START_AND_END_IDENTIFIER+TICKER_START_AND_END_IDENTIFIER)
		for _,v:= range result{
			if !strings.HasPrefix(v, TICKER_START_AND_END_IDENTIFIER){
				v=TICKER_START_AND_END_IDENTIFIER+v
			}
			if !strings.HasSuffix(v, TICKER_START_AND_END_IDENTIFIER){
				v=v+TICKER_START_AND_END_IDENTIFIER
			}
		}
	}
	return result
}

// Response Object for Interface Haltestellenmonitor

type Departure struct {
    Time string `json:"time"`
    Status string `json:"status"`
    Direction string `json:"direction"`
    Platform string `json:"platform"`
    Transportation string `json:"transportation"`
    TourId string `json:"tourId"`
    KindOfTour string `json:"kindOfTour"`
    PositionInTour string `json:"positionInTour"`
    StatusNote string `json:"statusNote"`
    LineId string `json:"lineId"`
    LineLabel string `json:"lineLabel"`
    DifferenceTime string `json:"differenceTime"`
    ForeignLine string `json:"foreignLine"`
    NewsAvailable string `json:"newsAvailable"`
}

// Check if Departure contains Realtime Data

func (d *Departure)IsRealtimeData()bool{
	return d.KindOfTour==KIND_OF_TOUR_REALTIME_DATA
}

// Get right pvalue for Paremter LineJourneyParams.TourType

func (d *Departure)GetTourType()string{
	return parseKindOfTourToTourType(d.KindOfTour)
}

// Is it a rnv Line or not

func (d *Departure)IsForeignLine()bool{
	return d.ForeignLine==TRUE
}
// Is a Bus

func (d *Departure)IsBus()bool{
	//return d.Transportation==TRANSPORTATION_KOM || d.Transportation==TRANSPORTATION_WEBU
	return !d.IsTram()
}
// Is a Tram

func (d *Departure)IsTram()bool{
	return d.Transportation==TRANSPORTATION_STRAB
}
// Returns the Extended Route Type of GTFS see https://support.google.com/transitpartners/answer/3520902
// rnv uses: 704  for Bus, 900  for Tram

func (d *Departure)GetGTFSExtendedRouteType()int{
	if d.IsTram(){
		return 900
	}
	return 704
}

// Is Departure cancelled

func (d *Departure)IsCancelled()bool{
	return d.Status==STATUS_CANCELLED
}
// 
func (d *Departure)AreNewsAvailable()bool{
	return d.NewsAvailable==TRUE
}

// Return value nil indicates no realtime data.

func (d *Departure)GetDelayIfRealtime()*int{
	if d.IsRealtimeData(){
		// delay is in Time string hh:mm+delay
		res:=strings.Split(d.Time,"+")
		if (len(res)==2){
			toInt,_:=strconv.Atoi(res[1])
			return &toInt
		}
	}
	return nil
}
// 
func (d *Departure)GetRNVGTFSRouteId()string{
	//RNV001_RNV001A
	tmp:=strings.Split(d.LineId,"_")
	result:=strings.Replace(tmp[0],"RNV","",-1)
	toInt,_:=strconv.Atoi(result)
	return strconv.Itoa(toInt)
}
// 
// 
// ## Requests
// 
// Request for Interface Haltestellenmonitor

func GetStationMonitor(params JourneyParams) (*Journey,error){
    urlString,err:=GetStationMonitorUrlAsString(params)   
	var record Journey
	log.Println(urlString)
	log.Println(err)
	if err!= nil {
		return &record,err
	}
    err=makeRequest(urlString,&record)
	log.Println(err)
	log.Println(record)
	return &record, err
}
// ## Request From Response Object

func (d *Departure)GetNewsIfAvailable()([]NewsEntry,error){
	result:=[]NewsEntry{}
	if d.AreNewsAvailable(){		
		news,err:=GetNews();
		if err!= nil {
			return result,nil
		}
		tick,err2:=GetTickerForLines([]string{d.LineLabel})
		if err2!= nil {
			return result,nil
		}
		// TODO test if LineLabel is value in seperated Line in newsEntry
		for _,v:= range news {
			if !v.IsOldNews && v.AffectsGivenLine(d.LineLabel){
				tick=append(tick,v)
			}
		}
		return tick,nil
	}
	return result,nil
}

// Gets LineJourney from Interface Farhtverläufe. Parameter TourId and Time will be overwritten with Departure Object values

func (d *Departure)GetLineJourneyWithAdditionalParams(params LineJourneyParams)(*LineJourney,error){
	params.TourId=d.TourId
	//params.LineId=&d.LineId
	time:=GetTimeFormatNow()
	params.Time=*time
	
	//params.TourType= d.GetTourType()
	return GetLines(params)
	
}

// Gets LineJourney from Interface Farhtverläufe. Only Parameter TourId and Time will applied.

func (d *Departure)GetLineJourney( stationId string)(*LineJourney,error){	
	time:=GetTimeFormatNow()
	return GetLines(LineJourneyParams{stationId,d.LineLabel,index, d.GetTourType(),d.TourId,*time})	
}

// ## Other
// These functions are needed to perform asynchronous Requests
// Request for Interface Haltestellenmonitor but returns the result as []byte Array (not Unmarshalled)

func GetStationMonitorForReverse(params JourneyParams) ([]byte,error){
 	urlString,err:=GetStationMonitorUrlAsString(params)
	if err!=nil{return nil,err}
    return getContent(urlString)
  
}

// Unmarshalls result from HTTP Response to Jouney Object 

func UnmarshalStationMonitorResponse(res *HttpResponse)(*Journey,error){
	body, err := ioutil.ReadAll(res.Response.Body)
	if err != nil {
		return nil, err
	}
	var record Journey
	err=parseResponseObject(body,record)
	return &record, err
}

// utiliy for Url for Interface Haltestellenmonitor

func GetStationMonitorUrlAsString(params JourneyParams)(string,error){
	par,err,addTimeManually :=params.GetAsUrlParams()    
	if err != nil {
		return "", err
	}
    values:=getValues(par)
    urlV:=getUrlWithValues(SITE_URL+MODUL_URL+DEPARTURES_URL,values)
    var urlString string
    if addTimeManually {        
        urlString=urlV.String()+"&time="+*params.Time
    } else {
        urlString=urlV.String()
    }
	return urlString,err
}