package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/fathurwithyou/Day3-Training-ARC-13523105/backend/internal/models"
)

var usersFile = filepath.Join("data", "users.json")

func ensureDataFile() error {
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		if err := os.Mkdir("data", 0755); err != nil {
			return err
		}
	}
	if _, err := os.Stat(usersFile); os.IsNotExist(err) {
		return os.WriteFile(usersFile, []byte("[]"), 0644)
	}
	return nil
}

func GetUsers(c *fiber.Ctx) error {
	if err := ensureDataFile(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	data, err := os.ReadFile(usersFile)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var users []models.User
	if err := json.Unmarshal(data, &users); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	if err := ensureDataFile(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var newUser models.User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	data, err := os.ReadFile(usersFile)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var users []models.User
	if err := json.Unmarshal(data, &users); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	maxID := 0
	for _, u := range users {
		if u.ID > maxID {
			maxID = u.ID
		}
	}
	newUser.ID = maxID + 1
	users = append(users, newUser)
	updatedData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err := os.WriteFile(usersFile, updatedData, 0644); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(newUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	if err := ensureDataFile(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	data, err := os.ReadFile(usersFile)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var users []models.User
	if err := json.Unmarshal(data, &users); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var updatedUser models.User
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	found := false
	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id
			users[i] = updatedUser
			found = true
			break
		}
	}
	if !found {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	updatedData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err := os.WriteFile(usersFile, updatedData, 0644); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(updatedUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	if err := ensureDataFile(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	data, err := os.ReadFile(usersFile)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var users []models.User
	if err := json.Unmarshal(data, &users); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	found := false
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	updatedData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err := os.WriteFile(usersFile, updatedData, 0644); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "User deleted"})
}