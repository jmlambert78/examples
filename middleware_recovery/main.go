package main

import (
	"github.com/iris-contrib/middleware/recovery"
	"github.com/kataras/iris"
)

func main() {
	//iris.Use(recovery.New(os.Stdout)) // this is an optional parameter, you can skip it, the default is os.Stderr
	iris.Use(recovery.New())
	i := 0
	iris.Get("/", func(ctx *iris.Context) {
		i++
		if i%2 == 0 {
			panic("a panic here")
			return

		}

		ctx.Next()

	}, func(ctx *iris.Context) {
		ctx.Writef("Hello, refresh one time more to get panic!")
	})

	iris.Listen(":8080")
}
