package main

import (
	"fmt"

	"github.com/jinghui0108/golib/jlog"
	"github.com/jinghui0108/golib/jtime"
)

func main() {

	t := jtime.New()


	fmt.Println(t.GetDateTime())

	logH := jlog.New(true, "bench_")

	if !logH.Init() {
		return
	}

	for i := 1000000; i > 0; i-- {
		logH.Debug(fmt.Sprintf("Test:%d", i))
	}

	fmt.Println(t.GetDateTime())



}
