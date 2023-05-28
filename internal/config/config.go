package config

import (
	"enceremony-be/internal/config/aws_conf"
	"encoding/json"
	"fmt"
	"log"
)

type Config struct {
	Aws string
}

func NewConfig() (*Config, error) {
	var config Config
	var confStr string
	var err error

	envKey := "enceremony"

	secretName := fmt.Sprintf("prod/%s/all", envKey)

	confSvc := aws_conf.NewConfigurationService()
	if confStr, err = confSvc.GetConfigString(envKey, secretName); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(confStr), &config); err != nil {
		log.Fatalf("Couldn't unmarshal json => %s to conf. Error %s", confStr, err)
		return nil, err
	}

	// use a single instance of Validate, it caches struct info
	//validatorSvc := validator.New()
	//
	//err = validatorSvc.Struct(config)
	//if err != nil {
	//	for _, err := range err.(validator.ValidationErrors) {
	//		fmt.Println(err)
	//	}
	//	return nil, err
	//}

	return &config, nil
}
