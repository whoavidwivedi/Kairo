package main


type LeadService struct{
	repository *LeadRepository
}

func NewLeadService(repository *LeadRepository) *LeadService {
	return &LeadService{
		repository: repository,
	}
}

func (s *LeadService) CreateLead(req CreateLeadRequest) CreateLeadResponse {
	lead := Lead{
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	leadID := s.repository.Save(lead)

	response := CreateLeadResponse{
		Message: "Lead created successfully",
		LeadID:  leadID,
	}

	return response

}

func (s *LeadService) ListLeads() []Lead {
	return s.repository.List()
}