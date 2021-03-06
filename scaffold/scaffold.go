// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package scaffold

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	textTemplate "text/template"
)

// A Scaffold holds data about how/where generate the scaffold of a new
// integration
type Scaffold struct {
	DestinationPath string
	Date            string
	Integration     Integration
}

// An Integration holds all the necessary data to generate an on-host
// integration
type Integration struct {
	Name            string
	Description     string
	ProtocolVersion int
	OS              string
	Prefix          string
	Interval        int
	CompanyPrefix   string
	CompanyName     string
	BinaryName      string
	EventType       string
	EntityType      string
}

type template struct {
	Path       string
	OutputPath string
	FileMode   os.FileMode
}

var (
	gometalinterPath = "resource/tmpl/.gometalinter.json.tmpl"

	commonTemplates = []template{
		{"resource/tmpl/README.md.tmpl", "README.md", 0644},
		{"resource/tmpl/CHANGELOG.md.tmpl", "CHANGELOG.md", 0644},
		{"resource/tmpl/definition.yml.tmpl", "{{ .CompanyPrefix }}-{{ .Name }}-definition.yml", 0644},
		{"resource/tmpl/configuration.yml.tmpl", "{{ .CompanyPrefix }}-{{ .Name }}-config.yml", 0644},
		{"resource/tmpl/configuration.yml.template.tmpl", "{{ .CompanyPrefix }}-{{ .Name }}-config.yml.template", 0644},
		{"resource/tmpl/Makefile.tmpl", "Makefile", 0644},
		{gometalinterPath, ".gometalinter.json", 0644},
		{"resource/tmpl/vendor/vendor.json.tmpl", "vendor/vendor.json", 0644},
	}
	localTemplates = []template{
		{"resource/tmpl/src/integration_local.go.tmpl", "src/{{ .Name }}.go", 0644},
		{"resource/tmpl/src/integration_test_local.go.tmpl", "src/{{ .Name }}_test.go", 0644},
	}
	remoteTemplates = []template{
		{"resource/tmpl/src/integration_remote.go.tmpl", "src/{{ .Name }}.go", 0644},
		{"resource/tmpl/src/integration_test_remote.go.tmpl", "src/{{ .Name }}_test.go", 0644},
	}
)

// InitVendoring setup the vendoring of the integration dependencies using
// 'govendor' tool.
func (s *Scaffold) InitVendoring() error {
	fmt.Println("Fetching external dependencies...")
	cmd := exec.Command("make", "compile")
	cmd.Dir = s.DestinationPath

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("cannot fetch external dependencies. %s failed with: %s", "make compile-deps", output)
	}

	return nil
}

func createIntegrationDirectories(destinationPath string, verbose bool) error {
	for _, path := range []string{destinationPath, destinationPath + "/src", destinationPath + "/vendor"} {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("cannot create directory '%s': \"%s\"", path, err)
		}
		if verbose {
			fmt.Printf("%s\n", path)
		}
	}
	return nil
}

func createIntegrationFile(directoryPath string, tmpl template, integration Integration, verbose bool) (*os.File, error) {
	// Prepare the parsing of file path
	outputPathTmpl, err := textTemplate.New("").Parse(tmpl.OutputPath)
	if err != nil {
		panic(err)
	}
	var outputPathBuf bytes.Buffer
	// Execute the parsing of file path
	err = outputPathTmpl.Execute(&outputPathBuf, integration)
	if err != nil {
		panic(err)
	}
	// Create the file
	outputFilePath := filepath.Join(directoryPath, outputPathBuf.String())
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return nil, fmt.Errorf("cannot create file '%s'. \"%s\"", outputFilePath, err)
	}
	// Change permissions
	err = os.Chmod(outputFilePath, tmpl.FileMode)
	if err != nil {
		return nil, fmt.Errorf("cannot change the permissions of '%s'. \"%s\"", outputFilePath, err)
	}
	if verbose {
		fmt.Printf("%s\n", outputFilePath)
	}
	return outputFile, nil
}

func writeTemplateContent(tmplPath string, outputFile *os.File, scaffold *Scaffold) error {
	// Load template file
	templateData, err := Asset(tmplPath)
	if err != nil {
		return fmt.Errorf("cannot read template file '%s'. \"%s\"", tmplPath, err)
	}
	// Parse file content
	outputTmpl, err := textTemplate.New("").Parse(string(templateData))
	if err != nil {
		return fmt.Errorf("cannot parse template file '%s'. \"%s\"", tmplPath, err)
	}
	err = outputTmpl.Execute(outputFile, scaffold)
	if err != nil {
		return fmt.Errorf("cannot write content for file '%s'. \"%s\"", outputFile.Name(), err)
	}
	return nil
}

// Generate creates all the directories and files of the scaffold
func (s *Scaffold) Generate(verbose bool) error {
	fmt.Printf("Generating scaffold for '%s' integration...\n", s.Integration.Name)
	exist, _ := exists(s.DestinationPath)
	if exist {
		return fmt.Errorf("Directory '%s' already exists", s.DestinationPath)
	}
	err := createIntegrationDirectories(s.DestinationPath, verbose)
	if err != nil {
		return err
	}

	if s.Integration.EntityType == "remote" {
		err = writeTemplates(remoteTemplates, s, verbose)
		if err != nil {
			return err
		}
	} else if s.Integration.EntityType == "local" {
		err = writeTemplates(localTemplates, s, verbose)
		if err != nil {
			return err
		}
	}

	err = writeTemplates(commonTemplates, s, verbose)
	if err != nil {
		return err
	}

	fmt.Printf("Integration scaffold successfully created in '%s/'\n", s.DestinationPath)
	return nil
}

// write templates
func writeTemplates(templates []template, s *Scaffold, verbose bool) error {
	for _, tmpl := range templates {
		outputFile, err := createIntegrationFile(s.DestinationPath, tmpl, s.Integration, verbose)
		if err != nil {
			return err
		}
		err = writeTemplateContent(tmpl.Path, outputFile, s)
		if err != nil {
			return err
		}
	}
	return nil
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
