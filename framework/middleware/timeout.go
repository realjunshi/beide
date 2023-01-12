package middleware

import (
	"beide/framework/gin"
	"context"
	"fmt"
	"log"
	"time"
)

func Timeout(d time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
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

			c.Next()

			finish <- struct{}{}
		}()

		// 执行业务逻辑后操作
		select {
		case p := <-panicChan:
			c.ISetStatus(500).IJson("time out")
			log.Println(p)

		case <-finish:
			fmt.Println("finish")

		case <-durationCtx.Done():
			c.ISetStatus(500).IJson("time out")
		}
	}
}
