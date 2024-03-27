package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetCommitSha(t *testing.T) {
	sha := "abcd1234"
	ioutil.WriteFile("sha.txt", []byte(sha), 0644)
	defer os.Remove("sha.txt")
	result := getCommitSha()
	if result != sha {
		t.Errorf("getCommitSha() = %s; want %s", result, sha)
	}
}

func TestGetColor(t *testing.T) {
	sha := "abcd1234"
	expectedColor := "#e9cee7"
	if result := getColor(sha); result != expectedColor {
		t.Errorf("getColor() = %s; want %s", result, expectedColor)
	}
}

func TestGetTextColor(t *testing.T) {
	tests := []struct {
		backgroundColor string
		expectedColor   string
	}{
		{"#FFFFFF", "#000000"},
		{"#000000", "#FFFFFF"},
	}

	for _, test := range tests {
		if result := getTextColor(test.backgroundColor); result != test.expectedColor {
			t.Errorf("getTextColor(%s) = %s; want %s", test.backgroundColor, result, test.expectedColor)
		}
	}
}

func TestGetEnvironment(t *testing.T) {
	os.Setenv("ENVIRONMENT", "production")
	expectedEnvironment := "production"
	if result := getEnvironment(); result != expectedEnvironment {
		t.Errorf("getEnvironment() = %s; want %s", result, expectedEnvironment)
	}
}
