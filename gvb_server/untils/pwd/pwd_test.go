package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("123456"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$TUpETNGGRDzD5FeH0tqzYO8UQSa/TBbe/8.4uHd5A0bHo7paeAvlO", ""))
}