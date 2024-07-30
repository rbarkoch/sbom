package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type InfoHandler struct{}

type InfoDto struct {
	PackageId string `json:"package,omitempty"`
	SbomPackageInfo
}

func (handler InfoHandler) Handle(args []string) error {

	if len(args) == 0 {
		return errors.New("no verb has been provided")
	}

	verb := args[0]
	tail := args[1:]

	var err error
	switch verb {
	case "ls":
		err = handler.print()
	case "add":
		err = handler.add(tail)
	case "rm":
		err = handler.remove(tail)

	default:
		err = errors.New("invalid verb")
	}

	if err != nil {
		return fmt.Errorf("failed to handle info verb\n  ⤷ %w", err)
	}

	return nil
}

func (handler InfoHandler) Help() {
	fmt.Println(`Manages package dependencies of the project.
	
Print package information.

    sbom info ls

Print specific package information.

    sbom info add [ARGUMENTS]

Add a package.

    sbom info rm [ARGUMENTS]


Arugments:

    --name <STRING>
    --description <STRING>
    --comment <STRING>
    --version <STRING>
    --type <STRING>
    --author <STRING>
    --company <STRING>
    --copyright <STRING>
    --source <STRING>
    --url <STRING>
    --repository <STRING>
    --branch <STRING>
    --commit <STRING>
    --license <STRING>
    --license-url <STRING>

	`)
}

func (handler InfoHandler) print() error {
	var sbom Sbom
	err := sbom.ReadFromFile()
	if err != nil {
		return fmt.Errorf("failed to read sbom from file\n  ⤷ %w", err)
	}

	dto := InfoDto{
		PackageId:       sbom.PackageId,
		SbomPackageInfo: sbom.SbomPackageInfo,
	}

	jsonData, err := json.MarshalIndent(dto, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to serialize sbom info\n  ⤷ %w", err)
	}

	fmt.Println(string(jsonData))

	return nil
}

func (handler InfoHandler) add(args []string) error {
	var sbom Sbom
	err := sbom.ReadFromFile()
	if err != nil {
		return fmt.Errorf("failed to read sbom from file\n  ⤷ %w", err)
	}

	err = sbom.PopulateFromFlags(args)
	if err != nil {
		return fmt.Errorf("failed to populate info from arguments\n  ⤷ %w", err)
	}

	err = sbom.WriteToFile()
	if err != nil {
		return fmt.Errorf("failed to write sbom file\n  ⤷ %w", err)
	}

	dto := InfoDto{
		PackageId:       sbom.PackageId,
		SbomPackageInfo: sbom.SbomPackageInfo,
	}

	jsonData, err := json.MarshalIndent(dto, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to serialize sbom info\n  ⤷ %w", err)
	}

	fmt.Println(string(jsonData))

	return nil
}

func (handler InfoHandler) remove(args []string) error {
	var sbom Sbom
	err := sbom.ReadFromFile()
	if err != nil {
		return fmt.Errorf("failed to read sbom from file\n  ⤷ %w", err)
	}

	err = sbom.ClearFromFlags(args)
	if err != nil {
		return fmt.Errorf("failed to clear sbom information from arguments\n  ⤷ %w", err)
	}

	err = sbom.WriteToFile()
	if err != nil {
		return fmt.Errorf("failed to write sbom file\n  ⤷ %w", err)
	}

	dto := InfoDto{
		PackageId:       sbom.PackageId,
		SbomPackageInfo: sbom.SbomPackageInfo,
	}

	jsonData, err := json.MarshalIndent(dto, "", "    ")

	if err != nil {
		return fmt.Errorf("failed to serialize sbom info\n  ⤷ %w", err)
	}

	fmt.Println(string(jsonData))

	return nil
}
