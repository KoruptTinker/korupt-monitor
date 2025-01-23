package main

func main() {
	engine := InitClient()

	engine.Start()

	select {}
}
