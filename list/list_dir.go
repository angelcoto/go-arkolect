package list

import (
	"fmt"
	"io/ioutil"
)

// SumDirectorio imprime el hash para los archivos
// de un directorio, sin incluir los subdirectorios
func ListDirectorio(dir string, algo string) {

	archivos, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Printf("* Error: %s\n", err)
	}

	fileLine := file{}
	for _, f := range archivos {
		if !f.IsDir() {

			// Se usa la ruta completa para poder localizar el archivo
			// al momento de calcular el hash

			fileLine, err := fileLine.getFileProp(f, dir, algo)
			if err != nil {
				fmt.Printf("* Error: %s\n", err)
			} else {
				printLine(fileLine)
			}

		}
	}
}
