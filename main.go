package main

import (
	"log"

	"github.com/Izunna-Norbert/busha-practice/initializers"
	"github.com/Izunna-Norbert/busha-practice/routers"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectDb()
	initializers.ConnectRedis()
}

func main() {
	r := routers.Routes()

	log.Fatalf("%v", r.Run())
}
