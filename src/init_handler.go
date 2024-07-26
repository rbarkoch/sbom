package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
)

type InitHandler struct{}

func (handler InitHandler) Handle(args []string) error {
	if len(args) == 0 {
		handler.Help()
		return errors.New("a package name must be provided")
	}

	// Validate package name.
	packageId := args[0]
	match, err := regexp.MatchString(`^[a-z0-9\-\.]+$`, packageId)
	if err != nil {
		return err
	}

	if !match {
		return errors.New("invalid package id. name must only contain lower-case letters, numbers, dashes, and periods")
	}

	// Check if an sbom file already exists.
	_, err = os.Stat("sbom.json")
	if err == nil {
		return errors.New("an sbom file already exists")
	}

	sbom := Sbom{
		PackageId: packageId,
		SbomPackage: SbomPackage{
			Packages: map[string]*SbomPackage{},
		},
	}

	err = sbom.PopulateFromFlags(args[1:])
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(sbom, "", "    ")
	if err != nil {
		return err
	}

	file, err := os.Create("sbom.json")
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	fmt.Println(string(jsonData))

	return nil
}

func (handler InitHandler) Help() {
	fmt.Println(`Creates a new software bill-of-materials file for your package.
	
Create a new SBOM.

    sbom init <PACKAGE NAME> [ARGUMENTS]


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
