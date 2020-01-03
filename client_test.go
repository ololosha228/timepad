package timepad

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEvents(t *testing.T) {
	res, err := GetEvents(10)
	assert.Nil(t, err, "got an error. except no one: ", err)

	assert.Equal(t, len(res), 10, "expected 10 events, got ", len(res))

	randomEvent := res[rand.Intn(10)]

	assert.NotEmpty(t, randomEvent.ID, "event id is 0")
	assert.NotEmpty(t, randomEvent.StartsAt, "starts_at is 1 Jan 1970")
	assert.NotEmpty(t, randomEvent.Name, "name is ''")
	assert.NotEmpty(t, randomEvent.DescriptonShort, "description is ''")
	assert.NotEmpty(t, randomEvent.URL, "url is ''")
	assert.NotEmpty(t, randomEvent.Category, "category is ''")
}
