package main

import "beide/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
