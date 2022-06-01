/**
 * @Author: vaynedu
 * @File: transport.go
 * @Date: 2022-06-01 22:39
 * @Description:
 */

package transport

import "context"

const (
	logTimeFormat = "2006-01-02T15:04:05.9999-07:00"
)

type Context interface {
	context.Context
}
