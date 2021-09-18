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
		os.Getenv("hostBD"),
		os.Getenv("PORT"),
		os.Getenv("nameBD"),
	)
)

// var ConnectionString = "beac0a503e9041:6b0c0bfd@tcp"

// const AllowedCORSDomain = "http://localhost"
