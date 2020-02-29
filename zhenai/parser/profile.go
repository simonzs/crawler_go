package parser

import (
	"crawler_go/engine"
	"crawler_go/model"
	"regexp"
	"strconv"
)

const ageRe = `"age":([0-9]+),`
const heightRe = `"heightString":"([0-9]*)cm",`

const incomeRe = `"月收入:([^"]*)"`
const genderRe = `"genderString":"([^"]*)",`
const marriageRe = `"marriageString":"([^"]*)"`
const educationRe = `"educationString":"([^"]*)"`
const occupationRe = `"workProvinceCityString":"([^"]*)"`

// ParserProfile ...
func ParserProfile(
	contents []byte,
	name string) engine.ParserResult {

	profile := model.Profile{}

	profile.Name = name

	age, err := strconv.Atoi(
		extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(
		extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	profile.Income = extractString(contents, incomeRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)

	result := engine.ParserResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, reStr string) string {
	re := regexp.MustCompile(reStr)
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
