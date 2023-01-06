package middleware

import (
	"beide/framework"
	"fmt"
)

func Test1() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test1")
		err := c.Next()
		if err != nil {
			return err
		}
		fmt.Println("middleware post test1")
		return nil
	}
}

func Test2() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test2")
		err := c.Next()
		if err != nil {
			return err
		}
		fmt.Println("middleware post test2")
		return nil
	}
}

func Test3() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test3")
		err := c.Next()
		if err != nil {
			return err
		}
		fmt.Println("middleware post test3")
		return nil
	}
}
