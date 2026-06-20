package texttemplates

import (
	"fmt"
	"html/template"
	"os"
)

func TemplateMust(show bool) {
	if show {
		fmt.Println("-- Text Templates Must")

		tmpl := template.Must(template.New("must").Parse("Welcome, {{.name}}! How are you?"))

		// Define data for welcome message template
		data := map[string]interface{}{
			"name": "John Doe",
		}

		// render the welcome message
		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			panic(err)
		}
	}
}
