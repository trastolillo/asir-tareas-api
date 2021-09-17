package main

import (
	"fmt"
	"os"
)

// var _ = godotenv.Load(".env")

var (
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("userBD"),
		os.Getenv("passBD"),
		os.Getenv("host"),
		os.Getenv("PORT"),
		os.Getenv("nameBD"),
	)
)

// const AllowedCORSDomain = "http://localhost"
