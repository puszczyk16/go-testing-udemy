package locations_provider

import (
	//"github.com/puszczyk16/go-testing-udemy/src/api/domain/locations"
	//"github.com/puszczyk16/go-testing-udemy/src/api/utils/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"net/http"
	//"fmt"
)

func TestGetCountryRestclientError(t *testing.T) {
	country, err := GetCountry("AR")
	assert.Nil(t,country)
	//fmt.Println(country)
	assert.NotNil(t,err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient error when getting country AR", err.Message)
}
	
func TestGetCountryNotFound(t *testing.T) {
	country, err := GetCountry("AR")
	assert.Nil(t,country)
	assert.NotNil(t,err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T) {
	country, err := GetCountry("AR")
	assert.Nil(t,country)
	assert.NotNil(t,err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestGetCountryInvalidJsonResponse(t *testing.T) {
	country, err := GetCountry("AR")
	assert.Nil(t,country)
	assert.NotNil(t,err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to un,arshal country data from AR", err.Message)
}


func TestGetCountryNoError(t *testing.T) {
	country, err := GetCountry("AR")
	assert.Nil(t,err)
	assert.NotNil(t,country)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
	assert.EqualValues(t, "GMT-03:00", country.TimeZone)
	assert.EqualValues(t, 24, len(country.States))

}

