package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

type clipboardTemplate struct {
	Label    string `json:"label"`
	Template string `json:"template"`
}

type response struct {
	Data               []map[string]interface{} `json:"results"`
	ClipboardTemplates []clipboardTemplate      `json:"clipboard_templates"`
}

var (
	temp                string
	input               string
	clipboard_templates string
	port                string
	dev                 bool
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a module",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[*] Acumen module is running http://127.0.0.1:" + port + "/" + temp)
		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/api/data", getData)

		router.PathPrefix("/").HandlerFunc(Index)

		log.Fatal(http.ListenAndServe(":"+port, router))
	},
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./client/build/index.html")
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&temp, "template", "t", "", "Name of the template")
	runCmd.MarkFlagRequired("template")
	runCmd.Flags().StringVarP(&input, "input", "i", "", "Input file")
	runCmd.MarkFlagRequired("input")
	runCmd.Flags().StringVarP(&clipboard_templates, "clipboard-templates", "c", "", "Copy-to-clipboard templates file")
	runCmd.MarkFlagRequired("clipboard-templates")
	runCmd.Flags().StringVarP(&port, "port", "p", "8888", "Port to run the HTTP server on")
	runCmd.Flags().BoolVarP(&dev, "dev", "d", false, "Run in development mode")
}

func getData(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	var results []map[string]interface{}
	json.Unmarshal(byteValue, &results)

	for i := 0; i < len(results); i++ {
		results[i]["rating"] = 0
		results[i]["id"] = i
	}

	file, err = os.Open(clipboard_templates)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var templates []clipboardTemplate
	tbyteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(tbyteValue, &templates)

	var response response
	response.Data = results
	response.ClipboardTemplates = templates

	w.Header().Set("Content-Type", "application/json")

	if dev {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	}

	json.NewEncoder(w).Encode(response)
}
