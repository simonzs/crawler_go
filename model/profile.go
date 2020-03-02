package model

import "encoding/json"

// Profile 人的基本信息
type Profile struct {
	Name   string // 姓名
	Gender string // 性别
	Age    int    // 年龄
	Height int    // 身高
	// Weight    int    // 体重
	Income    string // 收入
	Marriage  string // 婚姻状况
	Education string // 教育状况

	Occupation string // 工作地
	// Hokou      string // 户口，籍贯
	// Xinzuo     string // 星座
	// House string // 房子情况
	// Car   string // 车子情况
}

// FromJsonObj ...
func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	bytes, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(bytes, &profile)
	return profile, err
}
