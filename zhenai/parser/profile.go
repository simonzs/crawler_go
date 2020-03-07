package parser

import (
	"regexp"
	"strconv"

	"github.com/simonzs/crawler_go/engine"
	"github.com/simonzs/crawler_go/model"
)

var (
	ageRe        = regexp.MustCompile(`"age":([0-9]+),`)
	heightRe     = regexp.MustCompile(`"heightString":"([0-9]*)cm",`)
	incomeRe     = regexp.MustCompile(`"月收入:([^"]*)"`)
	genderRe     = regexp.MustCompile(`"genderString":"([^"]*)",`)
	marriageRe   = regexp.MustCompile(`"marriageString":"([^"]*)"`)
	educationRe  = regexp.MustCompile(`"educationString":"([^"]*)"`)
	occupationRe = regexp.MustCompile(`"workProvinceCityString":"([^"]*)"`)
	idURLRe      = regexp.MustCompile(`.*album.zhenai.com/u/([\d]+)`)
)

// ParserProfile ...
func parserProfile(
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

// ProfileParser ...
type ProfileParser struct {
	userName string
}

// Parse ...
func (p *ProfileParser) Parse(
	contents []byte, url string) engine.ParserResult {
	return parserProfile(contents, url, p.userName)
}

// Serialize ...
func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "parserProfile", p.userName
}

// NewProfileParser ...
func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}
