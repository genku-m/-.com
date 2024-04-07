package main

import "github.com/genku-m/upsider-cording-test/server"

func main() {
	cfg := server.NewConfig()
	svr := server.NewServer(nil, cfg)
	svr.Listen()
}
