package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Программа запущена")
	log.Warn("Внимание")
	log.Error("Произошла ошибка")
}
