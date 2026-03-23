package main

import "log"

func main()  {
	cfg := config{
		addr: ":8080",
	}

	app := &application{
		confg: cfg,
	}

	log.Fatal(app.run())
}