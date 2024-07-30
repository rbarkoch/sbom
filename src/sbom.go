package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const (
	sbomFilePath = "sbom.json"
)

// Represents a project which requires tracking of dependencies.
type Sbom struct {
	PackageId string `json:"package,omitempty"`
	SbomPackage
}

type SbomPackageInfo struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Version     string `json:"version,omitempty"`
	Type        string `json:"type,omitempty"`
	Author      string `json:"author,omitempty"`
	Company     string `json:"company,omitempty"`
	Copyright   string `json:"copyright,omitempty"`
	Source      string `json:"source,omitempty"`
	Url         string `json:"uri,omitempty"`
	Repository  string `json:"repository,omitempty"`
	Branch      string `json:"branch,omitempty"`
	Commit      string `json:"commit,omitempty"`
	License     string `json:"license,omitempty"`
	LicenseUrl  string `json:"license-url,omitempty"`
}

// Represents a single package which could be a dependency or build tool for a
// project.
type SbomPackage struct {
	SbomPackageInfo
	Packages map[string]*SbomPackage `json:"packages,omitempty"`
}

func (info *SbomPackageInfo) PopulateFromFlags(args []string) error {
	if len(args)%2 != 0 {
		return errors.New("invalid number of flags")
	}

	flags := args
	for len(flags) > 0 {
		flag := flags[0]
		value := flags[1]

		err := info.SetFromFlag(flag, value)
		if err != nil {
			return fmt.Errorf("failed to set field from flag\n  ⤷ %w", err)
		}

		flags = flags[2:]
	}

	return nil
}

func (info *SbomPackageInfo) ClearFromFlags(args []string) error {
	for i := 0; i < len(args); i++ {
		err := info.RemoveFromFlag(args[i])
		if err != nil {
			return fmt.Errorf("failed to clear field from flag\n  ⤷ %w", err)
		}
	}

	return nil
}

func (info *SbomPackageInfo) SetFromFlag(flag string, value string) error {
	switch flag {
	case "--name":
		info.Name = value
	case "--description":
		info.Description = value
	case "--comment":
		info.Comment = value
	case "--version":
		info.Version = value
	case "--type":
		info.Type = value
	case "--author":
		info.Author = value
	case "--company":
		info.Company = value
	case "--copyright":
		info.Copyright = value
	case "--source":
		info.Source = value
	case "--url":
		info.Url = value
	case "--repository":
		info.Repository = value
	case "--branch":
		info.Branch = value
	case "--commit":
		info.Commit = value
	case "--license":
		info.License = value
	case "--license-url":
		info.LicenseUrl = value
	default:
		return fmt.Errorf("unknown flag '%s'", flag)
	}

	return nil
}

func (info *SbomPackageInfo) RemoveFromFlag(flag string) error {
	switch flag {
	case "--name":
		info.Name = ""
	case "--description":
		info.Description = ""
	case "--comment":
		info.Comment = ""
	case "--version":
		info.Version = ""
	case "--type":
		info.Type = ""
	case "--author":
		info.Author = ""
	case "--company":
		info.Company = ""
	case "--copyright":
		info.Copyright = ""
	case "--source":
		info.Source = ""
	case "--url":
		info.Url = ""
	case "--repository":
		info.Repository = ""
	case "--branch":
		info.Branch = ""
	case "--commit":
		info.Commit = ""
	case "--license":
		info.License = ""
	case "--license-url":
		info.LicenseUrl = ""
	default:
		return fmt.Errorf("unknown flag '%s'", flag)
	}

	return nil
}

func (info *SbomPackageInfo) ReadFromFile() error {
	file, err := os.Open(sbomFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return errors.New("an sbom file has not been initialized. use `sbom init <PACKAGE NAME>` to create one")
		} else {
			return fmt.Errorf("failed to open sbom file\n  ⤷ %w", err)
		}
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(info)
	if err != nil {
		return fmt.Errorf("failed to decode sbom from json\n  ⤷ %w", err)
	}

	return nil
}

func (sbom *Sbom) ReadFromFile() error {
	file, err := os.Open(sbomFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return errors.New("an sbom file has not been initialized. use `sbom init <PACKAGE NAME>` to create one")
		} else {
			return fmt.Errorf("failed to open sbom file\n  ⤷ %w", err)
		}
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(sbom)
	if err != nil {
		return fmt.Errorf("failed to decode sbom from json\n  ⤷ %w", err)
	}

	if sbom.Packages == nil {
		sbom.Packages = map[string]*SbomPackage{}
	}

	return nil
}

func (sbom *Sbom) ReadFromFilePath(path string) error {
	file, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return errors.New("an sbom file has not been initialized. use `sbom init <PACKAGE NAME>` to create one")
		} else {
			return fmt.Errorf("failed to open sbom file\n  ⤷ %w", err)
		}
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(sbom)
	if err != nil {
		return fmt.Errorf("failed to decode sbom from json\n  ⤷ %w", err)
	}

	if sbom.Packages == nil {
		sbom.Packages = map[string]*SbomPackage{}
	}

	return nil
}

func (sbom *Sbom) WriteToFile() error {
	file, err := os.OpenFile(sbomFilePath, os.O_WRONLY, os.ModePerm)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return errors.New("an sbom file has not been initialized. use `sbom init <PACKAGE NAME>` to create one")
		} else {
			return fmt.Errorf("failed to open sbom file\n  ⤷ %w", err)
		}
	}
	defer file.Close()

	jsonString, err := sbom.ToJson()
	if err != nil {
		return fmt.Errorf("failed to convert sbom to json\n  ⤷ %w", err)
	}

	file.Truncate(0)
	file.Seek(0, 0)
	_, err = file.WriteString(jsonString)
	if err != nil {
		return fmt.Errorf("failed to write sbom to file\n  ⤷ %w", err)
	}

	return nil
}

func (sbom *Sbom) ToJson() (string, error) {
	jsonData, err := json.MarshalIndent(sbom, "", "    ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal sbom to json\n  ⤷ %w", err)
	}
	return string(jsonData), nil
}

func (info *SbomPackageInfo) ToJson() (string, error) {
	jsonData, err := json.MarshalIndent(info, "", "    ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal sbom info to json\n  ⤷ %w", err)
	}
	return string(jsonData), nil
}
