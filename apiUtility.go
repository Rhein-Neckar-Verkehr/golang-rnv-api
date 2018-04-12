package api

import (

    "os"
	"io/ioutil"
	"net/http"
    "encoding/json"
    con "strconv"
    "sync"
    "time"
    "log"
	"errors"
	"strings"
	"net"
	"net/url"

)

// Time Formats for Interface Änderungen an Haltestellen und Linienpakete ermitteln
//
const TIMEFORMATDATE string="2006-01-02"
const TIMEFORMATTIME string="15:04"
const TIMEFORMAT string="2006-01-02+15:04"

// TODO refactor this!
//
const TRUE string="true"
const FALSE string="false"
var booleanA =[2]string{"true","false"}

// used for async calls. See Station Monitor
//
type HttpResponse struct {
  Url      string
  Response *http.Response
  Err      error
}

//

// The API Token is stored in tokenConf.json. The struct Token is a singleton.
//
type token struct{
    UserToken string `json:"token"`
}
var apiToken *token
var once sync.Once

func parseKindOfTourToTourType(kindOfTour string)string{
	temp:=strings.Replace(kindOfTour,"REF","",-1)
	temp=strings.Replace(temp,"AUS","",-1)
	return temp
}

// try to parse the content to error description, if it works, take that message as error
// otherwise return the given error, which should be the json Parsing error.

func handleErrorsFromServer(content []byte,p_err error)error{
	errd:=string(content)
	if strings.Contains(errd,"errordescription") {
		// Create Error Message from given ErrorDescription
		p_err = errors.New(string(errd))
	}
	return p_err;
}

func parseResponseObject(content []byte,v interface{})error {
	err := json.Unmarshal(content, v)
	if err != nil {
		err=handleErrorsFromServer(content,err)
	}
	return err
}

func makeRequest(url string,responseObject interface{})error{
	log.Println(url)
	content,err:=getContent(url)
	if err != nil {
		return err
	}
	return parseResponseObject(content,responseObject)
}

func IsInCurrentlyInTimeRange(from uint64, to uint64)bool{
	fromT:=time.Unix(int64(from), 0)
	toT:=time.Unix(int64(to), 0)
	return time.Now().After(fromT) && time.Now().Before(toT)
}

func IsTimeRangeInTheFuture(from uint64)bool{
	fromT:=time.Unix(int64(from), 0)
	return time.Now().Before(fromT)
}
func GetLinesSeparated(lines []string)(string){
    var values string=""
    for _,l := range lines{
        values=values+l+";"
    }
    valuesTrimmed  := values[:len(values)-1]
    return valuesTrimmed
}

func SplitLinesSeperated(lines string)[]string{
	return strings.Split(lines,";")
}

func SplitLinesSeperatedComma(lines string)[]string{
	return strings.Split(lines,",")
}

func ListContainsValue(list []string, value string)bool{
	for _,v := range list {
		if v == value {
			return true
		}
	}
	return false
}

// Utility Functions for Time Parameters

func GetTimeFormat(t time.Time)(*string){
    result:=t.Format(TIMEFORMAT)
    return &result
}
func GetTimeFormatNow()(*string){
    return GetTimeFormat(time.Now())
}
func GetTimeUnixNow()(*string){
    return GetTimeUnix(time.Now())
}
func GetTimeUnix(t time.Time)(*string){
    unixTime:=t.Unix()
    result:=con.FormatInt(unixTime, 10)
    return &result
}
func arrayToString(arr[]string)string{
	/*result:="["
	for _,v:=range arr{
		result=result+v+" "
	}
	result=result[:len(result)-1]
	result=result+"]"
	return result	*/
	return "["+strings.Join(arr," ")+"]"
}
// Functions for api token

func GetToken() (*string,error) {
		// zuerst nach environment var suchen
		envToken:=os.Getenv("rnv_api_token")
		if (envToken!=""){
			return &envToken,nil
		}
		var err error
    once.Do(func() {
        apiToken,err = readToken()
    })
    return &apiToken.UserToken,err
}
func readToken() (*token,error){
	t := token{}
    file, errFile := os.Open("tokenConf.json")
	if  errFile!=nil{
		log.Println(errFile)
        errFile=errors.New("Cannot read Token. Please Check your File tokenConf.json or set environment variable rnv_api_token")
		return &t,errFile
    }
    decoder := json.NewDecoder(file)
    err := decoder.Decode(&t)
	if err!=nil {
		log.Println(err)
        err=errors.New("Cannot read Token. Please Check your content of tokenConf.json or set environment variable rnv_api_token")
    }

    return &t,err
}

func GetRnvApiHttpRequest(url string)(*http.Request,error){
	reqe, err := http.NewRequest("GET", url, nil)
	if err!=nil {
		return nil,err
	}
	tok,errT:= GetToken()
	if errT!=nil {
		return nil,errT
	}
	reqe.Header.Add("Accept", "application/json")
	reqe.Header.Add("RNV_API_TOKEN",*tok)

	return reqe,nil
}

// This Function is called for all API calls.
// It always requests json objekts. The Token is added to the Request Header.

func getContent(targetUrl string) ([]byte,error) {
	req, err := GetRnvApiHttpRequest(targetUrl)
	log.Println(targetUrl)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	u,errU:=url.Parse(os.Getenv("https_proxy"))
	if errU != nil {
		log.Println(errU)
		return nil, errU
	}
  tr:=&http.Transport{}
  if (u.String()!=""){
    tr = &http.Transport{
  		Dial: (&net.Dialer{
      Timeout: 5 * time.Second,
    	}).Dial,
  		TLSHandshakeTimeout: 5 * time.Second,
  	  Proxy: http.ProxyURL(u),
  	}
  } else {
    tr = &http.Transport{
  		Dial: (&net.Dialer{
      Timeout: 5 * time.Second,
    	}).Dial,
  		TLSHandshakeTimeout: 5 * time.Second,
  	}
  }

	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: tr,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

  //  log.Println("RES: "+resp.Status+" "+resp.Proto)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return body, nil
}
