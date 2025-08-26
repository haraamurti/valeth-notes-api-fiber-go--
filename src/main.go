package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Note struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

var notes = make(map[int]Note)
var nextID = 1

//response for home url route
func handlerHome(c *fiber.Ctx)error {
	return c.JSON(notes) //-> bakal ngembaliin json yaitu hashmap dari notes ini
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

//handler updating and editing a note
func handlerUpdateNote(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(400)
	}
	note := new(Note)
	c.BodyParser(note)
	note.ID = id
	notes[id] = *note
	return c.JSON(note)
}

//handler to delet a note
func handlertDeleteNote(c *fiber.Ctx)error{
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.SendStatus(400)
	}
	delete(notes,id)
	return c.JSON("file with the id "+ strconv.Itoa(id)+"is already deleted")
}

func main(){
	//started using fiber backend framework
	app := fiber.New() // -> fiber declaration make object called "app"
	app.Get("/",handlerHome) //home
	app.Get("/notes/:id",handlerGetNotes) //getnotes
	app.Post("/notes",handlerCreateNotes) //create notes
	app.Put("/notes/:id", handlerUpdateNote) // update notes and content
	app.Delete("/delete/notes/:id", handlertDeleteNote)//delete a note

	//->error handling
    err := app.Listen(":8181")
	if err != nil {
        fmt.Println(err.Error())
    }

}