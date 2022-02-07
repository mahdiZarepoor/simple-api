// main.go

package main

import "os"

func main() {
	// our app is nothing more than a structure which store 
	// two values . one for router second for database .
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8010")
}	