package model

import "encoding/json"

type Profile struct {
	Name	   string   //昵称
	Gender     string   //性别
	Age 	   int      //年龄
	Height     int      //身高
	Income     string   //收入
	Marriage   string   //婚况
	//Education  string   //教育
	//Occupation string   //职业
	XingZuo    string   //星座
	WorkPlace  string   //居住地
}


func FromJsonObj(o interface{}) (Profile, error){
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, err
}