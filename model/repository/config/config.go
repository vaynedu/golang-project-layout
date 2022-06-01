/**
 * @Author: vaynedu
 * @File: config.go
 * @Date: 2022-06-01 22:14
 * @Description:
 */

package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func NewServiceConfig(baseConfigFile string) (*ServiceConfig, error) {
	fileBytes, err := ioutil.ReadFile(baseConfigFile)
	if err != nil {
		return nil, err
	}

	config := new(ServiceConfig)
	if err = yaml.Unmarshal(fileBytes, &config); err != nil {
		panic(err)
	}

	return config, nil

}

type ServiceConfig struct {
	PSM *string `yaml:"PSM"`
}
