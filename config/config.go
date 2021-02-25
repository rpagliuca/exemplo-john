package config

import (
	"fmt"
	"john/config/schema"
)

func Hello() {
	fmt.Println("Hello,", schema.Provider)
}
