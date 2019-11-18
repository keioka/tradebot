package models

import (
	"fmt"
)

type Notification struct {
}


type Exchange struct {}

type Strategy struct {
	name string
	symbol string
	tradingviewAlert string
	fixedLot bool
	fixedLotSize float64
	fixableLevarage float64
	fixableMaxSize float64
	sfdEnable bool
}

type Runner struct {
	strategy Strategy
}

func execute(runner Runner) {
	fmt.Println("Strategy is running")
}

type Config struct {
	apiKey
	apiSecret
	tradingviewSecret
	lineNotificationToken
}