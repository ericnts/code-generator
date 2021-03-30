package main

import (
	"fmt"
	"github.com/ericnts/code-generator/constant"
	"github.com/ericnts/code-generator/db"
	"github.com/ericnts/tool-box/stringx"
	"io/ioutil"
	"os"
	"strings"
)

const (
	EntityFile     = "template/entity.tmp"
	DaoFile        = "template/dao.go"
	ServiceFile    = "template/service.go"
	ControllerFile = "template/controller.go"
)
const (
	DistDir       = "dist"
	EntityDir     = "entity"
	DaoDir        = "dao"
	ServiceDir    = "service"
	ControllerDir = "controller"
)

const (
	Package = "{Package}"
	Import  = "{Import}"
	Comment = "{Comment}"
	Entity  = "{Entity}"
	Field   = "{Field}"
	Table   = "{Table}"
)

func MkDir() error {
	if err := os.RemoveAll(DistDir); err != nil {
		return err
	}
	dirs := [4]string{EntityDir, DaoDir, ServiceDir, ControllerDir}
	for _, dir := range dirs {
		err := os.MkdirAll(fmt.Sprintf("%s/%s", DistDir, dir), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func GeneratorEntity(project string, table db.TableInfo, columns []db.ColumnInfo) error {
	var record EntityRecord
	fields := make([]string, len(columns))
	var hasTime bool
	for i, column := range columns {
		record.AddField(column.Name, i)
		fieldType := column.GetType()
		if fieldType == constant.FTTime {
			hasTime = true
		}
		fields[i] = fmt.Sprintf("\t%s\t\t%v\t\t//%s\n", stringx.CamelName(column.Name), fieldType, column.Comment)
	}
	var importStr string
	if record.HasBase() {
		importStr += fmt.Sprintf("\t\"%s/common\"\n", project)
		for _, index := range record.FieldIndex {
			fields[index] = ""
		}
		fields = append([]string{"\tcommon.Entity\n\n"}, fields...)
	}
	if hasTime {
		importStr += "\t\"time\""
	}
	file, err := ioutil.ReadFile(EntityFile)
	if err != nil {
		return err
	}
	str := string(file)
	str = strings.Replace(str, Package, EntityDir, 1)
	str = strings.Replace(str, Import, importStr, 1)
	str = strings.Replace(str, Comment, table.Comment, 1)
	str = strings.Replace(str, Entity, stringx.CamelName(table.Name), 2)
	str = strings.Replace(str, Field, strings.Join(fields, ""), 1)
	str = strings.Replace(str, Table, table.Name, 1)
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s/%s.go", DistDir, EntityDir, table.Name), []byte(str), 0644)
	return err
}
