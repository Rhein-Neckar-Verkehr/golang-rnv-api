// # Interfaces from rnv Start.Info API
//  - Haltestelleninformationen
//  - Fahrtinformatonen
// 
// 
package api
// 
import (
	"errors"
	"time"
)
// ## Parameter
// 
// Parameter for Interface Haltestelleninformationen
// 
type StationInfoTransferParams struct{
    Lines *[]string
    HafasId *string
    DepartureTime *string
}

// all Parameter can be nil
// 
func (params *StationInfoTransferParams)Validate()error{
	if params.Lines!=nil{
		for _,l:= range *params.Lines{
			err:=validateLine(l)
			if err!=nil{return err}
		}
	}
	if params.HafasId!=nil{
		err:=validateStationId(*params.HafasId)
		if err!=nil{return err}
	}
	if params.DepartureTime!=nil{
		err:=validateTimeAsUnixTimestamp(*params.DepartureTime)
		if err!=nil{return err}
	}
	return nil
}
func (params StationInfoTransferParams)String()string{
	result:="{"
	if params.Lines!=nil{result=result+arrayToString(*params.Lines)+" "}else{ result=result+"<nil> "}
	if params.HafasId!=nil{result=result+*params.HafasId+" "}else{ result=result+"<nil> "}
	if params.DepartureTime!=nil{result=result+*params.DepartureTime}else{ result=result+"<nil>"}
	return result+"}"
}
func (params *StationInfoTransferParams)SetDepartureTime(t time.Time){
	params.DepartureTime=GetTimeUnix(t)
}
func (params *StationInfoTransferParams)AddLine(line string)error{
	err:=validateLine(line)	
	if err!=nil{return err}
	if params.Lines==nil{
		ne:=[]string{line}
		params.Lines=&ne
	} else {
		ne:=append(*params.Lines,line)
		params.Lines=&ne
	}
	return nil
}
// 
func (params *StationInfoTransferParams)GetAsUrlParams()(map[string][]string,error){
	err:=params.Validate()	
	par :=make(map[string][]string)
	
// same result from function, but skip unnecesarry operations
// 
	if (err!=nil){
		return par,err
	}	
	if(params.Lines!=nil){
        par["lines"]=[]string{GetLinesSeparated(*params.Lines)}
    }  
    if(params.HafasId!=nil){
        par["hafasID"]=[]string{*params.HafasId} 
    } 
    if(params.DepartureTime!=nil){
        par["departureTime"]=[]string{*params.DepartureTime} 
    }   
	return par,err
}
// 
const SUPPORTED_COMBINATIONS_OF_JOURNEY_INFO_TRANSFER_PARAM string=" HafasId | HafasId + Poles | HafasId + DepartureTime | HafasId + DepartureTime + Poles | DepartureTime"
const ERROR_SUPPORTED_COMBINATIONS_OF_JOURNEY_INFO_TRANSFER_PARAM="Not valid Combination of JourneyInfoTranser. Valid Combinations are: "+SUPPORTED_COMBINATIONS_OF_JOURNEY_INFO_TRANSFER_PARAM

// Parameter for Interface Fahrtinformationen
// 
type JourneyInfoTranserParams struct{
    HafasId *string
    Poles *[]string
    DepartureTime *string
}
// 
func (params *JourneyInfoTranserParams)AddPole(pole string)error{
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
func (params *JourneyInfoTranserParams)SetDepartureTime(t time.Time){
	params.DepartureTime=GetTimeUnix(t)
}
func (params JourneyInfoTranserParams)String()string{
	p:=params.Poles
	h:=params.HafasId
	d:=params.DepartureTime
	result:="{"	
	if h!=nil {result=result+*h+" "} else {result=result+"<nil> "}
	if p!=nil {result=result+arrayToString(*p)+" "} else {result=result+"<nil> "}
	if d!=nil {result=result+*d} else {result=result+"<nil>"}
	return result+"}"
}
// 
// 
func (params *JourneyInfoTranserParams)Validate()error{
	if !(params.HafasId!=nil || (params.HafasId!=nil && params.Poles!=nil) || (params.HafasId!=nil && params.DepartureTime!=nil) || params.DepartureTime!=nil )  {
		return errors.New(ERROR_SUPPORTED_COMBINATIONS_OF_JOURNEY_INFO_TRANSFER_PARAM)
	}
	if params.HafasId!=nil{
		err:=validateStationId(*params.HafasId)
		if err!=nil{return err}
	}
	if params.Poles!=nil{
		for _,val:= range *params.Poles{
			err:=validatePoles(val)
			if err!=nil{return err}
		}
	}
	if params.DepartureTime!=nil{
		err:=validateTimeAsUnixTimestamp(*params.DepartureTime)
		if err!=nil{return err}
	}
	return nil
}
// 
func (params *JourneyInfoTranserParams)GetAsUrlParams()(map[string][]string,error){
	err:=params.Validate()	
	par :=make(map[string][]string)
	if (err!=nil){
		return par,err
	}
	if(params.Poles!=nil){
        par["poles"]=[]string{GetLinesSeparated(*params.Poles)}
    } 
    if(params.HafasId!=nil){
        par["hafasID"]=[]string{*params.HafasId} 
    } 
    if(params.DepartureTime!=nil){
        par["departureTime"]=[]string{*params.DepartureTime} 
    }     
	return par,err
}
// ## Response Objects
// 
// Response Object Interface Haltestelleninformationen
// 
type StationInfoTransfer struct{
    Title string `json:"title"`
    Text string `json:"text"`
    Id int `json:"Id"` 
    LineId string `json:"lineId"`
    StationIds []string `json:"stationIds"`
    StationNames []string `json:"stationNames"`    
    Url string `json:"url"`    
    Author string `json:"author"`
    Created uint64 `json:"created"`
    ValidFrom uint64 `json:"validFrom"`
    ValidTo uint64 `json:"validTo"`
    DisplayFrom uint64 `json:"displayFrom"`
    DisplayTo uint64 `json:"displayTo"` 
}

// TODO check if uint64 can be replaced by uint. Diffrent definitions for date from json
// 
func (c *StationInfoTransfer)IsCurrentlyValid()bool{
	return IsInCurrentlyInTimeRange(c.ValidFrom,c.ValidTo)
}
// 
func (c *StationInfoTransfer)IsValidInFuture()bool{
	return IsTimeRangeInTheFuture(c.ValidFrom)
}
// 
func (c *StationInfoTransfer)IsCurrentlyDisplayed()bool{
	return IsInCurrentlyInTimeRange(c.DisplayFrom,c.DisplayTo)
}
// 
func (n *StationInfoTransfer)AffectsGivenStation(stationId string)bool{
	return ListContainsValue(n.StationIds,stationId)
}

// Response Object Interface Fahrtinformationen

type JourneyInfoTranser struct{
    Title string `json:"title"`
    Text string `json:"text"`    
    Id int `json:"Id"`
    StationId string `json:"stationId"`
    StationName string `json:"stationName"`
    Poles []string `json:"poles"`    
    Author string `json:"author"`
    Created uint64 `json:"created"`
    ValidFrom uint64 `json:"validFrom"`
    ValidTo uint64 `json:"validTo"`  
}
// 
func (c *JourneyInfoTranser)IsCurrentlyValid()bool{
	return IsInCurrentlyInTimeRange(c.ValidFrom,c.ValidTo)
}
// 
func (c *JourneyInfoTranser)IsValidInFuture()bool{
	return IsTimeRangeInTheFuture(c.ValidFrom)
}
// 
// 
// 
// ## Requests
// 
// Request for Interface Fahrtinformation
// 
func GetLineInfo(params JourneyInfoTranserParams)(*[]JourneyInfoTranser,error){
    par,err:=params.GetAsUrlParams()
    if err!=nil{
		return nil,err
	}
    values:=getValues(par)
    urlV:=getUrlWithValues(SITE_URL+MODUL_URL+INFO_LINE_URL,values)
      
	var record []JourneyInfoTranser
	err=makeRequest(urlV.String(),&record)
    
	return &record, err
}

// Request for Haltestelleninformationen
// 
func GetStationInfo(params StationInfoTransferParams)(*[]StationInfoTransfer,error){
    par,err :=params.GetAsUrlParams()
    if err!=nil{
		return nil,err
	}    
    values:=getValues(par)
    urlV:=getUrlWithValues(SITE_URL+MODUL_URL+INFO_STATION_URL,values)
// 
	var record []StationInfoTransfer
	err=makeRequest(urlV.String(),&record)
    
	return &record, err
}

// ## Other
