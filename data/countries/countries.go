package countries

import (
	"fmt"
	"github.com/annakallo/travelog/mysql"
	"time"
)

type Country struct {
	Id          int       `json:"id"`
	CountryName string    `json:"country_name"`
	CountryCode string    `json:"country_code"`
	Visited     bool      `json:"visited"`
	OnWishList  bool      `json:"on_wish_list"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Countries []Country

// Load category
func (country *Country) Load(id int) error {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`select * from countries where id = ?`)
	defer stmt.Close()
	rows, e := stmt.Query(id)
	if e != nil {
		fmt.Printf("Error when preparing stmt id %d: %s", id, e.Error())
		return e
	}
	defer rows.Close()
	if rows.Next() {
		var createdAt string
		var updatedAt string
		e := rows.Scan(&country.Id, &country.CountryName, &country.CountryCode, &country.Visited, &country.OnWishList, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading id %v: %s", id, e.Error())
			return e
		}
		country.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		country.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
	}
	return nil
}

// Insert a new trade
func (country *Country) Insert() error {
	if country.CreatedAt.IsZero() {
		country.CreatedAt = time.Now().UTC()
	}
	if country.UpdatedAt.IsZero() {
		country.UpdatedAt = time.Now().UTC()
	}
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`insert countries set id=?, country_name=?, country_code=?, visited=?, on_wish_list=?, created_at=?, updated_at=?`)
	defer stmt.Close()

	res, e := stmt.Exec(country.Id, country.CountryName, &country.CountryCode, &country.Visited, &country.OnWishList, country.CreatedAt, country.UpdatedAt)
	if e != nil {
		fmt.Printf("Error when inserting country: %s", e.Error())
		return e
	}
	id, _ := res.LastInsertId()
	country.Id = int(id)
	return nil
}

func (country *Country) Save() error {
	if country.UpdatedAt.IsZero() {
		country.UpdatedAt = time.Now().UTC()
	}
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`update countries set country_name=?, country_code=?, visited=?, on_wish_list=?, created_at=?, updated_at=? where id=?`)
	defer stmt.Close()

	_, e := stmt.Exec(country.CountryName, &country.CountryCode, &country.Visited, &country.OnWishList, country.CreatedAt, country.UpdatedAt, country.Id)
	if e != nil {
		fmt.Printf("Error when saving country: %s", e.Error())
		return e
	}
	return nil
}

func (country *Country) Delete() error {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`delete from countries where id=?`)
	defer stmt.Close()
	_, e := stmt.Exec(country.Id)
	if e != nil {
		fmt.Printf("Error when deleting country: %s", e.Error())
		return e
	}
	return e
}

func GetCountries() Countries {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`select * from countries order by country_name ASC`)
	defer stmt.Close()
	rows, e := stmt.Query()
	if e != nil {
		fmt.Printf("Error when preparing stmt in getting all countries: %s", e.Error())
		return Countries{}
	}
	defer rows.Close()
	countries := Countries{}
	for rows.Next() {
		country := Country{}
		var createdAt string
		var updatedAt string
		e := rows.Scan(&country.Id, &country.CountryName, &country.CountryCode, &country.Visited, &country.OnWishList, &createdAt, &updatedAt)
		if e != nil {
			fmt.Printf("Error when loading countries: %s", e.Error())
			return Countries{}
		}
		country.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		country.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
		countries = append(countries, country)
	}
	return countries
}
