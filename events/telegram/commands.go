package telegram

import (
	e "TelegramBot/lib/error"
	"TelegramBot/storage"
	"errors"
	"log"
	"net/url"
	"strings"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	//add page: http://
	// rnd page: /rnd
	// help page:/help
	// start: /start: hi+help

	if isAddCmd(text) {
		return p.savePage(chatID, text, username)
	}

	switch text {
	case RndCmd:
		return p.sendRandom(chatID, username)
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID)
	default:
		return p.tg.SendMessages(chatID, msgUnknownCommand)

	}
}
func (p *Processor) savePage(chatID int, pageURL string, username string) (err error) {
	defer func() { err = e.WrapIfErr("cant do cmd save page", err) }()

	page := &storage.Page{
		URL:      pageURL,
		UserName: username,
	}
	IsExist, err := p.storage.IsExists(page)
	if err != nil {
		return err
	}
	if IsExist {

		return p.tg.SendMessages(chatID, msgAlreadyExists)

	}

	if err := p.storage.Save(page); err != nil {
		return err
	}
	if err := p.tg.SendMessages(chatID, msgSaved); err != nil {
		return err
	}
	return nil
}
func (p *Processor) sendRandom(chatID int, username string) (err error) {
	defer func() { err = e.WrapIfErr("cant do command: cant send random", err) }()
	page, err := p.storage.PickRandom(username)
	if err != nil && !errors.Is(err, storage.ErrNoSavedPages) {

		return err
	}
	if errors.Is(err, storage.ErrNoSavedPages) {
		return p.tg.SendMessages(chatID, msgNoSavedPages)
	}
	if err := p.tg.SendMessages(chatID, page.URL); err != nil {
		return err
	}
	return p.storage.Remove(page)
}
func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessages(chatID, msgHelp)
}
func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessages(chatID, msgHello)
}
func isAddCmd(text string) bool {
	if strings.HasPrefix(text, "Http") {
		return isUrl(text)
	} else if strings.HasPrefix(text, "http") {
		return isUrl(text)
	}
	return false
}

func isUrl(text string) bool {
	u, err := url.Parse(text)
	log.Print(u)
	return err == nil && u.Host != " "
}
