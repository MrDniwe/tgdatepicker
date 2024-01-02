package tgdatepicker

import (
	"context"
	"time"

	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mrdniwe/tgdatepicker/pkg/intf/tgdatepicker"
)

type datePicker struct {
}

func New() tgdatepicker.DatePicker {
	return &datePicker{}
}

func (dp *datePicker) PickDate(
	ctx context.Context,
	chatID, messageID int64,
	currentDate *time.Time,
	payload []string,
) (bot.Chattable, error) {
	var resp bot.Chattable
	return resp, nil
}
