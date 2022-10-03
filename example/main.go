package main

import (
	"fmt"
	"log"

	"github.com/noelukwa/coco"
)

func main() {
	app := coco.NewApp()

	// app.Use(logger, timer /*, ... */)

	// app.LoadTemplate("templates/*")
	// app.ServeStaticFrom("public")

	// app.Get("/", func(res http.ResponseWriter, req *http.Request) {
	// 	// do something
	// 	res.JSON(200, map[string]string{"message": "Hello World"})
	// })

	// bookRoute := app.NewRoute("/books")
	// bookRoute.Use(auth)
	// bookRoute.Get("/", func(res http.ResponseWriter, req *http.Request) {

	// 	// do something
	// 	res.Render("books/index.html", map[string]string{"message": "Hello World"})
	// })

	// bookRoute.Get("/:id", func(res http.ResponseWriter, req *http.Request) {
	// 	// do something
	// }).Use(rate)
	// app.Use(middlewareOne, middlewareTwo)

	app.Get("/hey", func(rw coco.Response, r *coco.Request, _ coco.NextFunc) {
		log.Println("final route")
		rw.Write([]byte("hey world\n"))
	})

	users := app.NewRoute("/users")

	users.Get("/", func(rw coco.Response, r *coco.Request, _ coco.NextFunc) {
		log.Println("users index")
		rw.Write([]byte("users index\n"))
	})

	users.Get("/:name", func(rw coco.Response, r *coco.Request, next coco.NextFunc) {
		log.Println("entering users route")
		name := r.Params()["name"]
		greeting := fmt.Sprintf("Hello %s\n", name)
		next(rw, r)
		rw.Write([]byte(greeting))
	})

	app.Listen(":8080")

}

// func one(rw http.ResponseWriter, r *http.Request, next coco.NextFunc) {
// 	log.Println("one")
// 	next(rw, r)
// }

// func two(rw http.ResponseWriter, r *http.Request, next coco.NextFunc) {
// 	log.Println("two")
// 	next(rw, r)
// }

// func logger(rw http.ResponseWriter, r *http.Request, next *coco.NextFunc) {
// 	log.Println("entering logger")
// 	next(rw, r)
// 	log.Println("leaving logger")
// }

// func middlewareOne(res http.ResponseWriter, req *http.Request, next coco.NextFunc) {
// 	log.Println("Executing middlewareOne")
// 	next.ServeHTTP(res, req)
// 	log.Println("Executing middlewareOne again")
// }

// func middlewareTwo(res http.ResponseWriter, req *http.Request, next coco.NextFunc) {

// 	log.Println("Executing middlewareTwo")
// 	if req.URL.Path == "/foo" {
// 		return
// 	}
// 	next.ServeHTTP(res, req)
// 	log.Println("Executing middlewareTwo again")

// }
