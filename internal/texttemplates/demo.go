package texttemplates

import (
	"fmt"
	"html/template"
	"os"
)

func Demo(show bool) {
	if show {
		fmt.Println("-- Text Templates Demo")

		tmpl, err := template.New("demo").Parse("Welcome, {{.name}}! How are you doing?")
		if err != nil {
			panic(err)
		}

		// Define data for welcome message template
		data := map[string]interface{}{
			"name": "John Doe",
		}

		// render the welcome message
		err = tmpl.Execute(os.Stdout, data)
		if err != nil {
			panic(err)
		}
	}
}
