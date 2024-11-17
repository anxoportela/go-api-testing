// Paquete csv contiene funciones para leer archivos CSV y convertirlos en estructuras de datos.
package csv

import (
	"encoding/csv"
	"go-api-testing/models"
	"os"
	"strconv"
)

// LeerCSV lee los casos de prueba desde un archivo CSV ubicado en la ruta especificada.
// La función convierte cada línea del archivo (excepto la primera línea de cabeceras) en una
// estructura TestCase y devuelve un slice de estas estructuras.
//
// Parameters:
//   - path (string): La ruta del archivo CSV que contiene los casos de prueba.
//
// Returns:
//   - []models.TestCase: Un slice de estructuras TestCase que representan los casos de prueba.
//   - error: Un error en caso de que haya un problema al abrir o leer el archivo CSV.
func LeerCSV(path string) ([]models.TestCase, error) {
	// Abrir el archivo CSV
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Crear un lector CSV
	reader := csv.NewReader(file)
	// Leer todas las líneas del archivo CSV
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Convertir los registros CSV en estructuras TestCase
	var testCases []models.TestCase
	for _, record := range records[1:] { // Ignorar la primera línea (cabeceras)
		testCases = append(testCases, models.TestCase{
			TestId:             record[0],
			TestCase:           record[1],
			Run:                record[2],
			Method:             record[3],
			URL:                record[4],
			Endpoint:           record[5],
			Authorization:      record[6],
			User:               record[7],
			Password:           record[8],
			Headers:            record[9],
			Body:               record[10],
			ExpectedStatusCode: atoi(record[11]),
			ExpectedResponse:   record[12],
		})
	}

	// Retornar el slice de casos de prueba y ningún error
	return testCases, nil
}

// atoi convierte un string a un número entero, manejando errores de conversión.
//
// Parameters:
//   - s (string): El string que se desea convertir a entero.
//
// Returns:
//   - int: El valor entero correspondiente al string proporcionado.
//   - Si la conversión falla, devuelve 0.
func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
