package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/fathurwithyou/Day3-Training-ARC-13523105/backend/internal/models"
)

func GetStudentCourses(c *fiber.Ctx) error {
	usersFile := filepath.Join("data", "users.json")
	coursesFile := filepath.Join("data", "courses.json")
	examScoresFile := filepath.Join("data", "examscores.json")

	usersData, err := os.ReadFile(usersFile)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var users []models.User
	if err := json.Unmarshal(usersData, &users); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	coursesData, err := os.ReadFile(coursesFile)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var courses []models.Course
	if err := json.Unmarshal(coursesData, &courses); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	examScoresData, err := os.ReadFile(examScoresFile)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var examScores []models.ExamScore
	if err := json.Unmarshal(examScoresData, &examScores); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	courseMap := make(map[int]models.Course)
	for _, course := range courses {
		courseMap[course.ID] = course
	}

	userCourseIDs := make(map[int]map[int]bool)
	for _, es := range examScores {
		if userCourseIDs[es.UserID] == nil {
			userCourseIDs[es.UserID] = make(map[int]bool)
		}
		userCourseIDs[es.UserID][es.CourseID] = true
	}

	type StudentCourses struct {
		User    models.User     `json:"user"`
		Courses []models.Course `json:"courses"`
	}

	var result []StudentCourses
	for _, user := range users {
		var userCourses []models.Course
		if courseIDs, exists := userCourseIDs[user.ID]; exists {
			for courseID := range courseIDs {
				if course, ok := courseMap[courseID]; ok {
					userCourses = append(userCourses, course)
				}
			}
		}
		result = append(result, StudentCourses{
			User:    user,
			Courses: userCourses,
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}
