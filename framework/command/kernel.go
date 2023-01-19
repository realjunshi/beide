package command

import "beide/framework/cobra"

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	// app 命令
	root.AddCommand(initAppCommand())
	// env 命令
	root.AddCommand(initEnvCommand())
	// cron 命令
	root.AddCommand(initCronCommand())
	// config 命令
	root.AddCommand(initConfigCommand())
	// build 命令
	root.AddCommand(initBuildCommand())
	// go build
	root.AddCommand(goCommand)
	// npm build
	root.AddCommand(npmCommand)
	// dev
	root.AddCommand(initDevCommand())
	// provider
	root.AddCommand(initProviderCommand())
	// demo
	root.AddCommand(DemoCommand)
	// cmd
	root.AddCommand(initCmdCommand())
	// middleware
	root.AddCommand(initMiddlewareCommand())

	// app
	//root.AddCommand(initAppCommand())
	//root.AddCommand(DemoCommand)
	// env
	//root.AddCommand(initEnvCommand())
	//root.AddCommand(deployCommand)
	//
	// cron
	//root.AddCommand(initCronCommand())
	//// cmd
	//cmdCommand.AddCommand(cmdListCommand)
	//cmdCommand.AddCommand(cmdCreateCommand)
	//root.AddCommand(cmdCommand)
	//
	// build
	//buildCommand.AddCommand(buildSelfCommand)
	//buildCommand.AddCommand(buildBackendCommand)
	//buildCommand.AddCommand(buildFrontendCommand)
	//buildCommand.AddCommand(buildAllCommand)
	//root.AddCommand(buildCommand)
	//

	//
	//// dev
	//root.AddCommand(initDevCommand())
	//
	//// middleware
	//middlewareCommand.AddCommand(middlewareAllCommand)
	//middlewareCommand.AddCommand(middlewareAddCommand)
	//middlewareCommand.AddCommand(middlewareRemoveCommand)
	//root.AddCommand(middlewareCommand)
	//
	//// swagger
	//swagger.IndexCommand.AddCommand(swagger.InitServeCommand())
	//swagger.IndexCommand.AddCommand(swagger.GenCommand)
	//root.AddCommand(swagger.IndexCommand)
	//
	//// provider
	//providerCommand.AddCommand(providerListCommand)
	//providerCommand.AddCommand(providerCreateCommand)
	//root.AddCommand(providerCommand)
	//
	//// new
	//root.AddCommand(initNewCommand())
}
