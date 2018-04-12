// # Interfaces from rnv Start.Info API
//  - Haltestellenpaket
//     - Änderungen am Haltestellen- und Linienpaket ermitteln
//     - Haltepunkte
//   
package api
import (
	"strconv"
)

// Mon Jan 2 15:04:05 MST 2006

const TIMESEPARATOR string="$"

// update?regionId=1&time=2014-07-08+13:40$2014-07-08+13:40$2011-11-11+11:11

const RNV_REGION_ID_PARAMS string="1"
// ## Parameter
// 
// Parameter for Interface 'Änderungen am Haltestellen- und Linienpaket ermitteln'

type UpdateParams struct{
    RegionId string
    //Time in Pattern yyyy-MM-dd+HH:mm separated by $
    LastUpdateStationsTime *string
    LastUpdateLinesTime *string
    UnusedTime *string
}

//  TODO implement validate, setParams stuff and NewUpdateParams. Understand the right parameters..
// ## Response Objects 
// 
// Response Object for Interface Haltestellenpaket

type StationPackage struct {
    Name string `json:"name"`
	RegionId int `json:"regionID"` // CHANGES stimmt nicht in Dokumentation, dort string, tatsächlich int
	ElementId int `json:"elementID"` // CHANGES stimmt nicht in Dokumentation, dort string, tatsächlich int
    Stations []*Station `json:"stations"`
}

// Get Station by Station Name 

func (stp *StationPackage)GetStationByName(name string)*Station{
	for _,v:= range stp.Stations {
		if v.LongName==name {
			return v
		}
	}
	return nil
}

// Response Object for Interface 'Änderungen am Haltestellen- und Linienpaket ermitteln'

type Update struct {
    Action string `json:"action"`
    Critical bool `json:"critical"`
    Description string `json:"description"`
    Element string `json:"element"`
    UpdateElementId int `json:"updateElementId"`
    ElementId int `json:"elementId"`
}

// Respone Object for Interface Haltepunkte

type StationDetail struct {
	Id int `json:"id"` // CHANGES stimmt nicht in Dokumentation, dort string, tatsächlich int
	Pole int `json:"pole"` // CHANGES stimmt nicht in Dokumentation, dort string, tatsächlich int
	PlatformText string `json:"platformText"`
	Platform string `json:"platform"`
	Longitude string `json:"longitude"`
	Latitude string `json:"latitude"`
	LeadTime int `json:"leadTime"` // CHANGES stimmt nicht in Dokumentation, dort string, tatsächlich int
	Active	bool `json:"active"` // CHANGES stimmt nicht in Dokumentation, dort string, tatsächlich bool
}

// Get Pole from Station Details as String

func (std *StationDetail)GetPoleAsString()string{
	return strconv.Itoa(std.Pole)
}

// Response Object  for Interface Haltestellenpaket

type Station struct {
    LongName string `json:"longName"`
    ShortName string `json:"shortName"`
    Longitude string `json:"longitude"`
    Latitude string `json:"latitude"`
    HafasId string `json:"hafasId"`
    ElementId int `json:"elementId"` // CHANGES stimmt nicht in Dokumentation, dort string, tatsächlich int
}

// ## Requests 
// 
// Request for Interface 'Änderungen am Haltestellen- und Linienpaket ermitteln'

func GetUpdates(params UpdateParams) (*[]Update,error){   
	var record []Update
    err:=makeRequest(getUrlForUpdate(params),&record)
	return &record, err
}

// Request for Interface Haltepunkte

func GetStationDetail(hafasId string) ([]StationDetail,error){
    par :=make(map[string][]string)
    par["stationId"]=[]string{hafasId}
     
    values:=getValues(par)
    urlV:=getUrlWithValues(SITE_URL+MODUL_URL+STATIONS_DETAIL,values)
   
	var record []StationDetail
	err:=makeRequest(urlV.String(),&record)
    
	return record, err
}

// Request for Interface Haltestellenpaket

func GetStationsPackage(Id string) (*StationPackage,error){	
	var record StationPackage
	err:=makeRequest(getUrl(SITE_URL+MODUL_URL+STATIONS_URL+Id).String(),&record)
	return &record, err
}
// ## Request From Response Object
// Get Station Detail from Station

func (st *Station)GetStationDetails()([]StationDetail,error){
	return GetStationDetail(st.HafasId)
}

// Get only active Station Detail Objects from Station 

func (st *Station)GetActiveStationDetails()([]StationDetail,error){
	temp,err:=GetStationDetail(st.HafasId)
	if err!=nil {
		result:=[]StationDetail{}
		for _,v:= range temp {
			if v.Active {
				result=append(result,v)
			}
		}
		return result,nil
	}
	return temp,err
}

// Get StationMonitor from Station

func (st *Station)GetStationMonitorTimeNow()(*Journey,error){
	time:="null"
	params:=JourneyParams{st.HafasId,&time,nil,GetSupportedModes()[0],GetSupportedNeedPlatformDetail()[0],nil}
	return GetStationMonitor(params)
}

// Get StationMonitor from Station with Additional Parameter (see stationMonitor.go)
// Parameter HafasId will be overwritte with HafasId of Station

func (st *Station)GetStationMonitorWithAdditionalParams(params JourneyParams)(*Journey,error){
	params.HafasId=st.HafasId
	return GetStationMonitor(params)
}

// Get StationMonitor from StationDetail Object
// TODO sollte auch ohne hafasID gehen

func (std *StationDetail)GetStationMonitorTimeNow(hafasId string)(*Journey,error){
	time:="null"
	poleAsStringArray:=[]string{std.GetPoleAsString()}
	params:=JourneyParams{hafasId,&time,&poleAsStringArray,GetSupportedModes()[0],GetSupportedNeedPlatformDetail()[0],nil}
	return GetStationMonitor(params)
}
// 
// 
// ## Other 
// 
func getUrlForUpdate(params UpdateParams) (string){
    return PROTCOL_AND_API_COMAIN+SITE_URL+UPDATE_URL+"?regionID="+params.RegionId+"&time="+*params.LastUpdateStationsTime+TIMESEPARATOR+*params.LastUpdateLinesTime+TIMESEPARATOR+*params.UnusedTime
}
