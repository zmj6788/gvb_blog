package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("123456"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$WSwsp/mW0iDzjE.fa9lUU.Ipxx9/m2S71zAXmmBQKgCFCDur0Kh9S", "zmj"))
}