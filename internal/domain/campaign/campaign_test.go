package campaign

import (
	"testing"

	"github.com/beevik/guid"
	"github.com/stretchr/testify/assert"
)

var (
	name = "coolname"
	content = "cinetete"
	contacts = []string{"email@hotmail.com", "email@hotmail.com"}
	//now = time.Now().Add(-time.Minute)
)

func TestNew(t *testing.T) {
	//Arrange
	assert := assert.New(t)

	//Act
	campaign, _ := New(name, content, contacts)

	//Assert
	assert.IsTypef(guid.NewString(), campaign.ID, "Id Must be an GUID")
	assert.NotEmpty(campaign.Name)
}

func TestNew_IdNotNil(t *testing.T) {
	//Arrange
	assert := assert.New(t)

	//Act
	campaign, _ := New(name, content, contacts)

	//Assert
	assert.NotNil(campaign.ID)

}

func TestNew_MustHaveContacts(t *testing.T){
	// Arrange
	assert := assert.New(t)
	 noContacts := make([]string,0)
	//Act
	campaign, _ := New(name, content, noContacts)
	//campaign.WithMailAddress(contacts...)

	assert.Equal(len(campaign.Contacts), len(contacts))
}

