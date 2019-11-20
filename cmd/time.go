package main

import (
	"fmt"
	"time"

	"github.com/eehut/hello-go/command"
)

// doString is command function
func doTime(argv []string) int {

	var a time.Time
	a = time.Now()

	fmt.Println(a)
	fmt.Println(a.Local())
	fmt.Println(a.UTC())
	fmt.Println(a.Zone())
	fmt.Println(a.Location())

	fmt.Println(a.Unix())

	b := time.Unix(2222222222, 444444)
	fmt.Println(b)

	if a.Before(b) {
		fmt.Println("time a is before b")
	}

	if b.Before(a) {
		fmt.Println("time b is before a")
	}

	lutc := time.UTC
	llocal := time.Local

	fmt.Println(lutc.String())
	fmt.Println(llocal.String())

	t := time.Date(2020, time.November, 10, 23, 11, 0, 0, time.Local)
	fmt.Println(t)

	ay, am, ad := a.Date()
	fmt.Println(ay, am, ad)

	fmt.Println("========== parse & format =================")

	const formt = "2006-01-02 15:04:05"
	const fmtZone = "2006-01-02 15:04:05 -0700 MST"

	if tt, ee := time.Parse(formt, "2019-11-18 20:30:30"); ee == nil {
		fmt.Println(tt)
	} else {
		fmt.Println(ee)
	}

	ll, _ := time.LoadLocation("Asia/Shanghai")

	if tt, ee := time.ParseInLocation(formt, "2019-11-20 20:30:30", ll); ee == nil {
		fmt.Println(tt)
	} else {
		fmt.Println(ee)
	}

	fmt.Println(a.Format(fmtZone))

	/*
		总结：
			使用自定义格式，必须使用固定的时间
			Mon Jan 2 15:04:05 -0700 MST 2006

			预定义的：

			    ANSIC       = "Mon Jan _2 15:04:05 2006"
			    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
			    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
			    RFC822      = "02 Jan 06 15:04 MST"
			    RFC822Z     = "02 Jan 06 15:04 -0700" // 使用数字表示时区的RFC822
			    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
			    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
			    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // 使用数字表示时区的RFC1123
			    RFC3339     = "2006-01-02T15:04:05Z07:00"
			    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"

	*/

	fmt.Println("========== duration =================")

	/* similar methods: Clock, Year, Month, Hour, Minute Second etc */

	/* time 计算， Add,Sub */
	/* 什么是 duration, 是一个整型数，以纳秒为单位，最大可表示 时间段为290年 */
	/*  常用的duration
	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)
	*/

	/*
		ParseDuration解析一个时间段字符串。一个时间段字符串是一个序列，
		每个片段包含可选的正负号、十进制数、可选的小数部分和单位后缀，如"300ms"、"-1.5h"、"2h45m"。
		合法的单位有"ns"、"us" /"µs"、"ms"、"s"、"m"、"h"。
	*/

	if d, e1 := time.ParseDuration("20h50m"); e1 == nil {
		fmt.Println(d)
		fmt.Println(d.Minutes())
		fmt.Println(d.Seconds())
	}

	tt, _ := time.Parse(formt, "2019-11-18 20:30:30")
	fmt.Println("since:", time.Since(tt))

	dd, _ := time.ParseDuration("30h")

	fmt.Println(tt.Add(dd))

	tr := time.NewTimer(5 * time.Second)
	te := time.NewTimer(60 * time.Second)
	time.AfterFunc(7*time.Second, func() {
		fmt.Println("<AfterFunc, only once>")
	})

	tk := time.NewTicker(8 * time.Second)

	for {

		select {

		//case <-time.After(1 * time.Second):
		//	fmt.Printf(".")
		case <-tk.C:
			fmt.Println("ticker up")

		case aa := <-tr.C:
			fmt.Println(aa)
			tr.Reset(5 * time.Second)

		case <-time.After(10 * time.Second):
			fmt.Println("<never reach, why?>")

		case <-te.C:
			fmt.Println("time's up")
			goto safeExit
		}
	}
safeExit:

	/*
		ntk := time.Tick(5 * time.Second)

		for v := range ntk {
			fmt.Println(v)
		}
	*/
	return 0
}

func init() {
	command.AddCommand("time", "time test", "Usage:\n time ", doTime)
}
