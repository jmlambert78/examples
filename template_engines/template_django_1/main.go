package main

import (
	"github.com/kataras/go-template/django"
	"github.com/kataras/iris"
)

func main() {
	iris.UseTemplate(django.New()).Directory("./templates", ".html")

	iris.Get("/", func(ctx *iris.Context) {
		ctx.MustRender("mypage.html", map[string]interface{}{"username": "iris", "is_admin": true}, iris.RenderOptions{"gzip": true})
	})

	iris.Listen(":8080")
}
