package countries

import (
	"github.com/annakallo/travel-log-server/mysql"
	"github.com/annakallo/travel-log-server/settings"
	"github.com/annakallo/travelog/log"
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

	return version
}
