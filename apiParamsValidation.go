//
//  Validation Functions used for several Params Objects
package api

import (
	"regexp"
	"strconv"
	"errors"
	"strings"
	"log"
)

// TODO compile all in init()
//
const ERROR_START string="Invalid value for"
const CATCH_ERROR_MSG_START string="Used default value for"

const ERROR_TIME_AS_STRING string =ERROR_START+ " time or departureTime"
const ERROR_TIME_UNIX_AS_STRING string =ERROR_START+ " time or departureTime"
const ERROR_STATION_ID string=ERROR_START+" stationId or hafasId"
const ERROR_POLES string=ERROR_START+" poles"
const ERROR_LINE string=ERROR_START+" lines or lineId"
const ERROR_INDEX string=ERROR_START+" LineJourneyParams.StopIndex"
const ERROR_TOUR_TYPE string=ERROR_START+" LineJourneyParams.TourType (452|454)"

// Validate departureTime and time
//
func validateTimeAsUnixTimestamp(departureTime string)error{
	_, err := strconv.ParseUint(departureTime, 10, 64)
    if err != nil {
        return errors.New(ERROR_TIME_UNIX_AS_STRING)
    }
	return nil
}
func validateTimeAsString(time string)error{
	valid:=regexp.MustCompile(`^\d\d\d\d-\d\d-\d\d \d\d:\d\d$|^\d\d\d\d-\d\d-\d\d\+\d\d:\d\d$`)
	res:=valid.MatchString(time)
	if !res {
		return errors.New(ERROR_TIME_AS_STRING)
	}
	return nil
}
func validateIndex(index string)error{
	valid:=regexp.MustCompile(`^\d{1,2}$`)
	res:=valid.MatchString(index)
	if !res {
		return errors.New(ERROR_TIME_AS_STRING)
	}
	return nil
}
func validateTourType(tourType string)error{
	valid:=regexp.MustCompile(`^45[24](AUS|REFAUS|)$`)
	res:=valid.MatchString(tourType)
	if !res {
		return errors.New(ERROR_TOUR_TYPE)
	}
	return nil
}


// Validate stationId and hafasId
//
func validateStationId(stationId string)error{
	if len(stationId)>4{
		return errors.New(ERROR_STATION_ID)
	}
	valid:=regexp.MustCompile(`^\d+$`) // NOTE muss mindestens aus eienr Zahl bestehen? Abfahrtsmonitor ^\d{3,4}$
	res:=valid.MatchString(stationId)
	if !res {
		return errors.New(ERROR_STATION_ID)
	}
	return nil
}

// Validate Poles
//
func validatePoles(poles string)error{
	splitted:=strings.Split(poles,";")
	valid:=regexp.MustCompile(`^\d\d$|^\d$`) // führende Null ein Fehler??
	for _,v:= range splitted {
		res:=valid.MatchString(v)
		if !res {
			return errors.New(ERROR_POLES)
		}
	}
	return nil
}

func validatePole(pole string)error{
	valid:=regexp.MustCompile(`^\d\d$|^\d$`) // führende Null ein Fehler??
	res:=valid.MatchString(pole)
	if !res {
		return errors.New(ERROR_POLES)
	}
	return nil
}

// Validate LineId or LineLabel
//
func validateLine(line string)error{
	valid:=regexp.MustCompile(`^(\d{1,3}|[456]A)$|^((\d{1,3}|[456]A),)+(\d{1,3}|[456]A)$|^(M[12345]|S[EFJKL]|X|632A)$`)
	res:=valid.MatchString(line)
		if !res {
			log.Println(line)
			return errors.New(ERROR_LINE)
	}
	return nil
}
func validateLines(lines string)error{
	splitted:=strings.Split(lines,";")
	for _,v:= range splitted {
		err:=validateLine(v)
		if err!=nil {
			return err
		}
	}
	return nil
}
