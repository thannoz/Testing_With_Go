package main

func Hello(name, lang string) (prefix string) {
	switch lang {
	case "Spanish":
		prefix = "Hola, " + name
	case "French":
		prefix = "Bonjour, " + name
	default:
		prefix = "Hello, " + name
	}
	return
}
