package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("123456"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$iUB1jHktKheSm6e7v8ofgOAabZLrCLt7Nzq810TaKURR1.Gb5svES", "123456"))
}