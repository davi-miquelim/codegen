package gen 

func GenerateMainFile(path string) {
	sourceCode := `
		package main

		func main() {

		}

	`

	os.WriteFile(path, []byte(sourceCode), 0644)
}
