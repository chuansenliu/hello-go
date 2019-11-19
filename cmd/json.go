package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/eehut/hello-go/command"
)

type ConfigStruct struct {
	Host              string   `json:"host"`
	Port              int      `json:"port"`
	AnalyticsFile     string   `json:"analytics_file"`
	StaticFileVersion int      `json:"static_file_version"`
	StaticDir         string   `json:"static_dir"`
	TemplatesDir      string   `json:"templates_dir"`
	SerTcpSocketHost  string   `json:"serTcpSocketHost"`
	SerTcpSocketPort  int      `json:"serTcpSocketPort"`
	Fruits            []string `json:"fruits"`
}

type Other struct {
	SerTcpSocketHost string   `json:"serTcpSocketHost"`
	SerTcpSocketPort int      `json:"serTcpSocketPort"`
	Fruits           []string `json:"fruits"`
}

// doString is command function
func doJson(argv []string) int {

	jsonStr := `{
		"host": "http://localhost:9090",
		"port": 9090,
		"analytics_file": "",
		"static_file_version": 1,
		"static_dir": "E:/Project/goTest/src/",
		"templates_dir": "E:/Project/goTest/src/templates/",
		"serTcpSocketHost": ":12340",
		"serTcpSocketPort": 12340,
		"fruits": ["apple", "peach"]}`

	var dat map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &dat)

	if err == nil {
		fmt.Println("==================== str -> map =================")
		fmt.Println(dat)
		fmt.Println(dat["port"])
	}

	var st ConfigStruct
	err = json.Unmarshal([]byte(jsonStr), &st)

	if err == nil {
		fmt.Println("==================== str -> struct =================")
		fmt.Println(st)
		fmt.Println(st.Host)
	}

	var dat1 map[string]interface{}
	err = json.NewDecoder(strings.NewReader(jsonStr)).Decode(&dat1)

	if err == nil {
		fmt.Println("==================== ioReader -> map =================")
		fmt.Println(dat1)
		fmt.Println(dat1["host"])
	}

	var st1 ConfigStruct
	err = json.NewDecoder(strings.NewReader(jsonStr)).Decode(&st1)

	if err == nil {
		fmt.Println("==================== ioReader -> struct =================")
		fmt.Println(st1)
		fmt.Println(st1.Host)
	}

	/* ====  */
	b2, e2 := json.Marshal(st1)
	if e2 == nil {
		fmt.Println("==================== struct -> str =================")
		fmt.Println(string(b2))
	}

	/* ====  */
	b3, e3 := json.MarshalIndent(st1, "PREFIX:", "  ")
	if e3 == nil {
		fmt.Println("==================== struct -> str(indent) =================")
		fmt.Println(string(b3))
	}

	fmt.Println("==================== map -> stdout =================")
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(dat1)

	return 0
}

func init() {
	command.AddCommand("json", "json test", "Usage:\n json [d|e]", doJson)
}
