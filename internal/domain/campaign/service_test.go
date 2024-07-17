package campaign

import (
	"errors"
	contract "mailsender/internal/contract/requests"
	ErrorType "mailsender/internal/domain/errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	campaign = contract.NewCampaign{
		Name:    "teste",
		Content: "testando 123",
		Emails:  []string{"sefwkefjskfjksefj", "yktukyukk"},
	}
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_Campaign_Validate_Domain(t *testing.T) {

	//Arrange
	assert := assert.New(t)
	var errors = make([]string, 0)
	campaign := contract.NewCampaign{}
	mockRepository := new(repositoryMock)
	service := Service{mockRepository}

	//Act
	_, err := service.Create(campaign)

	for _, e := range err {
		if e != nil {
			errors = append(errors, e.Error())
		}
	}

	//Assert
	assert.Contains(errors, "name is required")
	assert.Contains(errors, "must have an content")
	assert.Contains(errors, "destinary mail is necessary")
	assert.Equal(len(errors), 3)
}

func Test_Create_Campaign_Validate_Domain_Name(t *testing.T) {

	//Arrange
	assert := assert.New(t)
	var error = ""
	mockRepository := new(repositoryMock)
	service := Service{mockRepository}

	//Act
	_, err := service.Create(campaign)

	for _, e := range err {
		if e != nil {
			error = e.Error()
		}
	}

	//Assert
	assert.Contains(error, "must declare an name")
}

func Test_Create_Save_Campaign(t *testing.T) {

	//Arrange
	mockRepository := new(repositoryMock)
	mockRepository.On("Save", mock.MatchedBy(func(c *Campaign) bool {

		if campaign.Name != c.Name ||
			campaign.Content != c.Content ||
			len(campaign.Emails) != len(c.Contacts) {
			return false
		}

		return true
	})).Return(nil)

	service := Service{mockRepository}

	//Act
	service.Create(campaign)

	//Assert
	mockRepository.AssertExpectations(t)
}

func Test_Create_Save_Campaign_Error(t *testing.T) {

	//Arrange
	assert := assert.New(t)
	mockRepository := new(repositoryMock)
	mockRepository.On("Save", mock.Anything).Return(ErrorType.ErrInternal)
	service := Service{mockRepository}

	//Act
	_, err := service.Create(campaign)

	//Assert
	assert.True(errors.Is( ErrorType.ErrInternal, err[0]))
}
