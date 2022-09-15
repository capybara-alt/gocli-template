package utils_test

import (
	"testing"

	"github.com/capybara-alt/gocli-template/utils"
)

func TestContainsFoundCase1(t *testing.T) {
	tests := []string{"A", "B", "C", "D"}
	if !utils.Contains(tests, "A") {
		t.Fail()
	}
}

func TestContainsFoundCase2(t *testing.T) {
	tests := []string{"A", "B", "C", "D", "D"}
	if !utils.Contains(tests, "D") {
		t.Fail()
	}
}

func TestContainNotFoundCase(t *testing.T) {
	tests := []string{"A", "B", "C", "D", "D"}
	if utils.Contains(tests, "Z") {
		t.Fail()
	}
}
