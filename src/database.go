package webserver

//func InitDb() *gorm.DB {
//	db, err := gorm.Open("sqlite3", "./data.db")
//	db.LogMode(true)    // Error
//	if err != nil {
//		panic(err)
//	}    // Creating the table
//	if !db.HasTable(&Dish{}) {
//		db.CreateTable(&Dish{})
//		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Dish{})
//	}
//
//	return db
//}

//func InitDatabase(dishes []Dish) {
//	log.Println("Creating sqlite-database.db...")
//	file, err := os.Create("sqlite-database.db") // Create SQLite file
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//	file.Close()
//	log.Println("sqlite-database.db created")
//
//	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")
//	defer sqliteDatabase.Close()        // Defer Closing the database
//	createTable(sqliteDatabase)         // Create Database Tables
//
//	bulkInsert(dishes, sqliteDatabase)
//
//	displayDishes(sqliteDatabase)
//}
//
//func createTable(db *sql.DB) {
//	createDishesTable := `CREATE TABLE dishes (
//		"name" TEXT NOT NULL PRIMARY KEY,
//		"category" TEXT,
//		"preparationTime" TEXT,
//	  );`
//
//	log.Println("Create dishes table...")
//	statement, err := db.Prepare(createDishesTable)
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//	log.Println("FOOO")
//	statement.Exec()
//	log.Println("dishes table created")
//}
//
//func bulkInsert(unsavedRows []Dish, db *sql.DB) error {
//	valueStrings := make([]string, 0, len(unsavedRows))
//	fieldCount := reflect.TypeOf(Dish{}).NumField()
//	valueArgs := make([]interface{}, 0, len(unsavedRows) * fieldCount)
//	for _, post := range unsavedRows {
//		// TODO: the number of params should depend on the number of fields in the struct
//		valueStrings = append(valueStrings, "(?, ?, ?)")
//		valueArgs = append(valueArgs, post.Name)
//		valueArgs = append(valueArgs, post.Category)
//		valueArgs = append(valueArgs, post.PreparationTime)
//	}
//	stmt := fmt.Sprintf("INSERT INTO dishes (name, category, preparationTime) VALUES %s",
//		strings.Join(valueStrings, ","))
//	_, err := db.Exec(stmt, valueArgs...)
//	return err
//}
//
//func displayDishes(db *sql.DB) {
//	row, err := db.Query("SELECT * FROM dishes ORDER BY name")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer row.Close()
//	for row.Next() { // Iterate and fetch the records from result cursor
//		var name string
//		var category string
//		var preparationTime string
//		row.Scan(&name, &category, &preparationTime)
//		log.Println("Dish: ", name, " ", category, " ", preparationTime)
//	}
//}
//
