package main

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type config struct {
	Host    string `env:"HOST" envDefault:"localhost"`
	Port    int    `env:"PORT" envDefault:"3000"`
	Address string `env:"ADDRESS,expand" envDefault:"$HOST:${PORT}"`
}

func main() {
	var c config

	godotenv.Load()

	if err := env.Parse(&c); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Println(c.Host)
	fmt.Println(c.Port)
	// now do something with s3 or whatever
}
