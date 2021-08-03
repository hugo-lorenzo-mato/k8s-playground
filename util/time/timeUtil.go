package timeUtil

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func FromStart(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Infof("%s took %s", name, elapsed)
}