package encrypt

import (
	"fmt"
	"go_template/pkg/config"
	"testing"
)

func TestStringEncrypt(t *testing.T) {
	config.Init()
	p, err := StringEncrypt("mysql")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(p)
}

func TestStringDecrypt(t *testing.T) {
	p, err := StringDecrypt("AisBAQEBAQHkhF/z8qfr3CpI9L7oPas30bQ2/na0Egg=")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(p)
}
