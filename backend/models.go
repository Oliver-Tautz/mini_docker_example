package main

type User struct {
	FirstName string
	Lastname  string
	Email     string
	Id        string
}

type Submission struct {
	Motivation       string
	CustomMotivation string
	Id               string
	User             User
}

type submitResponse struct {
	Status  string
	Message string
}
