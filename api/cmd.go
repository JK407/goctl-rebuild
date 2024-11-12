package api

import (
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"gogen/api/apigen"
	"gogen/api/gogen"
	"gogen/internal/cobrax"
)

var (
	// Cmd describes an api command.
	Cmd   = cobrax.NewCommand("api", cobrax.WithRunE(apigen.CreateApiTemplate))
	goCmd = cobrax.NewCommand("go", cobrax.WithRunE(gogen.GoCommand))
)

func init() {
	var (
		apiCmdFlags = Cmd.Flags()
		goCmdFlags  = goCmd.Flags()
	)

	apiCmdFlags.StringVar(&apigen.VarStringOutput, "o")
	apiCmdFlags.StringVar(&apigen.VarStringHome, "home")
	apiCmdFlags.StringVar(&apigen.VarStringRemote, "remote")
	apiCmdFlags.StringVar(&apigen.VarStringBranch, "branch")

	goCmdFlags.StringVar(&gogen.VarStringDir, "dir")
	goCmdFlags.StringVar(&gogen.VarStringAPI, "api")
	goCmdFlags.StringVar(&gogen.VarStringHome, "home")
	goCmdFlags.StringVar(&gogen.VarStringRemote, "remote")
	goCmdFlags.StringVar(&gogen.VarStringBranch, "branch")
	goCmdFlags.StringVarWithDefaultValue(&gogen.VarStringStyle, "style", config.DefaultFormat)

	// Add sub-commands
	Cmd.AddCommand(goCmd)
}
