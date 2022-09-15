package utils_test

import (
	"testing"

	"github.com/capybara-alt/gocli-template/utils"
)

func TestSnakeUpperToCamelCase(t *testing.T) {
	tests := "THIS_IS_TEST_A"
	result := utils.SnakeUpperToCamelCase(tests)
	if result != "ThisIsTestA" {
		t.Fail()
	}
}
func TestSnakeUpperToCamelCaseWithCharacter(t *testing.T) {
	tests := "THIS_IS_TEST_@A"
	result := utils.SnakeUpperToCamelCase(tests)
	if result != "ThisIsTest@a" {
		t.Fail()
	}
}

func TestCamelToSnakeUpperCase(t *testing.T) {
	tests := "ThisIsTestA"
	result := utils.CamelToSnakeUpperCase(tests)
	if result != "THIS_IS_TEST_A" {
		t.Fail()
	}
}

func TestCamelToSnakeUpperCaseWithCharacter(t *testing.T) {
	tests := "ThisIsTest@a"
	result := utils.CamelToSnakeUpperCase(tests)
	if result != "THIS_IS_TEST_@A" {
		t.Fail()
	}
}
