package middleware

import "beide/framework"

func Recovery() framework.ControllerHandler {
	// 使用函数回掉
	return func(c *framework.Context) error {
		// 核心在增加这个recover机制，捕获c.Next()出现的panic
		defer func() {
			if err := recover(); err != nil {
				c.SetStatus(500).Json(err)
			}
		}()

		err := c.Next()
		if err != nil {
			return err
		}

		return nil
	}
}
