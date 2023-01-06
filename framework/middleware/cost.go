package middleware

import (
	"beide/framework"
	"log"
	"time"
)

func Cost() framework.ControllerHandler {
	return func(c *framework.Context) error {
		// 记录开始时间
		start := time.Now()

		// 使用next执行具体的业务逻辑
		err := c.Next()
		if err != nil {
			return err
		}

		// 记录结束的时间
		end := time.Now()
		cost := end.Sub(start)
		log.Printf("api uri: %v, cost: %v", c.GetRequest().RequestURI, cost.Seconds())

		return nil
	}
}
