package validations

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"regexp"
)

func ValidatePhone(phone string)bool{

	countryRegex, err :=GetCountryRegexByPhone(phone)
	if err !=nil {
		return false
	}else {
		re := regexp.MustCompile(*countryRegex)
		return re.MatchString(phone)
	}
}

func GetCountryRegexByPhone(phone string) (*string,error)  {
	type country struct {
		Name        string `json:"name"`
		CountryCode string `json:"country_code"`
		Regex       string `json:"regex"`
	}
	plan, readErr:= ioutil.ReadFile("./database/countries.json")
	if readErr != nil {log.Fatal(readErr)}
	var countries []country
	err := json.Unmarshal(plan, &countries)
	if err !=nil {
		log.Fatalln(err.Error())
	}
	for _,country :=range countries  {
		matched, _ :=regexp.MatchString(country.CountryCode+".*",phone)
		if matched {
			return &country.Regex,nil
		}
	}
	return nil,err


}