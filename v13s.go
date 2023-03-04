package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// type definition start
type Options struct {
	SiteName        string `json:"siteName"`
	SiteDescription string `json:"siteDescription"`
	ServerPort      int    `json:"serverPort"`
}

// type definition end

// http request handler function definition start
func entryPointHandler(w http.ResponseWriter, r *http.Request, options Options) {
	template := template.Must(template.ParseFiles("frontend/index.html"))

	siteDescription := func() string {
		if options.SiteName != "" {
			var desc string = "Let’s organise and archive your precious masterpieces."
			if options.SiteDescription != "" {
				desc = options.SiteDescription
			}
			return fmt.Sprintf("%s - %s Powered by Vivliothikarios(v13s)", options.SiteName, desc)
		} else {
			return "Vivliothikarios(v13s) - Let’s organise and archive your precious masterpieces."
		}
	}()

	variables := map[string]string{
		"siteName":        options.SiteName,
		"siteDescription": siteDescription,
	}

	template.ExecuteTemplate(w, "index.html", variables)
}

func optionsRequestHandler(w http.ResponseWriter, r *http.Request, rawOptions []byte) {
	w.Write(rawOptions)
}

// http request handler function definition end

func main() {
	_, err := os.Stat("v13s.config.json")
	if err != nil {
		options := Options{
			SiteName:        "Vivliothikarios",
			SiteDescription: "",
			ServerPort:      9090,
		}

		rawOptions, _ := json.MarshalIndent(options, "", "\t")

		f, _ := os.Create("v13s.config.json")
		f.Write(rawOptions)
	}

	rawOptions, _ := os.ReadFile("v13s.config.json")
	var options Options
	json.Unmarshal(rawOptions, &options)

	func() {
		errVal := func() string {
			if options.SiteName == "" {
				return "siteName"
			} else if options.ServerPort == 0 {
				return "serverPort"
			} else {
				return ""
			}
		}()
		if errVal != "" {
			fmt.Fprintf(os.Stderr, "Error: Option value \"%s\" (at v13s.config.json) is invalid.\n", errVal)
			os.Exit(2)
		}
	}()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/static/"))))
	http.HandleFunc("/api/options/", func(w http.ResponseWriter, r *http.Request) {
		optionsRequestHandler(w, r, rawOptions)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		entryPointHandler(w, r, options)
	})
	http.ListenAndServe(fmt.Sprintf(":%d", options.ServerPort), nil)
}
