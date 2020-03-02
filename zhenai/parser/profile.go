package parser

import (
	"crawler_go/engine"
	"crawler_go/model"
	"regexp"
	"strconv"
)

var (
	ageRe        = regexp.MustCompile(`"age":([0-9]+),`)
	heightRe     = regexp.MustCompile(`"heightString":"([0-9]*)cm",`)
	incomeRe     = regexp.MustCompile(`"月收入:([^"]*)"`)
	genderRe     = regexp.MustCompile(`"genderString":"([^"]*)",`)
	marriageRe   = regexp.MustCompile(`"marriageString":"([^"]*)"`)
	educationRe  = regexp.MustCompile(`"educationString":"([^"]*)"`)
	occupationRe = regexp.MustCompile(`"workProvinceCityString":"([^"]*)"`)
	idURLRe      = regexp.MustCompile(`https://album.zhenai.com/u/([\d]+)`)
)

// ParserProfile ...
func ParserProfile(
	contents []byte, url string,
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
		Items: []engine.Item{
			{
				URL:     url,
				Type:    "zhenai",
				ID:      extractString([]byte(url), idURLRe),
				Payload: profile,
			},
		},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
