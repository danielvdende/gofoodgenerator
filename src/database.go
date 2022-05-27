package webserver

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

var sqlDatabase = SetupDatabase()

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func SetupDatabase() *sql.DB {
	sqlitePath := "./sqlite-database.db"
	databaseExists := fileExists(sqlitePath)

	sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Fatal(err)
	}

	if !databaseExists {
		// If this is the first run, initialize the db with dishes from the csv file
		log.Printf("Database %s does not exist", sqlitePath)
		log.Println("Initializing sqlite-database.db...")

		createTable(sqliteDatabase) // Create Database Tables
		dishes := loadDishesFromFile()
		bulkInsert(dishes, sqliteDatabase)

	}
	return sqliteDatabase
}

func createTable(db *sql.DB) {
	createDishesTable := `CREATE TABLE IF NOT EXISTS dishes (
		"name" TEXT NOT NULL PRIMARY KEY,
		"category" TEXT,
		"preparationTime" TEXT
	 );`

	log.Println("Create dishes table...")
	statement, err := db.Prepare(createDishesTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("dishes table created")
}

func bulkInsert(unsavedRows []Dish, db *sql.DB) error {
	valueStrings := make([]string, 0, len(unsavedRows))
	fieldCount := reflect.TypeOf(Dish{}).NumField()
	valueArgs := make([]interface{}, 0, len(unsavedRows)*fieldCount)
	for _, post := range unsavedRows {
		// TODO: the number of params should depend on the number of fields in the struct
		valueStrings = append(valueStrings, "(?, ?, ?)")
		valueArgs = append(valueArgs, post.Name)
		valueArgs = append(valueArgs, post.Category)
		valueArgs = append(valueArgs, post.PreparationTime)
	}
	stmt := fmt.Sprintf("INSERT INTO dishes (name, category, preparationTime) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err := db.Exec(stmt, valueArgs...)
	return err
}

func loadDishesFromFile() []Dish {
	pathToDishes := "food.csv"
	var dishes []Dish
	f, err := os.Open(pathToDishes)
	if err != nil {
		log.Fatal("Unable to read input file "+pathToDishes, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	// Read the first line to skip the header
	csvReader.Read()
	for {
		record, error := csvReader.Read()
		if error == io.EOF {
			break
		}
		if error != nil {
			log.Fatal(error)
		}
		dishes = append(dishes, Dish{
			Name:            record[0],
			Category:        record[1],
			PreparationTime: record[2],
		})
	}
	return dishes
}

func GetDishWithConditions(conditions Rule) Dish {
	formattedPrepTime := "'" + strings.Join(conditions.PreparationTime, "','") + "'"
	query := fmt.Sprintf("SELECT * FROM dishes WHERE preparationTime in (%s) ORDER BY RANDOM() limit 1", formattedPrepTime)
	var name string
	var category string
	var preparationTime string

	err := sqlDatabase.QueryRow(query).Scan(&name, &category, &preparationTime)
	if err != nil {
		log.Fatal(err)
	}

	return Dish{Name: name, Category: category, PreparationTime: preparationTime}
}
