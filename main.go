package main

import "os"

func main() {

	a := App{} 

	// before calling os.Getenv we have to set these environment variables 
	// for our connection to database 
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8000")
}	


