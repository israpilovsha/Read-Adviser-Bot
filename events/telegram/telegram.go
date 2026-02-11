package telegram

import "github.com/israpilovsha/Read-Adviser-Bot/clients/telegram"

type Processor struct {
	tg *telegram.Client
	offset int
	//storage
}

func New(client *telegram.Client, storage) {

}  