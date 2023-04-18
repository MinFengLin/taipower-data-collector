package main

import (
	"fmt"
	"strconv"

	/*
	 * env
	 */
	"log"
	"os"
	"time"

	bot "github.com/MinFengLin/goTaipower/bot"
	taipower "github.com/MinFengLin/goTaipower/taipower"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	chat_id, _ := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)
	Token := os.Getenv("TOKEN")
	Env_Time, _ := strconv.ParseInt(os.Getenv("Timer_Minutes"), 10, 64)
	crontab_time := os.Getenv("Crontime")
	more_info := true

	if len(crontab_time) > 0 {
		fmt.Printf("chat_id: %d, Token: %s, Use crontab: %s \n", chat_id, Token, crontab_time)
		c := cron.New()
		_, _ = c.AddFunc(crontab_time, func() {
			data := taipower.Taipower_res(more_info)
			bot.Telegram_bot_run(&chat_id, &Token, data)
		})
		data := taipower.Taipower_res(more_info)
		bot.Telegram_bot_run(&chat_id, &Token, data)
		c.Start()
		select {}
	} else {
		fmt.Printf("chat_id: %d, Token: %s, How often to run: %d (Minutes)\n", chat_id, Token, Env_Time)
		tChannel := time.NewTimer(time.Duration(Env_Time) * time.Minute)
		for {
			data := taipower.Taipower_res(more_info)
			bot.Telegram_bot_run(&chat_id, &Token, data)
			tChannel.Reset(time.Duration(Env_Time) * time.Minute)
			<-tChannel.C
		}
	}
}
