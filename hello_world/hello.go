package hello

import "fmt"

type Lang struct {
    Greeting string
    World string
}

func Hello(name, lang string) string {
    langMap := getLangMap()
    langStruct, ok := langMap[lang]
    if !ok {
        lang = "english"
    }
    var toPrint string
    if len(name) == 0 {
        toPrint = langStruct.World
    } else {
        toPrint = name
    }
    return fmt.Sprintf("%s %s!", langStruct.Greeting, toPrint)
}

func getLangMap() map[string]Lang {
    langMap := map[string]Lang{
        "english": Lang{"Hello", "World"},
        "german": Lang{"Hallo", "Welt"},
        "greek": Lang{"Geia", "Kosmos"},
    }
    return langMap
}

