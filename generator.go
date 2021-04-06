package main

import (
	"fmt"
	"github.com/ericnts/code-generator/constant"
	"github.com/ericnts/code-generator/db"
	"github.com/ericnts/tool-box/stringx"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Generator struct {
	Project         string          //项目名称
	ObjectName      string          //对象名称
	TableName       string          //表名
	Comment         string          //描述信息
	Columns         []db.ColumnInfo //数据列
	fileName        string          //文件名
	lowerObjectName string          //变量名称
	columnSize      int             //字段长度
	record          EntityRecord
}

func MkDir() error {
	if err := os.RemoveAll(constant.DistDir); err != nil {
		return err
	}
	dirs := []string{constant.EntityDir, constant.DaoDir, constant.ServiceDir, constant.ControllerDir, constant.VODir, constant.RouterDir}
	for _, dir := range dirs {
		err := os.MkdirAll(fmt.Sprintf("%s/%s", constant.DistDir, dir), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewGenerator(project, tableName string) (p *Generator, err error) {
	p = &Generator{
		Project:   project,
		TableName: tableName,
	}

	p.ObjectName = stringx.CamelName(p.TableName)
	p.fileName = stringx.UnderscoreName(p.ObjectName)
	p.lowerObjectName = stringx.LowerCamelName(p.fileName)
	if comment, err := db.GetTableComment(p.TableName); err != nil {
		return p, err
	} else {
		p.Comment = comment
	}

	if columns, err := db.FindColumn(p.TableName); err != nil {
		return p, err
	} else {
		p.Columns = columns
		p.columnSize = len(columns)
	}
	for i, column := range p.Columns {
		p.record.CheckField(i, column.Name)
	}
	return p, nil
}

// 生成Entity
func (p *Generator) GenerateEntity() error {
	fields := make([]string, p.columnSize+1)
	var importStr string
	if p.record.HasCommon() {
		fields[0] = "\tcommon.Entity\n\n"
	}
	for i, column := range p.Columns {
		if p.record.IsCommonField(i) {
			continue
		}
		var notNull, defaultValue, comment string
		if column.Nullable == "NO" {
			notNull = "not null;"
		}
		if stringx.IsNotBlank(column.DefaultValue) {
			defaultValue = "default:" + column.DefaultValue + ";"
		}
		if stringx.IsNotBlank(column.Comment) {
			comment = "comment:" + column.Comment
		}
		fields[i+1] = fmt.Sprintf("\t%s\t\t%v\t`gorm:\"column:%s;%s%s%s\"`\n",
			stringx.CamelName(column.Name), column.GetType(), column.Name, notNull, defaultValue, comment)
	}
	fieldStr := strings.Join(fields, "")
	if strings.Contains(fieldStr, "time.") {
		importStr += "\n\t\"time\""
	}
	file, err := ioutil.ReadFile(constant.EntityTmp)
	if err != nil {
		return err
	}
	str := string(file)
	str = strings.Replace(str, constant.Package, path.Base(constant.EntityDir), 1)
	str = strings.ReplaceAll(str, constant.Project, p.Project)
	str = strings.Replace(str, constant.Import, importStr, 1)
	str = strings.Replace(str, constant.Comment, p.Comment, 1)
	str = strings.Replace(str, constant.Object, p.ObjectName, 2)
	str = strings.Replace(str, constant.Field, fieldStr, 1)
	str = strings.Replace(str, constant.Table, p.TableName, 1)
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s/%s_%s.go", constant.DistDir, constant.EntityDir, p.fileName, path.Base(constant.EntityDir)), []byte(str), 0644)
	return nil
}

// 生成VO
func (p *Generator) GenerateVO() error {
	var importStr string
	fields := make([]string, p.columnSize)
	toVOs := make([]string, p.columnSize)
	toEntities := make([]string, p.columnSize)
	var hasTime bool
	for i, column := range p.Columns {
		fieldType := column.GetType()
		if fieldType == constant.FTTime {
			hasTime = true
		}
		fields[i] = fmt.Sprintf("\t%s\t\t%v\t`json:\"%s\" form:\"%s\"`\t//%s\n", stringx.CamelName(column.Name), fieldType, stringx.LowerCamelName(column.Name), stringx.LowerCamelName(column.Name), column.Comment)
		if !p.record.IsCommonField(i) {
			toVOs[i] = fmt.Sprintf("\tv.%s = e.%s\n", stringx.CamelName(column.Name), stringx.CamelName(column.Name))
			toEntities[i] = fmt.Sprintf("\t\t%s:\t\t\tp.%s,\n", stringx.CamelName(column.Name), stringx.CamelName(column.Name))
		}
	}
	if hasTime {
		importStr += "\n\t\"time\""
	}
	file, err := ioutil.ReadFile(constant.VOTmp)
	if err != nil {
		return err
	}
	str := string(file)
	str = strings.Replace(str, constant.Package, path.Base(constant.VODir), 1)
	str = strings.Replace(str, constant.Import, importStr, 1)
	str = strings.ReplaceAll(str, constant.Project, p.Project)
	str = strings.ReplaceAll(str, constant.Object, p.ObjectName)
	str = strings.ReplaceAll(str, constant.Field, strings.Join(fields, ""))
	str = strings.ReplaceAll(str, constant.ToVO, strings.Join(toVOs, ""))
	str = strings.ReplaceAll(str, constant.ToEntity, strings.Join(toEntities, ""))
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s/%s_%s.go", constant.DistDir, constant.VODir, p.fileName, path.Base(constant.VODir)), []byte(str), 0644)
	return err
}

// 生成DAO
func (p *Generator) GenerateDAO() error {
	file, err := ioutil.ReadFile(constant.DaoTmp)
	if err != nil {
		return err
	}
	str := string(file)
	str = strings.Replace(str, constant.Package, path.Base(constant.DaoDir), 1)
	str = strings.ReplaceAll(str, constant.Project, p.Project)
	str = strings.ReplaceAll(str, constant.Object, p.ObjectName)
	str = strings.ReplaceAll(str, constant.LowerObject, p.lowerObjectName)
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s/%s_%s.go", constant.DistDir, constant.DaoDir, p.fileName, path.Base(constant.DaoDir)), []byte(str), 0644)
	return err
}

// 生成Service
func (p *Generator) GenerateService() error {
	file, err := ioutil.ReadFile(constant.ServiceTmp)
	if err != nil {
		return err
	}
	str := string(file)
	str = strings.Replace(str, constant.Package, path.Base(constant.ServiceDir), 1)
	str = strings.ReplaceAll(str, constant.Project, p.Project)
	str = strings.ReplaceAll(str, constant.Object, p.ObjectName)
	str = strings.ReplaceAll(str, constant.LowerObject, p.lowerObjectName)
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s/%s_%s.go", constant.DistDir, constant.ServiceDir, p.fileName, path.Base(constant.ServiceDir)), []byte(str), 0644)
	return err
}

// 生成Controller
func (p *Generator) GenerateController() error {
	file, err := ioutil.ReadFile(constant.ControllerTmp)
	if err != nil {
		return err
	}
	str := string(file)
	str = strings.Replace(str, constant.Package, path.Base(constant.ControllerDir), 1)
	str = strings.ReplaceAll(str, constant.Comment, p.Comment)
	str = strings.ReplaceAll(str, constant.Project, p.Project)
	str = strings.ReplaceAll(str, constant.Object, p.ObjectName)
	items := strings.Split(p.fileName, "_")
	if len(items) > 0 {
		items[len(items)-1] = stringx.Plural(items[len(items)-1])
	}
	str = strings.ReplaceAll(str, constant.Router, path.Base(constant.ControllerDir)+"/"+strings.Join(items, "/"))
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s/%s_%s.go", constant.DistDir, constant.ControllerDir, p.fileName, path.Base(constant.ControllerDir)), []byte(str), 0644)
	return err
}

// 生成Router
func (p *Generator) GenerateRouter() error {
	importStr := "\tv1 \"{Project}/controller/v1\"\n\t\"github.com/gin-gonic/gin\""
	file, err := ioutil.ReadFile(constant.RouterTmp)
	if err != nil {
		return err
	}
	str := string(file)
	str = strings.Replace(str, constant.Package, path.Base(constant.RouterDir), 1)
	str = strings.Replace(str, constant.Import, importStr, 1)
	str = strings.ReplaceAll(str, constant.Comment, p.Comment)
	str = strings.ReplaceAll(str, constant.Project, p.Project)
	str = strings.ReplaceAll(str, constant.Object, p.ObjectName)
	str = strings.ReplaceAll(str, constant.Permission, strings.ReplaceAll(p.fileName, "_", ":"))
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s/%s_%s.go", constant.DistDir, constant.RouterDir, p.fileName, path.Base(constant.RouterDir)), []byte(str), 0644)
	return err
}
