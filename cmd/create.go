package cmd

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const MODULES_PATH = "./client/src/modules/"
const MODULE_TEMPLATE_PATH = "./client/src/modules/module_template.js.tmpl"

const INDEX_PATH = "./client/src/modules/index.js"
const INDEX_TEMPLATE_PATH = "./client/src/modules/index.js.tmpl"

type ModuleDefinition struct {
	Metadata Metadata `json:"metadata"`
	Columns  []Column `json:"columns"`
}

type Metadata struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	ExampleSource string `json:"example_source"`
}

type Column struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	IsURL bool   `json:"is_url"`
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new module from JSON definition",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("[*] Reading module definition file: " + definitionFile)
		defintion, err := readDefinition(definitionFile)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("[*] Checking for name clashes")
		currentModules, err := getCurrentModules()
		if err != nil {
			log.Fatal(err)
		}
		for _, m := range currentModules {
			if m == defintion.Metadata.Name {
				log.Fatalf("A module with the same name exists")
			}
		}

		fmt.Println("[*] Creating module")
		err = createModule(defintion)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("[*] Module created at " + MODULES_PATH + defintion.Metadata.Name + "/index.js")

		fmt.Println("[*] Updating modules index: " + INDEX_PATH)
		err = updateIndex(defintion.Metadata.Name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("[*] Done")
	},
}

var definitionFile string

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&definitionFile, "file", "f", "", "Module definition file")
	createCmd.MarkFlagRequired("file")
}

func readDefinition(filePath string) (ModuleDefinition, error) {
	var definition ModuleDefinition
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return ModuleDefinition{}, err
	}

	err = json.Unmarshal(content, &definition)
	if err != nil {
		return ModuleDefinition{}, err
	}

	return definition, nil
}

func createTemplate(name string) *template.Template {
	funcMap := template.FuncMap{
		"Title": strings.Title,
	}

	t := template.New(name).Funcs(funcMap)
	return t
}

func createModule(definition ModuleDefinition) error {
	content, err := ioutil.ReadFile(MODULE_TEMPLATE_PATH)
	if err != nil {
		return err
	}
	t := string(content)

	temp, err := createTemplate("definition").Parse(t)
	if err != nil {
		return err
	}

	dir := MODULES_PATH + definition.Metadata.Name
	err = os.Mkdir(dir, 0755)
	if err != nil {
		return err
	}
	f, err := os.Create(dir + "/index.js")
	if err != nil {
		return err
	}

	err = temp.Execute(f, definition)
	if err != nil {
		return err
	}
	return nil
}

func updateIndex(moduleName string) error {
	content, err := ioutil.ReadFile(INDEX_TEMPLATE_PATH)
	if err != nil {
		return err
	}
	t := string(content)

	modules, err := getCurrentModules()

	f, err := os.Create(INDEX_PATH)
	if err != nil {
		return err
	}

	temp, err := createTemplate("index").Parse(t)
	if err != nil {
		return err
	}

	err = temp.Execute(f, modules)
	if err != nil {
		return err
	}
	return nil
}

func getCurrentModules() ([]string, error) {
	files, err := ioutil.ReadDir(MODULES_PATH)
	if err != nil {
		return nil, err
	}

	var modules []string
	for _, f := range files {
		if f.IsDir() {
			modules = append(modules, f.Name())
		}
	}

	return modules, nil
}
