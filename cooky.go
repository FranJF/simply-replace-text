package main

import (
    "bufio"
    "strings"
    "fmt"
    "os"
    "regexp"
    "github.com/AlecAivazis/survey/v2"
    "path/filepath"
)

func StringPrompt(label string) string {
    var s string
    r := bufio.NewReader(os.Stdin)
    for {
        fmt.Fprint(os.Stderr, label+" ")
        s, _ = r.ReadString('\n')
        if s != "" {
            break
        }
    }
    return strings.TrimSpace(s)
}

func suggestFiles(toComplete string) []string {
    files, _ := filepath.Glob("plantillas/" + "*")
    return files
}

var q = []*survey.Question{
	{
		Name: "file",
		Prompt: &survey.Input{
			Message: "Que archivo quieres utilizar?",
			Suggest: suggestFiles,
			Help:    "Si no encuentras el archivo recuerda ponerlo en la carpeta llamada plantillas",
		},
		Validate: survey.Required,
	},
}

func preguntarArchivo()string{
    	answers := struct {
		File string
	}{}

	err := survey.Ask(q, &answers)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
    return answers.File
}


func main() {
    archivo := preguntarArchivo()

    b, err := os.ReadFile(archivo)
    if err != nil {
        fmt.Print("No existe el archivo.")
    }
    str := string(b)

    re := regexp.MustCompile(`\{(.*?)\}`)

    var elementos []string
    var new_elementos []string

    submatchall := re.FindAllString(str, -1)
	for _, element := range submatchall {
		element = strings.Trim(element, "{")
		element = strings.Trim(element, "}")
        new_element := StringPrompt(element)
        elementos = append(elementos, "{"+element+"}")
        new_elementos = append(new_elementos,new_element)
	}

    if len(elementos) == 0 {
        panic("No se ha encontrado ninguna palabra en el archivo entre llaves {}.")
    }
    if len(elementos) != len(new_elementos) {
        panic("Ha ocurrido un error.")
    }

    new_str := ""
	for index, palabra_a_sustituir := range elementos {
        if index > 0 {
            str = new_str
        }
        palabra_a_poner := new_elementos[index] 
        new_str = strings.Replace(str,palabra_a_sustituir, palabra_a_poner, -1)
    }

    fmt.Println(new_str)

}
