package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

func main() {
	tables, err := FindTable("airobot_dev")
	if err != nil {
		return
	}

	var sb strings.Builder
	for i, table := range tables {
		columns, err := FindColumn(table.Name)
		if err != nil {
			continue
		}
		sb.WriteString(fmt.Sprintf("%v. %s(%s) [", i, table.Name, table.Comment))
		for _, column := range columns {
			sb.WriteString(column.Name + ",")
		}
		sb.WriteString(fmt.Sprintf("]\n"))
	}
	log.Info(sb.String())
}
