package main

import (
	"errors"
	"fmt"
)

type InfoHandler struct{}

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
		return err
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
	handler.readActAndPrint([]string{}, func(s *Sbom, a []string) error {
		return nil
	})
	return nil
}

func (handler InfoHandler) add(args []string) error {
	err := handler.readActAndPrint(args, func(s *Sbom, a []string) error {
		return s.PopulateFromFlags(a)
	})

	if err != nil {
		return err
	}

	return nil
}

func (handler InfoHandler) remove(args []string) error {
	err := handler.readActAndPrint(args, func(s *Sbom, a []string) error {
		return s.ClearFromFlags(a)
	})

	if err != nil {
		return err
	}

	return nil
}

func (handler InfoHandler) readActAndPrint(args []string, action func(*Sbom, []string) error) error {
	var sbom Sbom
	err := sbom.ReadFromFile()
	if err != nil {
		return err
	}

	action(&sbom, args)

	err = sbom.WriteToFile()
	if err != nil {
		return err
	}

	jsonStr, err := sbom.SbomPackageInfo.ToJson()
	if err != nil {
		return err
	}

	fmt.Println(jsonStr)

	return nil
}
