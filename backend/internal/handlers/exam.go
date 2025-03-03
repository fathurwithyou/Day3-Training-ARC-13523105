package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/fathurwithyou/Day3-Training-ARC-13523105/backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

var examScoresFile = filepath.Join("data", "examscores.json")

func ensureExamScoresFile() error {
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		if err := os.Mkdir("data", 0755); err != nil {
			return err
		}
	}
	if _, err := os.Stat(examScoresFile); os.IsNotExist(err) {
		return os.WriteFile(examScoresFile, []byte("[]"), 0644)
	}
	return nil
}

func GetExamScoresByUserID(c *fiber.Ctx) error {
	if err := ensureExamScoresFile(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	data, err := os.ReadFile(examScoresFile)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var scores []models.ExamScore
	if err := json.Unmarshal(data, &scores); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	userScores := make([]models.ExamScore, 0)
	for _, s := range scores {
		if s.UserID == userID {
			userScores = append(userScores, s)
		}
	}

	coursesFile := filepath.Join("data", "courses.json")
	coursesData, err := os.ReadFile(coursesFile)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var courses []models.Course
	if err := json.Unmarshal(coursesData, &courses); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	courseMap := make(map[int]string)
	for _, course := range courses {
		courseMap[course.ID] = course.CourseCode
	}

	type ExamScoreResponse struct {
		ID         int    `json:"id"`
		UserID     int    `json:"user_id"`
		CourseCode string `json:"course_code"`
		Score      int    `json:"score"`
	}

	var resp []ExamScoreResponse
	for _, s := range userScores {
		resp = append(resp, ExamScoreResponse{
			ID:         s.ID,
			UserID:     s.UserID,
			CourseCode: courseMap[s.CourseID],
			Score:      s.Score,
		})
	}
	return c.Status(http.StatusOK).JSON(resp)
}

func CreateExamScore(c *fiber.Ctx) error {
	if err := ensureExamScoresFile(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var newScore models.ExamScore
	if err := c.BodyParser(&newScore); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	data, err := os.ReadFile(examScoresFile)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var scores []models.ExamScore
	if err := json.Unmarshal(data, &scores); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	maxID := 0
	for _, s := range scores {
		if s.ID > maxID {
			maxID = s.ID
		}
	}
	newScore.ID = maxID + 1
	scores = append(scores, newScore)
	updatedData, err := json.MarshalIndent(scores, "", "  ")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err := os.WriteFile(examScoresFile, updatedData, 0644); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(newScore)
}

func UpdateExamScore(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid exam score ID"})
	}
	if err := ensureExamScoresFile(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	data, err := os.ReadFile(examScoresFile)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var scores []models.ExamScore
	if err := json.Unmarshal(data, &scores); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var updatedScore models.ExamScore
	if err := c.BodyParser(&updatedScore); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	found := false
	for i, s := range scores {
		if s.ID == id {
			updatedScore.ID = id
			scores[i] = updatedScore
			found = true
			break
		}
	}
	if !found {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Exam score not found"})
	}
	updatedData, err := json.MarshalIndent(scores, "", "  ")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err := os.WriteFile(examScoresFile, updatedData, 0644); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(updatedScore)
}

func DeleteExamScore(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid exam score ID"})
	}
	if err := ensureExamScoresFile(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	data, err := os.ReadFile(examScoresFile)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var scores []models.ExamScore
	if err := json.Unmarshal(data, &scores); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	found := false
	for i, s := range scores {
		if s.ID == id {
			scores = append(scores[:i], scores[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Exam score not found"})
	}
	updatedData, err := json.MarshalIndent(scores, "", "  ")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err := os.WriteFile(examScoresFile, updatedData, 0644); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Exam score deleted"})
}
