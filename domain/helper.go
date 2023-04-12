package domain

import (
	"strconv"
	"time"
)

func StringToInt(str string) int {
	i, _ := strconv.Atoi(str)

	return i
}

func IntToString(i int) string {
	str := strconv.Itoa(i)

	return str
}

func TimeToString(t time.Time) string {
	str := t.Format("2006-01-02 15:04:05")
	return str
}
