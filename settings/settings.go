// Copyright 2015 Bruno Albuquerque (bga@bug-br.org.br). All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

// Package settings provides simple settings handling. It handles settings
// files with the following format:
//
// setting1:value1
// setting2:value2
// [...]
//
// In other words, each line should have a single settings/value pair separated
// by ":". It does not handle extra spaces or comments for now so be careful.
package settings

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Settings represents all the configuration values in a specific file (or to
// be written to a specific file).
type Settings struct {
	filePath string

	settingsMap map[string]string
}

// New creates a new Settings instance by trying to load the settings from the
// given fiePath. If filePath does not exist, it returns an empty Settings
// instance that can have settings added to it and written to the given
// filePath when Save() is called. In case of any error, a nil Settings is
// returned together with a non-nil error.
func New(filePath string) (*Settings, error) {
	settingsMap, err := loadFile(filePath)
	if err != nil {
		return nil, err
	}

	return &Settings{
		filePath,
		settingsMap,
	}, nil
}

// Get returns the value associate with the given setting. In case the setting
// is not found, it returns a non-nil error.
func (s *Settings) Get(setting string) (string, error) {
	value, ok := s.settingsMap[setting]
	if !ok {
		return "", fmt.Errorf("setting %q not found", setting)
	}

	return value, nil
}

// Set overrides the given setting with the given value. If the settings does
// not exist yet, it will be created.
func (s *Settings) Set(setting, value string) {
	s.settingsMap[setting] = value
}

// Save writes the configured settings to the filePath given when New() was
// called. It returns a non-nil error in case any error is detected.
func (s *Settings) Save() error {
	if len(s.settingsMap) == 0 {
		return fmt.Errorf("no settings to write")
	}

	f, err := os.OpenFile(s.filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0600)
	if err != nil {
		return err
	}

	defer f.Close()

	for key, value := range s.settingsMap {
		_, err = f.WriteString(key + ":" + value + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

// loadFile tries to load settings from the given filePath. It returns a
// non-nil error in case of failure.
func loadFile(filePath string) (map[string]string, error) {
	settingsMap := make(map[string]string)

	f, err := os.Open(filePath)
	if os.IsNotExist(err) {
		// File does not exist. That is fine.
		return settingsMap, nil
	}

	if err != nil {
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if len(strings.TrimSpace(scanner.Text())) == 0 {
			// Empty line. Ignore it.
			continue
		}

		parsedLine := strings.Split(scanner.Text(), ":")
		if len(parsedLine) != 2 {
			return nil, fmt.Errorf(
				"invalid line %q in settings file",
				scanner.Text())
		}
		settingsMap[parsedLine[0]] = parsedLine[1]

	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return settingsMap, nil
}
