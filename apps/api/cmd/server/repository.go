package main

import "github.com/google/uuid"
type LeadRepository struct{
	leads []Lead
}

func (r *LeadRepository) Save(lead Lead) string {
	lead.ID = uuid.New().String()
	r.leads = append(r.leads, lead)
	return lead.ID
}

func (r *LeadRepository) List() []Lead {
	return r.leads
}