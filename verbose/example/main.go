package main

import "github.com/koykov/helpers/verbose"

func main() {
	v := verbose.NewVerbose(verbose.LevelDebug3)
	v.Info("default")
	v.Ok("success message")
	v.Warning("warning message")
	v.Fail("fail message")
	v.Debug1("message")
	v.Debug2("message")
	v.Debug3("message")
}
