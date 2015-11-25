// basic project main.go
package main

import (
	"github.com/nivance/go-example/basic/logs"
)

func main() {
	logs.Logger.Debug("start log")
	logs.Logger.Debugf("Debugf")
	logs.Logger.Info("hello world")
	logs.Logger.Infof("hahaha")
	logs.Logger.Critical("Critical")
	logs.Logger.Criticalf("Criticalf")
	logs.Logger.Error("error")
	logs.Logger.Errorf("errorf")
	logs.Logger.Trace("trace")
	logs.Logger.Tracef("Tracef")
	logs.Logger.Warn("Warn")
	logs.Logger.Warnf("Warn")
}
