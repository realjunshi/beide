package middleware

import (
	"beide/framework"
	"context"
	"fmt"
	"log"
	"time"
)

func Timeout(d time.Duration) framework.ControllerHandler {
	return func(c *framework.Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()

			err := c.Next()
			if err != nil {
				return
			}

			finish <- struct{}{}
		}()

		// 执行业务逻辑后操作
		select {
		case p := <-panicChan:
			err := c.Json(500, "time out")
			if err != nil {
				return err
			}
			log.Println(p)

		case <-finish:
			fmt.Println("finish")

		case <-durationCtx.Done():
			c.SetHasTimeout()
			err := c.Json(500, "time out")
			if err != nil {
				return err
			}
		}
		return nil
	}
}
