package main

import (
	"github.com/ericnts/code-generator/db"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := MkDir(); err != nil {
		log.Error(err)
	}
	tables, err := db.FindTable("airobot_dev")
	if err != nil {
		return
	}

	for i, table := range tables {
		if i > 5 {
			break
		}
		columns, err := db.FindColumn(table.Name)
		if err != nil {
			continue
		}
		err = GeneratorEntity("test", table, columns)
		if err != nil {
			log.Error(err)
		}
	}
}
