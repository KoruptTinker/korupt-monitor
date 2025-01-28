package main

import "fmt"

func main() {
	engine := InitClient()
	fmt.Println("Starting Cron Engine")
	engine.Start()
	fmt.Println("Started Cron Engine")
	select {}
}
