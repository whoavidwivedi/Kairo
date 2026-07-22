package main

type CreateLeadRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type CreateLeadResponse struct {
	Message string `json:"message"`
	LeadID  string `json:"leadId"`
}

type Lead struct {
	ID string
	Name string
	Email string
	Phone string
}