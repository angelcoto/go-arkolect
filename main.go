package main

import (
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

func main() {

	start := time.Now()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dirPtr := flag.String("d", dir, "Directorio a recorrer")
	algoPtr := flag.String("m", "sha1", "Algoritmo: md5, sha1, sha256")
	recPtr := flag.Bool("r", false, "Recorrido recursivo")

	flag.Parse()

	header(start, *dirPtr)
	defer footer(start)

	if !*recPtr {
		list.ListDirectory(*dirPtr, *algoPtr)
	} else {
		list.ListRecursive(*dirPtr, *algoPtr)
	}

}
