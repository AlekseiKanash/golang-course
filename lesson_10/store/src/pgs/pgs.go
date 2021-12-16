package pgs

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/AlekseiKanash/golang-course/lesson_10/openweather"
	_ "github.com/lib/pq"
)

type CityWeatherInfo struct {
	City        string
	Temperature float32
}

func NewPostgreDS(auth_data DbAuthData) *PosgtreDS {
	var ds = &PosgtreDS{
		Auth_data: auth_data,
		Driver:    "postgres",
		Db:        nil,
	}
	return ds
}

type DbAuthData struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

type IDB interface {
	Connect() (*sql.DB, error)
	Close()
	Create(city string) (int, error)
	Update(int, openweather.Response) error
}

type PosgtreDS struct {
	Auth_data DbAuthData
	Driver    string
	Db        *sql.DB
}

func (p *PosgtreDS) Connect() error {
	if p.Db != nil {
		p.Db.Close()
		p.Db = nil
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		p.Auth_data.Host, p.Auth_data.Port, p.Auth_data.User, p.Auth_data.Password, p.Auth_data.Dbname)

	db, err := sql.Open(p.Driver, psqlInfo)
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	err = db.Ping()
	if err != nil {
		defer db.Close()
		fmt.Printf("%v\n", err)
		return err
	}

	p.Db = db
	return nil
}

func (p *PosgtreDS) Close() {
	if p.Db != nil {
		p.Db.Close()
	}
}

func (p *PosgtreDS) Create(city string) (int, error) {
	if p.Db == nil {
		err := p.Connect()
		if err != nil {
			return 0, err
		}
	}
	sqlStatement := `
	INSERT INTO weather (city, temperature)
	VALUES ($1, $2)
	RETURNING id`
	id := 0
	err := p.Db.QueryRow(sqlStatement, strings.ToLower(city), .0).Scan(&id)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	return id, err
}

func (p *PosgtreDS) Update(id int, wdata openweather.Response) error {
	if p.Db == nil {
		err := p.Connect()
		if err != nil {
			return err
		}
	}

	sqlStatement := `
	UPDATE weather
	SET city=$1, temperature=$2
	WHERE id=$3`
	_, err := p.Db.Exec(sqlStatement, strings.ToLower(wdata.Name), wdata.Main.Temp, id)

	return err
}

func (p *PosgtreDS) Get(city string) float32 {
	if p.Db == nil {
		err := p.Connect()
		if err != nil {
			return .0
		}
	}

	sqlStatement := `
	SELECT id, temperature
	FROM weather
	WHERE city=$1
	ORDER by id DESC
	LIMIT 1`
	row := p.Db.QueryRow(sqlStatement, strings.ToLower(city))
	id := 0
	city = ""
	var temperature float32 = .0
	err := row.Scan(&id, &temperature)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	return temperature
}

func getAuthData() DbAuthData {
	auth_data := DbAuthData{
		Host:     "posgtres",
		Port:     5432,
		User:     "adm",
		Password: "adm",
		Dbname:   "weather",
	}
	return auth_data
}

func SaveWeather(city string) {
	auth_data := getAuthData()
	db := NewPostgreDS(auth_data)
	db.Connect()
	id, _ := db.Create(city)
	go func(db *PosgtreDS) {
		weather, err := openweather.GetWeather(city)
		if err == nil {
			fmt.Printf("%s\n", fmt.Sprintf("%f", weather.Main.Temp))
			db.Update(id, weather)
		}
		db.Close()
	}(db)
}

func GetWeather(city string) CityWeatherInfo {
	ret_info := CityWeatherInfo{City: city, Temperature: .0}

	auth_data := getAuthData()
	db := NewPostgreDS(auth_data)
	err := db.Connect()
	if err != nil {
		return ret_info
	}
	temperature := db.Get(city)
	ret_info.Temperature = temperature
	return ret_info
}

func ListWeather() []CityWeatherInfo {
	var ret_info []CityWeatherInfo

	auth_data := getAuthData()
	pdb := NewPostgreDS(auth_data)
	err := pdb.Connect()
	if err != nil {
		return ret_info
	}
	defer pdb.Close()

	sqlStatement := `
	SELECT id, city, temperature
	FROM weather
	ORDER by id DESC`
	rows, err := pdb.Db.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		fmt.Printf("%v\n", err)
		return ret_info
	}

	for rows.Next() {
		id := 0
		city := ""
		var temperature float32 = .0
		err := rows.Scan(&id, &city, &temperature)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		row_info := CityWeatherInfo{
			City:        city,
			Temperature: temperature,
		}
		ret_info = append(ret_info, row_info)
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("%v\n", err)
	}

	return ret_info
}
