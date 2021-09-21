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

func TestGetCategories(t *testing.T) {
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
	country = Country{
		CountryName: "Italy",
		CountryCode: "IT",
		Visited:     true,
		OnWishList:  false,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	e = country.Insert()
	assert.Nil(t, e)

	res := GetCountries()
	//assert.Equal(t, len(res), 2)
	assert.Equal(t, res[0].CountryName, "Hungary")
	assert.Equal(t, res[0].CountryCode, "HU")
	assert.Equal(t, res[0].Visited, true)
	assert.Equal(t, res[0].OnWishList, false)
	assert.Equal(t, res[1].CountryName, "Italy")
	for _, result := range res {
		country.Load(result.Id)
		country.Delete()
	}
}
