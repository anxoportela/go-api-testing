// Paquete csv contiene funciones para manejar archivos CSV, incluyendo la lectura y escritura de datos.
package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

// EscribirResultados escribe los resultados proporcionados en un archivo CSV sin realizar
// el escapado de comillas ni otros caracteres especiales. La función crea un archivo CSV nuevo
// o sobrescribe uno existente con los datos proporcionados.
//
// Parameters:
//   - results ([][]string): Un slice de slices de strings, donde cada slice interno representa
//     una fila que se escribirá en el archivo CSV.
//   - filename (string): El nombre del archivo CSV en el que se guardarán los resultados.
//
// Returns:
//   - error: Un error en caso de que ocurra algún problema al crear o escribir en el archivo CSV.
func EscribirResultados(results [][]string, filename string) error {
	// Abrir o crear el archivo CSV para escritura
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error al crear el archivo CSV: %v", err)
	}
	defer file.Close()

	// Crear un escritor CSV que manejará la escritura de las filas
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribir cada fila de resultados en el archivo CSV
	for _, record := range results {
		// Escribir la fila como una línea CSV
		// NOTA: Aquí no se realiza un escapado automático de comillas u otros caracteres especiales,
		//       lo que significa que los datos deben estar correctamente formateados antes de ser escritos.
		err := writer.Write(record)
		if err != nil {
			return fmt.Errorf("error al escribir en el archivo CSV: %v", err)
		}
	}

	// Retornar nil si todo salió bien
	return nil
}
