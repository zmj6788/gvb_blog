package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)
func inner (name string) func() {
	return func() {
		fmt.Printf("%s %s \n", name, time.Now())
	}
}

type Job struct {
	Name string
}
func (j Job) Run() {
	fmt.Printf("%s 这是一个job方法 %s\n", j.Name, time.Now())
}
func main() {
	Cron := cron.New(cron.WithSeconds())
	// Cron.AddFunc("* * * * * *", func() {
	// 	fmt.Println("每秒执行一次", time.Now())
	// })
	// 闭包
	// Cron.AddFunc("*/5 * * * * *", inner("zmj"))
	// 对象
	Cron.AddJob("*/2 * * * * *", Job{Name: "zmj"})
	Cron.Start()
	select {}
}
