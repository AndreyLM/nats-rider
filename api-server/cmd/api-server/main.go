package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	showHelp     bool
	showVersion  bool
	serverListen string
	natsServers  string
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: api-server [options...]\n\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
	}
	flag.BoolVar(&showHelp, "help", false, "Show help")
	flag.BoolVar(&showVersion, "version", false, "Show version")
	flag.StringVar(&serverListen, "listen", "0.0.0.0:9090", "Network host:port to listen to")
	flag.StringVar(&natsServers, "nats", nats.DefaultURL, "Network host:port to listen to")
}

func main() {

}
