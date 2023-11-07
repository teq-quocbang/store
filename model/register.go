package model

import "time"

type Register struct {
	AccountID  uint      `json:"account_id"`
	SemesterID string    `json:"semester_id"`
	ClassID    string    `json:"class_id"`
	CourseID   string    `json:"course_id"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  uint      `json:"created_by"`
	IsCanceled bool      `json:"is_canceled"`
}

func (Register) TableName() string {
	return "register"
}
