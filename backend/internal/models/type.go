package models

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	NIM   string `json:"nim"`
	Email string `json:"email"`
}

type Course struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	CourseCode string `json:"course_code"`
}

type ExamScore struct {
	ID    int `json:"id"`
	UserID int `json:"user_id"`
	CourseID int `json:"course_id"`
	Score int `json:"score"`
}

