package main

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {

	var (
		// Universal markup builders.
		menu     = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
		// selector = &tb.ReplyMarkup{}

		// Reply buttons.
		btnHelp     = menu.Text("ℹ Help")
		btnSettings = menu.Text("⚙ Settings")

		// Inline buttons.
		//
		// Pressing it will cause the client to
		// send the bot a callback.
		//
		// Make sure Unique stays unique as per button kind,
		// as it has to be for callback routing to work.
		//
		// btnPrev = selector.Data("⬅", "prev", ...)
		// btnNext = selector.Data("➡", "next", ...)
	)

	menu.Reply(
		menu.Row(btnHelp),
		menu.Row(btnSettings),
	)

	b, err := tb.NewBot(tb.Settings{
		Token:  Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(m *tb.Message) {
                b.Send(m.Sender, "Hello World!", menu)
        })

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hello World!")
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		log.Println(m.Text)
		b.Send(m.Sender, m.Text)
	})


	b.Start()
}

