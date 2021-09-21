package countries

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCountrySaveAndFetch(t *testing.T) {
	country := Country{
		CountryName: "Hungary",
		CountryCode: "HU",
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	e := country.Insert()
	assert.Nil(t, e)
	categories := GetCountries()
	assert.NotEqual(t, len(categories), 0)
	e = country.Delete()
	assert.Nil(t, e)
}

func TestCategorySave(t *testing.T) {
	country := Country{
		CountryName: "Hungary",
		CountryCode: "HU",
		Visited:     true,
		OnWishList:  false,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	e := country.Insert()
	assert.Nil(t, e)
	country.CountryCode = "ES"
	country.Save()
	assert.Equal(t, country.CountryCode, "ES")
	e = country.Delete()
	assert.Nil(t, e)
}
