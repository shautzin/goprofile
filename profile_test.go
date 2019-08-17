package goprofile

import (
	"fmt"
	"testing"
)

func TestLoad(T *testing.T) {
	Load()
	fmt.Println(GetEnv("goprofile.server.port"))
}
