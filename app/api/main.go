package main

func main() {
	db := GetDBClient()
	defer db.Close()
	Init(db)
	// Create(db)
	UpdateRepositories()
}
