package main

import (
	"fmt"
	"os"
)

// var _ = godotenv.Load(".env")

var (
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		os.Getenv("userBD"),
		os.Getenv("passBD"),
		os.Getenv("host"),
		// os.Getenv("PORT"),
		os.Getenv("nameBD"),
	)
)

// var ConnectionString = "beac0a503e9041:6b0c0bfd@tcp"

// const AllowedCORSDomain = "http://localhost"
