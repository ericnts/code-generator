package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := MkDir(); err != nil {
		log.Error(err)
		return
	}
	generator, err := NewGenerator("epshealth-airobot-admin", "sys_dict")
	if err != nil {
		log.Error(err)
		return
	}
	if err := generator.GenerateEntity(); err != nil {
		log.Error(err)
	}
	if err := generator.GenerateVO(); err != nil {
		log.Error(err)
	}
}
