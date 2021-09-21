package countries

import (
	"github.com/annakallo/travelog/log"
	"github.com/annakallo/travelog/mysql"
	"github.com/annakallo/travelog/settings"
)

const (
	LogPrefix   = "Version countries"
	PackageName = "countries"
)

func UpdateCountriesTable() string {
	version := settings.GetCurrentVersion(PackageName)
	version = updateV1M0(version)
	return version
}

func updateV1M0(version string) string {
	db := mysql.GetInstance().GetConn()

	if version == "" {
		query := `create table if not exists countries (
					id int(11) unsigned not null auto_increment,
					country_name varchar(255)  not null,
					country_code varchar(255)  not null,
					visited boolean,
					on_wish_list boolean,
					created_at datetime not null default now(),
					updated_at datetime not null default now(),
					PRIMARY KEY (id)
					);`
		_, e := db.Exec(query)
		if e != nil {
			log.GetInstance().Errorf(LogPrefix, "Trouble at creating countries table: ", e)
			return version
		}
		log.GetInstance().Infof(LogPrefix, "Table countries created.")

		version = "v1.0-0"
		settings.UpdateVersion(PackageName, version)
	}
	if version == "v1.0-0" {
		defaultCountries := []Country{
			{CountryName: "Romania", CountryCode: "RO", Visited: true, OnWishList: false},
			{CountryName: "Spain", CountryCode: "ES", Visited: true, OnWishList: false},
			{CountryName: "Hungary", CountryCode: "HU", Visited: true, OnWishList: false},
			{CountryName: "Netherlands", CountryCode: "NL", Visited: true, OnWishList: false},
			{CountryName: "France", CountryCode: "FR", Visited: true, OnWishList: false},
			{CountryName: "United Kingdom", CountryCode: "GB", Visited: true, OnWishList: false},
			{CountryName: "Italy", CountryCode: "IT", Visited: true, OnWishList: false}}
		for _, country := range defaultCountries {
			stmt, _ := db.Prepare(`insert countries set country_name=?, country_code=?, visited=?, on_wish_list=?`)
			_, e := stmt.Exec(country.CountryName, country.CountryName, country.Visited, country.OnWishList)
			if e != nil {
				log.GetInstance().Errorf(LogPrefix, "Trouble when inserting default country in countries table: ", e.Error())
				return version
			}
			stmt.Close()
		}

		log.GetInstance().Infof(LogPrefix, "Table countries updated.")
		version = "v1.0-1"
		settings.UpdateVersion(PackageName, version)
	}

	return version
}
