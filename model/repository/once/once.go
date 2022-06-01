/**
 * @Author: vaynedu
 * @File: once.go
 * @Date: 2022-06-01 22:23
 * @Description:
 */

package once

import "golang-project-layout/model/repository/config"

type IOnce interface {
}

type _once struct {
}

func NewOnce(c *config.ServiceConfig) (IOnce, error) {
	return &_once{}, nil
}
