package encryption

import (
	"testing"
)

func TestTextEncrypt(t *testing.T) {
	text := "1ye7hgb!@#K(EJ12rb42<!@ijdi1g"
	pass := "!@(EJ*D(H#@FG6qrniduwge#@R_O#$)K(gjvna"

	t.Log("Text Original: ", text)

	x := TextEncrypt(text, pass)

	t.Log("Text Encrypted: ", x)

	d := TextDecrypt(x, pass)

	if d != text {
		t.Error("Texts not match!")
	} else {
		t.Log("Text Match!")
	}
}
