package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)


type Alarm struct {
	Date string `json:"date"`
	RecID int `json:"recID"`
	AlarmArea string `json:"alarmArea"`
	AlarmType string `json:"alarmType"`
	AlarmColor string `json:"alarmColor"`
}

type RspData struct {
	subAlarm []Alarm
}

func main(){
	random_value := rand.Int31()  // Float32()
	time_value := time.Now().Unix() * 1000
	data_url := fmt.Sprintf("https://weather.121.com.cn/data_cache/szWeather/alarm/szAlarm.js?random=0.%d&_=%d" , random_value , time_value)
	fmt.Println(data_url)

	data_req, err_new_request := http.NewRequest("GET", data_url, nil)
	if err_new_request != nil {
		panic(err_new_request.Error())
	}
	data_rsp, err_get_response := http.DefaultClient.Do(data_req)
	if err_get_response != nil {
		panic(err_get_response.Error())
	}
	defer data_rsp.Body.Close()

	data_js_bytes, err_read_bytes := ioutil.ReadAll(data_rsp.Body)
	if err_read_bytes != nil {
		panic(err_read_bytes.Error())
	}
	data_js := string(data_js_bytes)
	//fmt.Println(data_js)
	data_json_str := data_js[150 :len(data_js)-12]
	//fmt.Println(data_json_str)

	var data_json_obj RspData
	err_unmarshal := json.Unmarshal([]byte(data_json_str), &data_json_obj)
	if err_unmarshal != nil {
		panic(err_unmarshal.Error())
	}
	fmt.Println(data_json_obj)  // currently failed
}