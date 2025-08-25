package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Note struct{
	ID int `json:"id"`
	Title string `json:"title"`

}

var notes = make(map[int]Note)
var nextID = 1

//response for home url route
func handlerHome(c *fiber.Ctx)error {
	return c.SendString("Hello, Fiber!")
}

//getting notes with id route
func handlerGetNotes(c *fiber.Ctx)error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.SendStatus(400)
	}
	note := notes[id]
	return c.JSON(note)

}

//response from posting notes or making a note
func handlerCreateNotes(c *fiber.Ctx)error {
	note := new (Note)
	if err := c.BodyParser(note); err !=nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	note.ID = nextID
	notes[note.ID] = *note
	nextID++

	return c.Status(fiber.StatusCreated).JSON(note)
}

func main(){
	//started using fiber backend framework
	app := fiber.New() // -> fiber declaration make object called "app"
	app.Get("/",handlerHome)
	app.Get("/notes/:id",handlerGetNotes)
	app.Post("/notes",handlerCreateNotes)

	//->error handling
    err := app.Listen(":8181")
	if err != nil {
        fmt.Println(err.Error())
    }

}