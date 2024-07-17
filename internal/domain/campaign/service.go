package campaign

import (
	contract "mailsender/internal/contract/requests"
	ErrorType "mailsender/internal/domain/errors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, []error) {

	campaign, err := New(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if hasError := s.Repository.Save(campaign); hasError != nil {
		err = append(err, ErrorType.ErrInternal)
	}

	if err != nil {
		return "", err
	}

	return campaign.ID, nil
}
