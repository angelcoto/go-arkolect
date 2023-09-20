package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"time"

	"github.com/angelcoto/go-arkolect/list"
)

func header(t time.Time, dir string) {
	usuario, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inventario generado por:", usuario.Username)
	fmt.Println("Ruta: ", dir)
	fmt.Println("Inicio:", t.Format(time.RFC3339))
	fmt.Println("----------------------------------------------------------------")
}

func footer(start time.Time) {
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Fin:", time.Now().Format(time.RFC3339))
	fmt.Println("Duración: ", time.Since(start))
}

const appVersion = "1.1.3"

func main() {

	start := time.Now()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dirPtr := flag.String("d", dir, "Directorio a recorrer")
	algoPtr := flag.String("m", "sha1", "Algoritmo: md5, sha1, sha256")
	recPtr := flag.Bool("r", false, "Recorrido recursivo")
	verPtr := flag.Bool("v", false, "Muestra la versión del programa")
	wrkPtr := flag.Int("w", 1, "Cantidad de workers (entre 1 y 6)")

	flag.Parse()

	// Imprime la versión
	if *verPtr {
		fmt.Println("arkolect versión", appVersion)
		os.Exit(0)
	}

	// Verifica existencia de directorio a recorrer
	if _, err := os.Stat(*dirPtr); os.IsNotExist(err) {
		list.PrintError(err)
		os.Exit(1)
	}

	if *wrkPtr < 1 || *wrkPtr > 6 {
		list.PrintError(errors.New("valor fuera de rango permitido"))
		os.Exit(1)
	}

	header(start, *dirPtr)
	defer footer(start)

	if !*recPtr {
		list.ListDirectory(*dirPtr, *algoPtr)
	} else {
		list.ListRecursive(*dirPtr, *algoPtr, *wrkPtr)
	}

}
