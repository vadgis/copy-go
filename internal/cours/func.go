package cours

import (
	"fmt"
	"time"
)

func MafunctionPublic(argumen1 string, argument2 int) (string, error) {
	var random = time.Now().Unix()
	if mafunctionPrivate(int(random)) {
		return "success", nil
	}
	return "", fmt.Errorf("Fail")
}

func mafunctionPrivate(random int) bool {
	result := random % 2
	return result == 0
}