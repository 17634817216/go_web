package utils

import (
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetStatmPath() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(" Error:", err)
	}
	return cwd
}
