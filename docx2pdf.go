package main

import (
	// "fmt"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	// "os/exec"
)

func main() {
	// arg0 := "libreoffice"
	// arg1 := "--headless"
	// arg2 := "--convert-to"
	// arg3 := "pdf"

	// path := "/home/ada/Загрузки/Ануфриев Василий Иванович.docx"

	currDir := "/home/ada/Загрузки/ИНТЕРАКТИВНАЯ СТЕНА ПАМЯТИ"
	outDir := "/home/ada/Загрузки/pdfs"

	files, err := os.ReadDir(currDir)

	if err != nil {
		panic(err)
	}


	convertToPdfAll(files, currDir, outDir)
}

func convertToPdfAll(files []fs.DirEntry, path string, outDir string)  {
	for _, file := range files {
		if file.IsDir() {
			outDir = filepath.Join(outDir, file.Name())
			os.Mkdir(outDir, 0777)
			path = filepath.Join(path, file.Name())
			files, err := os.ReadDir(path)
			if err != nil {
				panic(err)
			}
			convertToPdfAll(files, path, outDir)
			fmt.Println("MK DIR")
		} else {
			
			convertToPdf(filepath.Join(path, file.Name()), outDir)
			fmt.Printf("Convert :%v\n", path)
		}	
	}
}

func convertToPdf(path string, outPath string) {
	arg0 := "libreoffice"
	arg1 := "--headless"
	arg2 := "--convert-to"
	arg3 := "pdf"
	arg4 := "--outdir"


	_, err := exec.Command(arg0, arg1, arg2, arg3, path, arg4, outPath).Output()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", outPath)
}


