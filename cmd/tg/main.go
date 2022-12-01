package main

import (
	"deadlines/internal/repo"
	"deadlines/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	token := ""
	bot, err := tgbotapi.NewBotAPI(token)
	bot.Debug = true
	if err != nil {
		panic(err)
	}

	psql, err := initRepo("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	cronPsql, err := initCron("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	tgService := services.NewReminderBot(bot, psql)
	cronService := services.NewReminderCron(bot, cronPsql)
	cronService.Start()

	log.Printf("Authorized on account %s", bot.Self.UserName)
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		tgService.HandleUpdates(&update)
	}
}

func initRepo(addr string) (repo.IReminderRepo, error) {
	psql, err := sqlx.Connect("postgres", addr)
	if err != nil {
		return nil, err
	}

	return repo.NewReminderRepo(psql), nil
}

func initCron(addr string) (repo.IReminderRepoCron, error) {
	psql, err := sqlx.Connect("postgres", addr)
	if err != nil {
		return nil, err
	}

	return repo.NewReminderRepoCron(psql), nil
}
