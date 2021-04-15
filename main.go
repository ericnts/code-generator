package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := MkDir(); err != nil {
		log.Error(err)
		return
	}
	generate("epshealth-airobot-admin", "robot_position")
}

func generate(project, table string) {
	generator, err := NewGenerator(project, table)
	if err != nil {
		log.Error(err)
		return
	}
	if err := generator.GenerateEntity(); err != nil {
		log.Error(err)
	}
	if generator.record.HasCommon() {
		if err := generator.GenerateVO(); err != nil {
			log.Error(err)
		}
		if err := generator.GenerateController(); err != nil {
			log.Error(err)
		}
		if err := generator.GenerateRouter(); err != nil {
			log.Error(err)
		}
		if err := generator.GenerateService(); err != nil {
			log.Error(err)
		}
		if err := generator.GenerateDAO(); err != nil {
			log.Error(err)
		}
	}
}
