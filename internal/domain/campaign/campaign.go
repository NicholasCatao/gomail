package campaign

import (
	"mailsender/internal/domain/contact"
	validate "mailsender/internal/domain/validations" 
	"time"

	guid "github.com/beevik/guid"
)

type Campaign struct {
	ID        string            `validate:"required"`
	Name      string            `validate:"min=5,max=25"`
	CreatedOn time.Time         `validate:"required"`
	Content   string            `validate:"min=5,max=1024"`
	Contacts  []contact.Contact `validate:"min=1",dive`
}

func New(name string, content string, mails []string) (r *Campaign, error []error) {

	contacts := make([]contact.Contact, len(mails))
	for index, value := range mails {
		contacts[index].Mail = value
	}

	campaign := &Campaign{
		ID:        guid.NewString(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}

	err := validate.ValidateStruct(campaign)

	if err != nil {
		return nil, err
	}

	return campaign, nil
}

// func validate(name string, content string, mailAdress []string) []error {

// 	errs := make([]error, 0)

// 	if name == "" {
// 		errs = append(errs, errors.New("must declare an name"))
// 	}

// 	if content == "" {
// 		errs = append(errs, errors.New("must have an content"))
// 	}

// 	if len(mailAdress) <= 0 {
// 		errs = append(errs, errors.New("destinary mail is necessary"))
// 	}

// 	if len(errs) > 1 {
// 		return errs
// 	}

// 	return nil
// }

func (c *Campaign) WithMailAddress(mails ...string) {

	addressContacts := make([]contact.Contact, len(mails))

	for index, value := range mails {
		addressContacts[index].Mail = value
	}

	c.Contacts = addressContacts
}
