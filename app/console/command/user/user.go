package user

import (
	"fmt"

	"beide/framework/cobra"
)

var UserCommand = &cobra.Command{
	Use:   "user",
	Short: "user",
	RunE: func(c *cobra.Command, args []string) error {
		//container := c.GetContainer()
		fmt.Println("测试一下命令工具")
		return nil
	},
}
