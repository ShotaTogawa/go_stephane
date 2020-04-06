package main

import "fmt"

func main() {
	fmt.Println(Hello("Tom", "ja"))
}

const japaneseHElloPrifix = "こんにちは、 "
const englishHelloPrifix = "Hello, "
const spanishHelloPrifix = "Hola, "

// Hello return greeting depends on a language.
func Hello(name string, language string) string {

	// if language == "ja" {
	// 	return japaneseHElloPrifix + name
	// } else if language == "sp" {
	// 	return spanishHelloPrifix + name
	// } else {
	// 	return englishHelloPrifix + name
	// }
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prifix string) {
	switch language {
	case "ja":
		prifix = japaneseHElloPrifix
	case "sp":
		prifix = spanishHelloPrifix
	default:
		prifix = englishHelloPrifix
	}
	return
}
