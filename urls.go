// # Urls
// All URLs for the Interfaces are provided as Constants
// 
package api
// 
import (
    "net/url"
)
// 
// Definition of Protocol, Domain etc.
// 
const PROTOCOL="http"
const API_DOMAIN = "rnv.the-agent-factory.de:8080"
const PROTCOL_AND_API_COMAIN="http://rnv.the-agent-factory.de:8080"
const SITE_URL = "/easygo2/api/"
const MODUL_URL="regions/rnv/modules/"

// ## Interface URL
// Each of the following constant variables represent an Interface
//  - Linienpaket
// 
const LINE_PACKAGE_URL=LINE_URL+"/allJourney"
// 
//  - Haltestellenpaket

const STATIONS_URL = "stations/packages/"
// 
//  - Änderungen am Haltestellen- und Linienpaket ermitteln

const UPDATE_URL="update"

//  - Haltepunkte
// 
const STATIONS_DETAIL="stations/detail"

//  - Haltestellenmonitor

const DEPARTURES_URL = "stationmonitor/element"
// 
//  - Fahrtverläufe

const LINE_URL = "lines"
const LINE_ALL_URL = LINE_URL + "/all" // Linieninformation siehe Linienpaket

//  - News

const NEWS_URL = "news"

//  - Ticker

const TICKER_URL = "ticker"
// 
//  - Entfallene Linien

const CANCELED_LINE_URL="canceled/line"
// 
//  - Haltestelleninformation

const INFO_STATION_URL="info/station"
// 
//  - Fahrtinformationen
const INFO_LINE_URL="info/journey"

//  - Liniennetzpläne

const MAPS_URL="maps"
// 
// ### Url helper Functions
// These functions interact with golang net/url package with the interface url strings.
// 
// #### getValues
// Returns the given *params* in net/url Values struct.
// *params* map[string][]string: Queryparameter for GET Request
// 
func getValues(params map[string][]string)(*url.Values){
    val:=url.Values{}
    for k,v :=range params{
        val.Set(k,v[0]) 
    }
    return &val
}
// 
// #### getUrl
// Returns the net/url.URL struct from given *path*. *Host* and *Scheme* will be set.
// *path* string: Interface route.

func getUrl(path string)(*url.URL){
    return &url.URL{
			Host:   API_DOMAIN,
			Scheme: PROTOCOL,
            Path:path,
		}
}
// 
// #### getUrlWithValues
// Like getUrl but additionally it sets the *RawQuery* of *net/http.URL* struct with the given *params*.
// *params* net/url.Values: Queryparameter for GET Request
// *path* string: Interface Route. 

func getUrlWithValues(path string,params *url.Values)(*url.URL){
    return &url.URL{
			Host:   API_DOMAIN,
			Scheme: PROTOCOL,
            Path:path,
        RawQuery:params.Encode(),
		}
}
