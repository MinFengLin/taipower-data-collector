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

	taipower "github.com/MinFengLin/goTaipower/taipower"

	bot "github.com/MinFengLin/goTaipower/bot"
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

	more_info, _ := strconv.ParseBool(os.Getenv("MORE_INFO"))
	Env_Time, _ := strconv.ParseInt(os.Getenv("Timer_Minutes"), 10, 64)
	crontab_time := os.Getenv("Crontime")

	if len(crontab_time) > 0 {
		fmt.Printf("chat_id: %d, Token: %s, more_info: %t, Use crontab: %s \n", chat_id, Token, more_info, crontab_time)
		c := cron.New()
		_, _ = c.AddFunc(crontab_time, func() {
			data := taipower.Parser_Taipower(&more_info)
			bot.Telegram_bot_run(&chat_id, &Token, data)
		})
		data := taipower.Parser_Taipower(&more_info)
		bot.Telegram_bot_run(&chat_id, &Token, data)
		c.Start()
		select {}
	} else {
		fmt.Printf("chat_id: %d, Token: %s, more_info: %t, How often to run: %d (Minutes)\n", chat_id, Token, more_info, Env_Time)
		tChannel := time.NewTimer(time.Duration(Env_Time) * time.Minute)
		for {
			data := taipower.Parser_Taipower(&more_info)
			bot.Telegram_bot_run(&chat_id, &Token, data)
			tChannel.Reset(time.Duration(Env_Time) * time.Minute)
			<-tChannel.C
		}
	}
}
