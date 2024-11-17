// Paquete main contiene la función principal que orquesta la carga de configuración,
// la ejecución de pruebas, la visualización de resultados y la generación de reportes.
package main

import (
	"fmt"
	"go-api-testing/config"
	"go-api-testing/internal/csv"
	"go-api-testing/internal/report"
	"go-api-testing/internal/test"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

// main es la función principal que ejecuta todo el proceso de pruebas automatizadas.
// Esto incluye cargar la configuración desde las variables de entorno, leer los casos de prueba,
// ejecutar las pruebas, mostrar los resultados en la consola y generar un reporte en HTML y un archivo CSV.
func main() {
	// Cargar la configuración desde las variables de entorno
	config.LoadConfig()

	// Leer los casos de prueba desde el archivo CSV especificado en la configuración
	testCases, err := csv.LeerCSV(config.AppConfig.TestCasesFile)
	if err != nil {
		log.Fatalf("Error al leer el archivo CSV: %v", err)
	}

	// Crear una tabla para mostrar los resultados de las pruebas en la consola
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"TestId", "TestCase", "Resultado", "Mensaje"})

	// Crear una lista para almacenar los resultados que se escribirán en un archivo CSV
	var results [][]string
	results = append(results, []string{"TestId", "TestCase", "Resultado", "Mensaje"})

	// Iterar sobre cada caso de prueba y ejecutarlo si tiene la etiqueta 'Y' para ser ejecutado
	for _, testCase := range testCases {
		if testCase.Run == "Y" {
			// Ejecutar la prueba correspondiente utilizando la función EjecutarPrueba
			success, message := test.EjecutarPrueba(testCase)

			// Determinar el mensaje que se mostrará en consola dependiendo del resultado de la prueba
			var consoleMessage string
			if success {
				consoleMessage = "Prueba exitosa"
			} else {
				consoleMessage = "Prueba fallida"
			}

			// Agregar el resultado de la prueba a la tabla que se mostrará en consola
			result := []string{testCase.TestId, testCase.TestCase, fmt.Sprintf("%t", success), consoleMessage}
			table.Append(result)

			// Agregar los resultados al slice de resultados que se escribirá en el archivo CSV
			results = append(results, []string{testCase.TestId, testCase.TestCase, fmt.Sprintf("%t", success), message})
		}
	}

	// Imprimir la tabla con los resultados de las pruebas en la consola
	table.Render()

	// Escribir los resultados de las pruebas en un archivo CSV especificado en la configuración
	err = csv.EscribirResultados(results, config.AppConfig.ResultsFile)
	if err != nil {
		log.Fatalf("Error al escribir el archivo de resultados: %v", err)
	}

	// Generar un reporte HTML con los resultados de las pruebas
	err = report.GenerarReporteHTML(results, config.AppConfig.ReportFile)
	if err != nil {
		log.Fatalf("Error al generar el reporte HTML: %v", err)
	}

	// Imprimir mensaje de éxito en la consola una vez que las pruebas se hayan ejecutado correctamente
	fmt.Println("Las pruebas se ejecutaron correctamente y los resultados se guardaron en 'results.csv'")
}
