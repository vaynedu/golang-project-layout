/**
 * @Author: vaynedu
 * @File: providers.go
 * @Date: 2022-06-01 22:11
 * @Description:
 */

package container

import (
	"github.com/google/wire"
	"golang-project-layout/model/repository/config"
	"golang-project-layout/model/repository/once"
)

var onceSet = wire.NewSet(
	config.NewServiceConfig,
	once.NewOnce,
)
