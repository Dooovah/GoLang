package utils

import (
	"fmt"
)

func PressEnterToExit() {
	fmt.Println("Press enter to exit...")
	fmt.Scanln()
}
