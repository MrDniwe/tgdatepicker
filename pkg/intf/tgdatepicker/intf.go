package tgdatepicker

import (
	"context"
	"time"

	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DatePicker interface {
	PickDate(
		ctx context.Context,
		chatID, messageID int64,
		currentDate *time.Time,
		payload []string,
	) (bot.Chattable, error)
}
