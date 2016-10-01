package main

import (
	"flag"
	"math/rand"
	"time"

	_ "github.com/fsr/matthias/plugins/apbdoor"
	_ "github.com/fsr/matthias/plugins/bday"
	_ "github.com/fsr/matthias/plugins/btc"
	_ "github.com/fsr/matthias/plugins/drsommer"
	_ "github.com/fsr/matthias/plugins/dvb"
	_ "github.com/fsr/matthias/plugins/elbe"
	_ "github.com/fsr/matthias/plugins/firat"
	_ "github.com/fsr/matthias/plugins/fsr"
	_ "github.com/fsr/matthias/plugins/googleimages"
	_ "github.com/fsr/matthias/plugins/help"
	_ "github.com/fsr/matthias/plugins/maintain"
	_ "github.com/fsr/matthias/plugins/mensa"
	_ "github.com/fsr/matthias/plugins/ping"
	_ "github.com/fsr/matthias/plugins/porn"
	_ "github.com/fsr/matthias/plugins/portalstate"
	_ "github.com/fsr/matthias/plugins/random"
	_ "github.com/fsr/matthias/plugins/simpleresponses"
	_ "github.com/fsr/matthias/plugins/stoll"
	_ "github.com/fsr/matthias/plugins/urbandictionary"
	_ "github.com/fsr/matthias/plugins/version"

	"github.com/abourget/slick"
)

var isProduction = flag.Bool("production", false, "production mode, uses matthias.conf")

func main() {
	flag.Parse()

	configFile := "./matthias_dev.conf"

	if *isProduction {
		configFile = "./matthias.conf"
	}

	bot := slick.New(configFile)
	bot.Run()
}

func init() {
	rand.Seed(time.Now().UnixNano())
}