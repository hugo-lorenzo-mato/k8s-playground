package logUtil

import (
	fileUtil "k8s-playground/util/file"
	"reflect"
)

type LogFiles struct {
	WrongNps   string
	CreatedNps string
	DeletedNps string
}

func CreateLogFiles(files LogFiles) error {
	f := reflect.ValueOf(files)
	for i := 0; i < f.NumField(); i++ {
		if fileUtil.Exists(f.Field(i).String()) {
			err := fileUtil.DeleteFile(f.Field(i).String())
			if err != nil {
				return err
			}
		}
	}
	return nil
}