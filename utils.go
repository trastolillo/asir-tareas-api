package main

import "strconv"

func stringToInt(s string) (int, error) {
	numero, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, err
	}
	return int(numero), err
}
