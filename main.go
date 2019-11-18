package main

import (
	"fmt"
	"os"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"github.com/joho/godotenv"
	"tradebot-go/models"
	// "github.com/Darkbladecr/ccxt"
)

func main() {
	err := godotenv.Load()
	if err != nil {
    log.Error(err)
  } else {
		log.Info("Load env file")
	}

	fmt.Println(os.Getenv("API_KEY"))

	// bitFlyer := bitFlyer{
	// 	sfdEnable: "hello",
	// }
	cfg, err := ini.Load("./strategy/bitflyer_st7.ini")
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Load ini file")
	}

	strategy := Strategy{
		exchange: cfg.Section("settings").Key("EXCHANGE").String(),
	}

	// fmt.Println(strategy)

	fmt.Println(cfg.Section("settings").Key("EXCHANGE").String())

	log.Info("Trade Bot Started...")
	fmt.Println("hello")
}
