package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/noelukwa/coco"
)

func main() {
	// create a new app and set the global prefix
	app := coco.NewApp().GlobalPrefix("api/v1/")

	app.Get("/hey", reqLogger, func(rw coco.Response, r *coco.Request, next coco.NextFunc) {
		rw.JSON(200, map[string]string{
			"message": "hello world",
		})
	})

	// named param
	app.Get("/users/:name", func(rw coco.Response, r *coco.Request, next coco.NextFunc) {
		// returns a map of params
		name := r.Params()["name"]

		rw.JSON(200, map[string]interface{}{
			"message": fmt.Sprintf("hello %s", name),
		})
	})

	// subroute with prefix
	fruits := app.NewRoute("/fruits")

	store := []string{"apple", "orange", "banana"}

	fruits.Get("/", func(rw coco.Response, r *coco.Request, next coco.NextFunc) {
		find := r.Query().Get("find")
		if find == "" {
			rw.JSON(200, store)
			return
		}

		var found []string
		for _, f := range store {
			if strings.Contains(f, find) {
				found = append(found, f)
			}
		}
		rw.JSON(200, map[string]interface{}{
			"found": found,
			"total": len(found),
		})
	})

	if err := app.Listen(":8040", context.Background()); err != nil {
		log.Fatal(err)
	}
}

// middleware
func reqLogger(rw coco.Response, r *coco.Request, next coco.NextFunc) {
	fmt.Println("request received")
	next(rw, r)
	fmt.Println("request completed")
}
