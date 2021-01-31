package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComparePasswords(t *testing.T) {

	var defaultPass = "123"

	ok := ComparePasswords(defaultPass, defaultPass)
	assert.Equal(t, false, ok)

	hashPass, err := HashAndSalt([]byte(defaultPass))

	if err != nil {
		t.Fatalf(err.Error())
	}

	ok = ComparePasswords(defaultPass, hashPass)
	assert.Equal(t, true, ok)

	hashPass2, err := HashAndSalt([]byte("321"))

	if err != nil {
		t.Fatalf(err.Error())
	}

	ok = ComparePasswords(defaultPass, hashPass2)
	assert.Equal(t, false, ok)

}


