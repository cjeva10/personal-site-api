package noteRoutes

import (
    "github.com/gofiber/fiber/v2"
    noteHandler "git.chrisevanko.com/personal-site-api.git/internal/handlers/note"
)

func SetupNoteRoutes(router fiber.Router) {
    note := router.Group("/note")

    note.Post("/", noteHandler.CreateNotes)

    note.Get("/", noteHandler.GetNotes)

    note.Get("/:noteId", noteHandler.GetNote)

    note.Put("/:noteId", noteHandler.UpdateNote)

    note.Delete("/:noteId", noteHandler.DeleteNote)
}
