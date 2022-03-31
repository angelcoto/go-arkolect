package list

import (
	"os"
	"path/filepath"
	"sync"
)

// imprimeSalida lee el buffer de resultados para imprimir cada resultado a terminal.
// La rutina se mantiene en ejecución hasta que el buffer es cerrado.
func imprimeSalida(done chan struct{}) {
	for resultado := range resultados {
		if resultado.err != nil {
			printError(resultado.err)
		} else {
			printLine(resultado.fileProp)
		}
	}
	done <- struct{}{} // Informa a la función de llamado que el trabajo ha finalizado
}

// Definición del buffer para jobs.
type tjob struct {
	fileInfo os.FileInfo
	path     string
}

// Definición del buffer para resultados
type tresultado struct {
	fileProp file
	err      error
}

const totaljobs = 20
const totalworkers = 3
const totalresultados = 10

var jobs = make(chan tjob, totaljobs)
var resultados = make(chan tresultado, totalresultados)

// workerHash lee el buffer "jobs" para obtener el próximo job a ejecutar
// Cuando la tarea ha sido ejecutada el resultado se escribe en el buffer "resultado"
// El worker seguirá ejecutándose mientras el buffer no esté cerrado.
// Cuando el buffer se cierra el worker se declara como finalizado a través de
// wg.Done()
func workerHash(wg *sync.WaitGroup, algo string) {
	fileProp := file{}

	//Recorre el buffer de jobs
	for job := range jobs {

		//hash, err := hash.SumArchivo(job.path, job.algo)
		fileProp, err := fileProp.getFileProp(job.fileInfo, job.path, algo)
		resultado := tresultado{fileProp, err}
		resultados <- resultado
	}
	wg.Done()
}

// creaWorkerPool inicia los worker que estarán leyendo la cola de jobs
func creaWorkerPool(nWorkers int, algo string) {
	var wg sync.WaitGroup
	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go workerHash(&wg, algo)
	}
	wg.Wait()
	close(resultados) // Para indicarle al lector del buffer que no hay más valores a enviar
}

// ListRecursive imprime el hash para los archivos de un directorio,
// incluyendo los subdirectorios.
func ListRecursive(dir string, algo string) {

	done := make(chan struct{})

	go func() {

		i := 0
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if !info.IsDir() {
				jobs <- tjob{info, filepath.Dir(path)}
				i++
			}
			return nil
		})
		close(jobs) // Para indicarle al lector del buffer que no hay más valores a enviar
	}()

	go imprimeSalida(done)
	creaWorkerPool(totalworkers, algo)

	<-done // Genera un bloqueo hasta que ha finalizado la impresión de todos los resultados

}
