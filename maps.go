// # Interfaces from rnv Start.Info API
//  - Liniennetzpläne
// 
package api

import (
	"errors"
	"log"
)

// ## Parameter
// 
// default is 128

var supportedThumbnailSizes  = [5]string{"32","64","128","256","512"} 
func GetSupportedThumbnailSizes()[5]string{return supportedThumbnailSizes}
const ERROR_THUMBNAIL_SIZE string=ERROR_START+" MapEntityParams.ThumbnailSize"
const CATCH_ERROR_MSG_THUMBNAIL_SIZE=CATCH_ERROR_MSG_START+" MapEntityParams.ThumbnailSize: 128"

// default is png
// 
var supportedFormats =[2]string{"pdf","png"} 
func GetSupportedFormats()[2]string{return supportedFormats}
const ERROR_FORMAT string=ERROR_START+" MapEntityParams.Format"
const CATCH_ERROR_MSG_FORMAT=CATCH_ERROR_MSG_START+" MapEntityParams.Format: png"
// 
type MapEntityParams struct {
    ThumbnailSize *string
    Format *string
}
// 
func (params MapEntityParams)String()string{
	result:="{"
	if params.ThumbnailSize==nil {result=result+"<nil> "}else {result=result+*params.ThumbnailSize+" "}
	if params.Format==nil {result=result+"<nil>"}else {result=result+*params.Format}
	return result+"}"
}

// resets Parameter no nil if not valid
// 
func (params *MapEntityParams)Validate()error{
	if(params.ThumbnailSize!=nil){
		found:=false
		for _,v:= range supportedThumbnailSizes{
			if *params.ThumbnailSize==v{
				found=true
			}
		}
		if !found{
			params.ThumbnailSize=nil
			return errors.New(ERROR_THUMBNAIL_SIZE)
		}
	}
	if(params.Format!=nil){
		found:=false
		for _,v:= range supportedFormats{
			if *params.Format==v{
				found=true
			}
		}
		if !found{
			params.Format=nil
			return errors.New(ERROR_FORMAT)
		}
	}
	return nil
}
// 
func (params *MapEntityParams)GetAsUrlParams()(map[string][]string,error){
	err:=params.Validate()
	par :=make(map[string][]string)
	
// Catch errors and use default values (in this case nil)
// 
	if err==errors.New(ERROR_THUMBNAIL_SIZE){
		log.Println(CATCH_ERROR_MSG_MODE)
		params.ThumbnailSize=nil
		err=nil
	}
	if err==errors.New(ERROR_FORMAT){
		log.Println(CATCH_ERROR_MSG_FORMAT)
		params.Format=nil
		err=nil
	}
	if err!=nil{
		return par,err
	}
	if (params.ThumbnailSize!=nil){
		par["thumbnailSize"]=[]string{*params.ThumbnailSize} 
	}
	if (params.Format!=nil){
		par["format"]=[]string{*params.Format} 
	}
	return par,err
}

// ## Response Objects
// 
// Response Object MapEntity from Interface Lieniennetzpläne
// 
type MapEntity struct{
    
    Name string `json:"name"`
    ShortName string `json:"shortName"`
    Author string `json:"author"`
    ThumbnailUrl string `json:"thumbnailUrl"`
    MapUrl string `json:"mapUrl"`
    ValidFrom string `json:"validFrom"`
}

// ## Requests
// 
// Request to Interface Lieniennetzpläne
// 
func GetMaps(params MapEntityParams)(*[]MapEntity,error){
    par,err:=params.GetAsUrlParams()    
    if err!=nil{
		return nil,err
	}
    values:=getValues(par)
    urlV:=getUrlWithValues(SITE_URL+MODUL_URL+MAPS_URL,values)
	var record []MapEntity   
	err=makeRequest(urlV.String(),&record)
	return &record, err
}

// ## Other
