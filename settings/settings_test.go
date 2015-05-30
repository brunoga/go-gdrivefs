package settings

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	// Inexistent file.
	s, err := New("./doesnotexist")
	if err != nil {
		t.Fatalf("Expected no error on non-existing file. Got %q.", err)
	}

	if len(s.settingsMap) != 0 {
		t.Errorf("Expected no settings. Got %d settings.", len(s.settingsMap))
	}

	// Invalid settings.
	s, err = New("./invalid_settings.txt")
	if err == nil {
		t.Fatalf("Expected error but got none.")
	}

	if s != nil {
		t.Fatalf("Expected nil Settings but got non-nil one.")
	}

	// Calid settings.
	s, err = New("./settings.txt")
	if err != nil {
		t.Fatalf("Expected no error. Got %q.", err)
	}

	if len(s.settingsMap) != 2 {
		t.Fatalf("Expected 2 settings. Got %d settings.", len(s.settingsMap))
	}

	setting1, ok := s.settingsMap["setting1"]
	if !ok {
		t.Errorf("setting1 setting not found in file.")
	}

	if setting1 != "value1" {
		t.Errorf("setting1 has incorrect value. Expected 'value1', "+
			"got %q.", setting1)
	}

	setting2, ok := s.settingsMap["setting2"]
	if !ok {
		t.Errorf("setting2 setting not found in file.")
	}

	if setting2 != "value2" {
		t.Errorf("setting2 has incorrect value. Expected 'value1', "+
			"got %q.", setting2)
	}
}

func TestGet(t *testing.T) {
	s, err := New("./settings.txt")
	if err != nil {
		t.Fatalf("Expected no error. Got %q.", err)
	}

	value, err := s.Get("setting2")
	if err != nil {
		t.Errorf("setting2 setting not found.")
	}

	if value != "value2" {
		t.Errorf("incorrect setting2 value found. Expected 'value2', got %q.", value)
	}

	value, err = s.Get("setting3")
	if err == nil {
		t.Errorf("Found unexpected setting3 setting.")
	}
}

func TestSet(t *testing.T) {
	s, err := New("./doesnotexist")
	if err != nil {
		t.Fatalf("Expected no error on non-existing file. Got %q.", err)
	}

	if len(s.settingsMap) != 0 {
		t.Errorf("Expected no settings. Got %d settings.", len(s.settingsMap))
	}

	s.Set("setting1", "value1")

	if len(s.settingsMap) != 1 {
		t.Errorf("Expected 1 setting. Got %d settings.", len(s.settingsMap))

	}

	value, ok := s.settingsMap["setting1"]
	if !ok {
		t.Errorf("setting1 setting not found.")
	}

	if value != "value1" {
		t.Errorf("incorrect setting1 value found. Expected 'value1', got %q.", value)
	}
}

func TestSave(t *testing.T) {
	s, err := New("./new_settings.txt")
	if err != nil {
		t.Fatalf("Expected no error on non-existing file. Got %q.", err)
	}

	if len(s.settingsMap) != 0 {
		t.Errorf("Expected no settings. Got %d settings.", len(s.settingsMap))
	}

	err = s.Save()
	if err == nil {
		t.Errorf("Expected error whwn writting file with no settings. Got none.")
	}

	s.Set("setting1", "value1")

	s.Save()

	_, err = os.Stat("./new_settings.txt")
	if err != nil {
		t.Fatalf("Expected no error when checking settings file. Got %q.", err)
	}
	defer os.Remove("./new_settings.txt")

	s2, err := New("./new_settings.txt")
	if err != nil {
		t.Fatalf("Expected no error reading settings file. Got %q.", err)
	}

	value, err := s2.Get("setting1")
	if err != nil {
		t.Errorf("Expected no error fetting setting1. Got %q.", err)
	}

	if value != "value1" {
		t.Errorf("incorrect setting1 value found. Expected 'value1', got %q.", value)
	}
}
