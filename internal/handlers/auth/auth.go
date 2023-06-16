package authHandler

import (
    "time"
	"crypto/sha256"
    "encoding/base64"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"git.chrisevanko.com/personal-site-api.git/database"
	"git.chrisevanko.com/personal-site-api.git/internal/model"
)
func hashPassword(password string) string {
    rawHash := sha256.Sum256([]byte(password))
    hash := base64.URLEncoding.EncodeToString(rawHash[:])
    return hash 
}

// given user id, get the active api key or else generate a new one
// note the user id must be valid since Register and Login handle validating the user
func generateApiKey(userId uint) uuid.UUID {
	// if valid, check the current api key for validity
    var apiKey model.ApiKey
    db := database.DB

    db.Where("user_id = ?", userId).Find(&apiKey)

    // no api key issued yet
    if apiKey.ID == 0 {
        apiKey.UserId = userId
        apiKey.Key = uuid.New()
        apiKey.Duration = time.Duration(-1)
    }

    db.Save(&apiKey)

    // get updated at time and duration
    updatedAt := apiKey.UpdatedAt
    duration := apiKey.Duration

    // if api key duration < 0, then we have an unlimimted duration
    if updatedAt.Add(duration).Before(updatedAt) {
        return apiKey.Key
    }

	// if api key is expired, generate and return a new api key
    now := time.Now()
    if updatedAt.Add(duration).After(now) {
        // generate a new api key
        apiKey.Key = uuid.New()
        db.Where("user_id = ?", userId).Save(&apiKey)
    }

    return apiKey.Key
}

func Register(c *fiber.Ctx) error {
    user := new(model.User)
	db := database.DB

    // parse the input into a user object
    err := c.BodyParser(&user)
    if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
    }

    // validate inputs
    if user.Username == "" {
        return c.Status(403).JSON(fiber.Map{"status": "error", "message": "Missing username", "data": nil})
    }
    if user.Password == "" {
        return c.Status(403).JSON(fiber.Map{"status": "error", "message": "Missing password", "data": nil})
    }

	// check db for a duplicate username
    var userCheck model.User
    db.Find(&userCheck, "username = ?", user.Username)
    if userCheck.ID != 0 {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "user already exists", "data": nil})
    }

    // make sure password is sufficiently long
    if len(user.Password) < 8 {
        return c.Status(403).JSON(fiber.Map{"status": "error", "message": "Password must be at least 8 characters", "data": nil})
    }

	// hash password
    user.Password = hashPassword(user.Password)

	// insert user and hashed password to db
    err = db.Create(&user).Error
    if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "could not create user", "data": err})
    }

    // login the new user
    return Login(c)
}

func Login(c *fiber.Ctx) error {
	db := database.DB
	var user model.User

    // parse the input into a user object
    err := c.BodyParser(&user)
    if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
    }

	// hash password
    user.Password = hashPassword(user.Password)

	// check user and hashed password with table in db
    db.Where("username = ? AND password = ?", user.Username, user.Password).Find(&user)
    if user.ID == 0 {
        return c.Status(401).JSON(fiber.Map{"status": "error", "message": "incorrect username or password", "data": nil})
    }

    apiKey := generateApiKey(user.ID)

	return c.JSON(fiber.Map{
        "status": "success", 
        "message": "user logged in", 
        "data": fiber.Map{
            "key": apiKey,
        },
    })
}

