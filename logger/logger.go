package logger

import (
	log "github.com/sirupsen/logrus"
)

func Log(msg string) {
	log.Info("[Custom Info Msg]"+msg)
}
