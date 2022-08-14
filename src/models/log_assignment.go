package models

type LogAssignment struct {
	Id           string `json:"id"`
	AssignmentId int    `json:"assignmentId"`
	Action       string `json:"action"`
	UserId       int    `json:"userId"`
}
