package main

import (
	"os"

	"github.com/jhjaggars/dice/pkg/dice"
	"github.com/kataras/iris"
)

func main() {
	f, err := os.OpenFile("access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	app := iris.Default()
	app.Logger().SetOutput(f).SetLevel("info")

	app.Get("/", func(ctx iris.Context) {
		d, _ := dice.ParseDie(ctx.URLParam("dice"))
		ctx.JSON(d.Roll())
	})
	app.Get("/roll/{dice:string}", func(ctx iris.Context) {
		d, _ := dice.ParseDie(ctx.Params().Get("dice"))
		ctx.JSON(d.Roll())
	})
	// listen and serve on http://0.0.0.0:8080.
	app.Run(iris.Addr(":8080"))
}
