package main

import (
	"fmt"

	"github.com/noelukwa/coco"
)

func main() {
	app := coco.NewApp()

	app.Get("hey", log, func(rw coco.Response, r *coco.Request, _ coco.NextFunc) {
		fmt.Println("enter handler")
		rw.Write([]byte("hey world\n"))
		fmt.Println("exit handler")
	})

	app.Get("hey/:id", log, func(rw coco.Response, r *coco.Request, _ coco.NextFunc) {

		rw.Write([]byte(fmt.Sprintf("hey world %s\n", r.Params()["id"])))
	})

	users := app.NewRoute("/users")

	users.Use(userLog)

	users.Get("", log, func(rw coco.Response, r *coco.Request, _ coco.NextFunc) {
		fmt.Println("enter index handler")
		rw.Write([]byte("users index\n"))
		fmt.Println("exit index handler")
	})

	users.Get(":name", log, func(rw coco.Response, r *coco.Request, next coco.NextFunc) {

		name := r.Params()["name"]
		greeting := fmt.Sprintf("Hello %s\n", name)

		rw.Write([]byte(greeting))
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
