package main

import (
	"fmt"
	"os"
)

// var _ = godotenv.Load(".env")

const IsLocal = true

var (
	RemoteConnectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		os.Getenv("userBD"),
		os.Getenv("passBD"),
		os.Getenv("hostBD"),
		os.Getenv("nameBD"),
	)
)

const LocalConnectionString = "root:root@tcp(localhost:3306)/Agenda"

// const AllowedCORSDomain = "http://localhost"
