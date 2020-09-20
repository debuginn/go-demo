package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	//i := 0
	//c := cron.New(
	//	cron.WithLogger(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron ", log.Lshortfile))))
	//
	//_, _ = c.AddFunc("@every 30s", func() {
	//	i++
	//	fmt.Printf("task %d\n", i)
	//})
	//
	////time.ParseDuration()
	//
	//c.Start()
	//
	////cron.Parser{}
	//
	//for {
	//	time.Sleep(time.Second)
	//}

	c := cron.New()
	i := 0
	entryId, err := c.AddFunc("1 * * * *", func() {
		fmt.Printf("%d\n", i)
		i++
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", entryId)

	for {
		time.Sleep(time.Second)
	}

	//a, _ := time.ParseDuration("1h30m")
	//
	//fmt.Printf("%v", a)

}

// Minutes：分钟，取值范围[0-59]，支持特殊字符* / , -；
// Hours：小时，取值范围[0-23]，支持特殊字符* / , -；
// Day of month：每月的第几天，取值范围[1-31]，支持特殊字符* / , - ?；
// Month：月，取值范围[1-12]或者使用月份名字缩写[JAN-DEC]，支持特殊字符* / , -；
// Day of week：周历，取值范围[0-6]或名字缩写[JUN-SAT]，支持特殊字符* / , - ?。

// 特殊字符含义如下：
//
// *：使用*的域可以匹配任何值，例如将月份域（第 4 个）设置为*，表示每个月；
// /：用来指定范围的步长，例如将小时域（第 2 个）设置为3-59/15表示第 3 分钟触发，以后每隔 15 分钟触发一次，因此第 2 次触发为第 18 分钟，第 3 次为 33 分钟。。。直到分钟大于 59；
// ,：用来列举一些离散的值和多个范围，例如将周历的域（第 5 个）设置为MON,WED,FRI表示周一、三和五；
// -：用来表示范围，例如将小时的域（第 1 个）设置为9-17表示上午 9 点到下午 17 点（包括 9 和 17）；
// ?：只能用在月历和周历的域中，用来代替*，表示每月/周的任意一天。
