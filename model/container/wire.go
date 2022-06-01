//go:build wireinject
// +build wireinject

/**
 * @Author: vaynedu
 * @File: wire.go
 * @Date: 2022-06-01 22:11
 * @Description:
 */

package container

import (
	"github.com/google/wire"
	"golang-project-layout/model/repository/once"
)

type Container struct {
	// once config
	_forbidden once.IOnce
}

func Init(baseConfigFile string) (*Container, error) {
	wire.Build(
		wire.Struct(new(Container), "*"),
		onceSet,
	)

	return &Container{}, nil
}