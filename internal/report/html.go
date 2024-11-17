// Paquete report contiene funciones para generar reportes en formato HTML a partir de los resultados de las pruebas.
package report

import (
	"fmt"
	"os"
	"text/template"
)

// ReportData estructura que contiene los resultados de las pruebas a ser utilizados en la generación de un reporte HTML.
type ReportData struct {
	Results [][]string // Los resultados de las pruebas, cada fila contiene datos de un caso de prueba.
}

// GenerarReporteHTML recibe los resultados de las pruebas y genera un reporte en formato HTML.
// La función toma los resultados, los filtra (ignorando la primera fila de cabeceras), y genera un archivo HTML con
// una tabla que contiene la información de los casos de prueba.
//
// Parameters:
//   - results ([][]string): Un slice bidimensional de strings que contiene los resultados de las pruebas.
//   - filePath (string): La ruta donde se guardará el archivo HTML generado.
//
// Returns:
//   - error: Un error en caso de que ocurra algún problema al generar el reporte HTML o crear el archivo.
func GenerarReporteHTML(results [][]string, filePath string) error {
	// Filtrar las cabeceras, asumiendo que la primera fila es la cabecera
	if len(results) > 0 {
		results = results[1:] // Ignorar la primera fila (cabeceras)
	}

	// Preparar los datos del reporte
	reportData := ReportData{Results: results}

	// Crear el archivo HTML
	reportFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error al crear el archivo HTML: %v", err)
	}
	defer reportFile.Close()

	// Template HTML que define la estructura de la tabla
	tmpl := `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Reporte de Resultados de Pruebas</title>
				<style>
					/* Estilos del reporte HTML */
					body {
						font-family: 'Roboto', Arial, sans-serif;
						background-color: #f4f4f9;
						margin: 0;
						padding: 20px;
						color: #333;
					}
					h2 {
						text-align: center;
						color: #444;
						margin-bottom: 20px;
					}
					table {
						width: 100%;
						border-collapse: collapse;
						background: white;
						margin: 20px auto;
						box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
						border-radius: 8px;
						overflow: hidden;
					}
					th, td {
						padding: 12px 15px;
						text-align: left;
						border: 1px solid #ddd;
					}
					th {
						background-color: #007bff;
						color: white;
						text-transform: uppercase;
						font-size: 14px;
					}
					td {
						font-size: 14px;
						color: #555;
						vertical-align: top;
					}
					tr:nth-child(even) {
						background-color: #f9f9f9;
					}
					tr:hover {
						background-color: #f1f7ff;
					}
					.status-pass {
						color: #28a745;
						font-weight: bold;
					}
					.status-fail {
						color: #dc3545;
						font-weight: bold;
					}
					.emoji {
						font-size: 18px;
						margin-right: 8px;
					}
					.message-cell {
						max-width: 300px; /* Limitar el ancho máximo */
						overflow: hidden;
						white-space: nowrap; /* Evitar salto de línea por defecto */
						text-overflow: ellipsis; /* Mostrar "..." cuando el texto es demasiado largo */
						word-wrap: break-word;
					}
					.message-full {
						white-space: normal; /* Expandir texto completo */
					}
					.expand-btn {
						color: #007bff;
						cursor: pointer;
						text-decoration: underline;
						font-size: 12px;
					}
				</style>
				<script>
					function toggleMessage(id) {
						const cell = document.getElementById(id);
						const isExpanded = cell.classList.contains('message-full');
						cell.classList.toggle('message-full', !isExpanded);
						const btn = document.getElementById('btn-' + id);
						if (isExpanded) {
							btn.innerText = 'Ver más';
						} else {
							btn.innerText = 'Ver menos';
						}
					}
				</script>
			</head>
			<body>
				<h2>📋 Reporte de Resultados de Pruebas</h2>
				<table>
					<tr>
						<th><span class="emoji">🆔</span>TestId</th>
						<th><span class="emoji">📝</span>TestCase</th>
						<th><span class="emoji">✅</span>Resultado</th>
						<th><span class="emoji">📜</span>Mensaje</th>
					</tr>
					{{range $index, $row := .Results}} <!-- Iterar sobre los resultados -->
					<tr>
						<td>{{index $row 0}}</td> <!-- Mostrar TestId -->
						<td>{{index $row 1}}</td> <!-- Mostrar TestCase -->
						<td class="{{if eq (index $row 2) "true"}}status-pass{{else}}status-fail{{end}}">
							{{if eq (index $row 2) "true"}}Passed{{else}}Failed{{end}}
						</td> <!-- Mostrar Resultado con texto y estilo -->
						<td>
							<div id="message-{{$index}}" class="message-cell">
								{{index $row 3}}
							</div>
							{{if gt (len (index $row 3)) 15}} <!-- Mostrar botón si el mensaje es largo -->
							<span class="expand-btn" id="btn-message-{{$index}}" onclick="toggleMessage('message-{{$index}}')">Ver más</span>
							{{end}}
						</td> <!-- Mostrar Mensaje -->
					</tr>
					{{end}}
				</table>
			</body>
			</html>
			`

	// Parsear la plantilla y ejecutarla para generar el reporte
	t, err := template.New("report").Parse(tmpl)
	if err != nil {
		return fmt.Errorf("error al parsear la plantilla HTML: %v", err)
	}

	// Generar el archivo HTML con los resultados
	err = t.Execute(reportFile, reportData)
	if err != nil {
		return fmt.Errorf("error al generar el reporte HTML: %v", err)
	}

	return nil
}
