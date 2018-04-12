// # Interfaces from rnv Start.Info API
//  - News
//  - Ticker
//
package api
import ("strings")
// ## Parameter
//
type TickerParams []string
//
func NewTickerParams()TickerParams{
	return TickerParams([]string{})
}

// Functions with pointer possible??

func (params TickerParams)AddLine(line string)(TickerParams,error){
	line=checkLineId(line)
	err:=validateLine(line)
	if err==nil{
		params=append(params,line)
	}
	return params,err
}

// lines contains several lines as string, seperated by "," TODO ? auch ;

func (params TickerParams)AddLines(lines string)(TickerParams,error){
	arr:=strings.Split(lines,",")
	var err error
	for _,v:= range arr {
		params,err=params.AddLine(v)
		if err!=nil{return params,err}
	}
	return params,err
}
//
func (params TickerParams)Validate()error{
	for _,v:= range params {
		err:=validateLine(v)
		if err!=nil{return err}
	}
	return nil
}
//
func (params TickerParams)GetQueryParameterAsString()(string,error){
	err:=params.Validate()
	result:=""
	if err==nil {
		var values string="?lines="
		for _,l := range params{
			values=values+l+";"
		}
		result= values[:len(values)-1]
	}
	return result,err
}

// ## Response Objects
//
// Response Object for Interface News and Ticker

type NewsEntry struct{
    RegionId int `json:"regionID"` //NOT IN RESULT!! APIBUG
    Title string `json:"title"`
	Text	string `json:"text"`
    ValidFrom uint64 `json:"validFrom"`
    dateAsString string `json:"dateAsString"`
    ValidTo uint64 `json:"validTo"`
    Lines string `json:"lines"`
    SeperatedLines []string `json:"seperatedLines"`
    ImgUrl string `json:"imgUrl"` // kein bild enthalten
    TextAsHtml string `json:"textAsHtml"`
    IsOldNews bool `json:"oldNews"`   //heißt nur oldNews
    ThumbUrl string `json:"thumbUrl"`   // no thumbUrl
    FurtherLink string `json:"furtherLink"` // no further link!
    ElementId int `json:"elmentID"`
}
//
func (n *NewsEntry)AffectsGivenLine(line string)bool{
	return ListContainsValue(n.SeperatedLines,line)
}

// ## Requests
//
// Request for Interface Ticker

func GetTickerForLines(params TickerParams)([]NewsEntry,error){
	query,err:=params.GetQueryParameterAsString()
	var record []NewsEntry
	if err !=nil {return record,err}
    urlV:=getUrl(SITE_URL+MODUL_URL+TICKER_URL)
//

	err=makeRequest(urlV.String()+query,&record)
	return record, err
}

// Request for Interface News

func GetNews()([]NewsEntry,error){
    urlV:=getUrl(SITE_URL+MODUL_URL+NEWS_URL)
	var record []NewsEntry
	err:=makeRequest(urlV.String(),&record)
	return record, err
}

// ## Other
