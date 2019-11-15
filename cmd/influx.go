package main

import (
	"fmt"

	"github.com/eehut/hello-go/command"
	"github.com/influxdata/influxdb/client/v2"
)

func queryDB(clnt client.Client, cmd string, db string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: db,
	}
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

// CommandInflux is for influxdb test
type CommandInflux struct {
	Desc command.CommandDesc
}

// GetDesc return the description
func (cmd CommandInflux) GetDesc() *command.CommandDesc {
	return &cmd.Desc
}

// Execute is command function
func (cmd CommandInflux) Execute(argv []string) int {

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://10.72.7.24:8086",
	})
	if err != nil {
		fmt.Println("connect to db failed")
		fmt.Println(err)
		return 0
	}
	defer c.Close()

	if len(argv) > 0 && argv[0] == "c" {
		_, err = queryDB(c, "create database hello", "hello")

		if err != nil {
			fmt.Println("Create database failed")
		}
	}

	if len(argv) > 0 && argv[0] == "d" {
		resp, ret := queryDB(c, "show databases", "")

		if ret != nil {
			fmt.Println("show databases failed")
		} else {
			fmt.Println(resp)
		}
	}

	if len(argv) > 0 && argv[0] == "cc" {
		resp, ret := queryDB(c, "select count(cpu) from apple", "system")

		if ret != nil {
			fmt.Println("show databases failed")
		} else {
			fmt.Println(resp)
		}
	}

	if len(argv) > 0 && argv[0] == "l" {
		resp, ret := queryDB(c, "select * from apple limit 10", "system")

		if ret != nil {
			fmt.Println("op databases failed")
		} else {

			fmt.Printf("result num = %d", len(resp))

			fmt.Printf("columns:\n")
			fmt.Println(resp[0].Series[0].Columns)

			fmt.Printf("name:\n")
			fmt.Println(resp[0].Series[0].Name)

			for i, row := range resp[0].Series[0].Values {
				fmt.Printf(" --- %d\n", i)
				fmt.Println(row)
			}
		}

	}

	return 0
}

func init() {
	command.AddCommand("influx", CommandInflux{Desc: command.CommandDesc{Info: "influx package test", Help: "Usage:\n influx"}})
}
