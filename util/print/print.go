package printUtil

import log "github.com/sirupsen/logrus"

func Flags(m map[string]interface{}) {
	log.Info("\t---------------------------- flags values ----------------------------")
	for key, value := range m {
		log.Infof("\tKey: %-25s --> %-35v", key, value)
	}
	log.Info("\t----------------------------------------------------------------------")
}