package main

import (
	"github.com/zeromicro/go-zero/core/load"
	"github.com/zeromicro/go-zero/core/logx"
	"gogen/cmd"
)

func main() {
	logx.Disable()
	load.Disable()
	cmd.Execute()
}
