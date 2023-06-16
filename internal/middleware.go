package middleware

import (
	"time"

	"git.chrisevanko.com/personal-site-api.git/database"
	"git.chrisevanko.com/personal-site-api.git/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/google/uuid"
)

// check whether a given api key is valid
func ValidateApiKey(c *fiber.Ctx, key string) (bool, error) {
    db := database.DB
    var apiKey model.ApiKey

    // search for this key in the database
    db.Find(&apiKey, "key = ?", key)

    // make sure that the key is not empty
    if apiKey.Key == uuid.Nil {
        return false, keyauth.ErrMissingOrMalformedAPIKey
    }

    // check that the key is not expired
    duration := apiKey.Duration
    updatedAt := apiKey.UpdatedAt

    if updatedAt.Add(duration).After(time.Now()) {
        return false, keyauth.ErrMissingOrMalformedAPIKey 
    }

    return true, nil 
}