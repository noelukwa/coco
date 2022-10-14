package main

import (
	"fmt"

	"github.com/noelukwa/coco"
)

func main() {
	app := coco.NewApp()

	app.Get("/hey", log, func(rw coco.Response, r *coco.Request, _ coco.NextFunc) {
		rw.JSON(200, map[string]string{"hello": "world"})
	})

	app.Get("/hey/:id", log, func(rw coco.Response, r *coco.Request, _ coco.NextFunc) {

		rw.JSON(200, map[string]string{"message": "hello " + r.Params()["id"]})
	})

	users := app.NewRoute("/users")

	users.Use(userLog)

	users.Get("", log, func(rw coco.Response, r *coco.Request, _ coco.NextFunc) {

		rw.JSON(200, map[string]string{"hello": "world!"})
	})

	users.Get("/", log, func(rw coco.Response, r *coco.Request, next coco.NextFunc) {
		rw.JSON(200, map[string]string{"hello": r.Params()["name"]})
	})

	app.Listen(":8080")
}

func log(rw coco.Response, r *coco.Request, next coco.NextFunc) {
	fmt.Println("enter logging")
	next(rw, r)
	fmt.Println("exit logging")
}

func userLog(rw coco.Response, r *coco.Request, next coco.NextFunc) {
	fmt.Println("enter  user logger")
	next(rw, r)
	fmt.Println("exiting user logger")
}
