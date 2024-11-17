// Paquete config maneja la carga de la configuración de la aplicación desde
// variables de entorno, especialmente los archivos de casos de prueba, resultados y reporte.
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config representa la estructura que contiene las rutas de los archivos utilizados por la aplicación.
// TestCasesFile: Ruta del archivo CSV que contiene los casos de prueba.
// ResultsFile: Ruta del archivo CSV donde se guardan los resultados de las pruebas.
// ReportFile: Ruta del archivo HTML donde se generará el reporte de las pruebas.
type Config struct {
	TestCasesFile string // Ruta del archivo CSV de casos de prueba
	ResultsFile   string // Ruta del archivo CSV de resultados
	ReportFile    string // Ruta del archivo HTML de reporte
}

// AppConfig es una instancia global de la configuración de la aplicación.
var AppConfig Config

// LoadConfig carga la configuración desde las variables de entorno.
// Esta función carga las rutas de los archivos desde un archivo .env usando la librería godotenv.
// Si alguna de las variables de entorno no está configurada, el programa termina con un error.
func LoadConfig() {
	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	// Obtener la ruta del archivo de casos de prueba desde la variable de entorno TEST_CASES_FILE
	testCasesFile := os.Getenv("TEST_CASES_FILE")
	if testCasesFile == "" {
		log.Fatal("TEST_CASES_FILE no está configurado. Establezca la variable en el archivo .env.")
	}

	// Obtener la ruta del archivo de resultados desde la variable de entorno RESULTS_FILE
	resultsFile := os.Getenv("RESULTS_FILE")
	if resultsFile == "" {
		log.Fatal("RESULTS_FILE no está configurado. Establezca la variable en el archivo .env.")
	}

	// Obtener la ruta del archivo de reporte HTML desde la variable de entorno REPORT_FILE
	reportFile := os.Getenv("REPORT_FILE")
	if reportFile == "" {
		log.Fatal("REPORT_FILE no está configurado. Establezca la variable en el archivo .env.")
	}

	// Asignar las rutas de los archivos a la estructura AppConfig
	AppConfig = Config{
		TestCasesFile: testCasesFile,
		ResultsFile:   resultsFile,
		ReportFile:    reportFile,
	}

	// Confirmar que la configuración se cargó correctamente mostrando las rutas de los archivos.
	log.Printf("Configuración cargada con éxito. TestCasesFile: %s, ResultsFile: %s, ReportFile: %s", AppConfig.TestCasesFile, AppConfig.ResultsFile, AppConfig.ReportFile)
}
