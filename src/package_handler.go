package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type PackageHandler struct{}

func (handler PackageHandler) Handle(args []string) error {

	if len(args) == 0 {
		return errors.New("no verb has been provided")
	}

	verb := args[0]
	tail := args[1:]

	var err error
	switch verb {
	case "ls":
		err = handler.print(tail)
	case "add":
		err = handler.add(tail)
	case "rm":
		err = handler.remove(tail)
	default:
		err = errors.New("invalid verb")
	}

	if err != nil {
		return err
	}

	return nil
}

func (handler PackageHandler) Help() {
	fmt.Println(`Manages package dependencies of the project.
	
Print all packages.

    sbom package ls

Print specific package information.

    sbom package ls <PACKAGE NAME>

Add a package.

    sbom package add <PACKAGE NAME> [ARGUMENTS]

Remove a package.

    sbom package rm <PACKAGE NAME>


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

func (handler PackageHandler) print(args []string) error {
	var sbom Sbom
	err := sbom.ReadFromFile()
	if err != nil {
		return err
	}

	var jsonData []byte
	if len(args) == 0 {

		// No package name provided. Print all packages.
		jsonData, err = json.MarshalIndent(sbom.Packages, "", "    ")

	} else {

		// Package name was provided. Attempt to print just that package.
		if len(args) > 1 {
			return errors.New("invalid number of arguments")
		}

		pkgName := args[0]
		_, ok := sbom.Packages[pkgName]
		if !ok {
			return errors.New("package does not exist")
		}

		dto := map[string]*SbomPackage{
			pkgName: sbom.Packages[pkgName],
		}

		jsonData, err = json.MarshalIndent(dto, "", "    ")

	}

	if err != nil {
		return err
	}

	fmt.Println(string(jsonData))

	return nil
}

func (handler PackageHandler) add(args []string) error {
	var sbom Sbom
	err := sbom.ReadFromFile()
	if err != nil {
		return err
	}

	if len(args) == 0 {
		return errors.New("must provide package name")
	}

	pkgName := args[0]
	_, ok := sbom.Packages[pkgName]
	var pkg *SbomPackage
	if ok {
		if len(args) == 1 {
			return errors.New("package already exists. either remove the package first or provide arguments to update the existing package")
		}

		pkg = sbom.Packages[pkgName]
	} else {
		pkg = new(SbomPackage)
	}

	err = pkg.PopulateFromFlags(args[1:])
	if err != nil {
		return err
	}

	sbom.Packages[pkgName] = pkg

	err = sbom.WriteToFile()
	if err != nil {
		return err
	}

	dto := map[string]*SbomPackage{
		pkgName: pkg,
	}

	jsonData, err := json.MarshalIndent(dto, "", "    ")
	if err != nil {
		return err
	}

	fmt.Println(string(jsonData))

	return nil
}

func (handler PackageHandler) remove(args []string) error {
	var sbom Sbom
	err := sbom.ReadFromFile()
	if err != nil {
		return err
	}

	if len(args) == 0 {
		return errors.New("must provide package name")
	}

	pkgName := args[0]
	_, ok := sbom.Packages[pkgName]
	if !ok {
		return errors.New("package does not exist")
	}

	// Delete the entire package.
	if len(args) == 1 {
		delete(sbom.Packages, pkgName)
		err = sbom.WriteToFile()
		if err != nil {
			return err
		}
		return nil
	}

	// Delete properties from the package.
	sbom.Packages[pkgName].ClearFromFlags(args[1:])

	err = sbom.WriteToFile()
	if err != nil {
		return err
	}

	dto := map[string]*SbomPackage{
		pkgName: sbom.Packages[pkgName],
	}

	jsonData, err := json.MarshalIndent(dto, "", "    ")
	if err != nil {
		return err
	}

	fmt.Println(string(jsonData))

	return nil
}
