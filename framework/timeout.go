package framework

//func TimeoutHandler(fun ControllerHandler, d time.Time) ControllerHandler {
//	return func(c *Context) error {
//
//		finish := make(chan struct{}, 1)
//		panicChan := make(chan interface{}, 1)
//
//		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
//		defer cancel()
//
//		c.request.WithContext(durationCtx)
//
//		go func() {
//			defer func() {
//				if p := recover(); p != nil {
//					panicChan <- p
//				}
//			}()
//
//			// 处理具体的业务
//
//			finish <- struct{}{}
//		}()
//
//		select {
//		case p := <-panicChan:
//			log.Println(p)
//			c.responseWriter.WriteHeader(500)
//		}
//	}
//}
