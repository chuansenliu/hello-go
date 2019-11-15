/*


{
    "resp": {
        "respCode": "000000",
        "respMsg": "成功",
        "app": {
            "appId": "xxxxxx"
        }
    }
}

*/

package main

import (
	"encoding/json"
	"fmt"
)

type AppInfo struct {
	AppId string `json "appId"`
}

type Response struct {
	RespCode string  `json "respCode"`
	RespMsg  string  `json "respMsg"`
	App      AppInfo `json "app"`
}

type JsonResult struct {
	response Response `json "resp"`
}

func main() {

	jsonstr := []byte(`{"resp": {"respCode": "000000","respMsg": "ssss","app": {"appId": "xxxxxx"}}}`)

	//fmt.Println(jsonstr)

	var resp JsonResult
	err := json.Unmarshal(jsonstr, &resp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v", resp)
	}
}
