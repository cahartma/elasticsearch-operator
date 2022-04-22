package esclient

import (
	"testing"

	"github.com/ViaQ/logerr/v2/log"
)

func TestHeaderGenEmptyToken(t *testing.T) {
	logger := log.NewLogger("client-testing")
	tokenFile := "../../../test/files/emptyToken"

	_, ok := readSAToken(logger, tokenFile)

	if ok {
		t.Errorf("Expected to be unable to read file [%s] due to empty file but could", tokenFile)
	}
}

func TestHeaderGenWithToken(t *testing.T) {
	logger := log.NewLogger("client-testing")
	tokenFile := "../../../test/files/testToken"

	expected := "test\n"

	actual, ok := readSAToken(logger, tokenFile)

	if !ok {
		t.Errorf("Expected to be able to read file [%s] but couldn't", tokenFile)
	} else {
		if actual != expected {
			t.Errorf("Expected %q but got %q", expected, actual)
		}
	}
}

func TestHeaderGenWithNoToken(t *testing.T) {
	logger := log.NewLogger("client-testing")
	tokenFile := "../../../test/files/errorToken"

	_, ok := readSAToken(logger, tokenFile)

	if ok {
		t.Errorf("Expected to be unable to read file [%s]", tokenFile)
	}
}
