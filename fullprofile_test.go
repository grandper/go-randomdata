package randomdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FullProfileGenerator(t *testing.T) {
	r := FromSeed(1234)

	profile := r.GenerateProfile(1)
	assert.Equal(t, "female", profile.Gender)

	profile = r.GenerateProfile(0)
	assert.Equal(t, "male", profile.Gender)

	profile = r.GenerateProfile(2)
	assert.NotNil(t, profile, "profile failed to generate")

	CheckPhoneNumber(profile.Cell, t, "expected Cell# to be a valid phone number: %v", profile.Cell)
	CheckPhoneNumber(profile.Phone, t, "expected Phone# to be a valid phone number: %v", profile.Phone)

	assert.NotEmpty(t, profile.Login.Username, "profile Username failed to generate")
	assert.NotEmpty(t, profile.Location.Street, "profile Street failed to generate")
	assert.Equal(t, "SSN", profile.ID.Name, "profile ID Name to be SSN, but got %s\n", profile.ID.Name)
	assert.NotEmpty(t, profile.Picture.Large, "profile Picture Large failed to generate", profile.Picture.Large)
}
