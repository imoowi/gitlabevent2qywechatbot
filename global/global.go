/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package global

import (
	"gitlab2wechatbot/apps/bot/models"

	"go.uber.org/dig"
)

var DigContainer *dig.Container
var MigrateContainer *dig.Container

var DigContainerProviders []any
var DigContainerMigrateProviders []any

func init() {
	DigContainer = dig.New()
	MigrateContainer = dig.New()
	DigContainerProviders = make([]any, 0)
	DigContainerMigrateProviders = make([]any, 0)
}
func RegisterContainerProviders(provider any) {
	DigContainerProviders = append(DigContainerProviders, provider)
}
func RegisterMigrateContainerProviders(provider any) {
	DigContainerMigrateProviders = append(DigContainerMigrateProviders, provider)
}
func Bootstrap() {
	InitMsgChan()
}

var MsgChan chan *models.Msg

func InitMsgChan() {
	MsgChan = make(chan *models.Msg, 1000)
}
