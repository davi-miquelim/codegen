package gen

import (
	"golang.org/x/mod/modfile"
)

func GenProjectFolders(projectPath string) error {
	dirs := []string{"handlers", "middlewares", "models", "base", "cmd", "db"}

	fmt.Println("Generating project structure...")
	for dir := range(dirs) {
		fmt.Println("Creating %s directory", dir)
		err := os.Mkdir(fmt.Sprintf("%s/%s", porjectPath, dir))
		if err != nil {
			return err
		}
	}
}

// TODO: handle errors properly
func GenProjectGoMod(projectPath, projectName string) error {
	modulePath := fmt.Sprintf("github.com/%s", projectName)
	goVersion := "1.23.0"
	filePath := fmt.Sprintf("%s/go.mod", projectPath)

	f := modfile.Newfile(filePath, modulePath, nil)
	if err := f.AddModuleStmt(modulePath); err != nil {
		return err
	}
	if err := f.AddGoStmt(goVersion); err != nil {
		return err
	}

	f.Cleanup()
	formattedBytes, err := f.Format()
	if err != nil {
		return err
	}

	err = os.WriteFile(modulePath, formattedBytes, 0644)
	if err != nil {
		return err
	}

	fmt.Println("go.mod generated")
}

func GenProjectFramework(projectPath, projectName string) {
	GenProjectFolders(projectPath)
	GenProjectGoMod(projectName)
}
