// Paquete test contiene funciones relacionadas con la ejecución de pruebas y comparación de respuestas
// de una API frente a las respuestas esperadas definidas en los casos de prueba.
package test

import (
	"encoding/json"
	"fmt"
	"go-api-testing/internal/api"
	"go-api-testing/models"
	"reflect"
)

// EjecutarPrueba ejecuta una prueba basada en un caso de prueba especificado. La función realiza una solicitud HTTP
// utilizando los datos del caso de prueba, luego compara la respuesta obtenida con la respuesta esperada.
// Si el código de estado y el cuerpo de la respuesta coinciden con lo esperado, la prueba se considera exitosa.
//
// Parameters:
//   - test (models.TestCase): Un caso de prueba que contiene los detalles sobre el método HTTP,
//     la URL, las cabeceras, el cuerpo, la autenticación, y las respuestas esperadas.
//
// Returns:
//   - bool: Retorna `true` si la prueba fue exitosa, `false` en caso contrario.
//   - string: Un mensaje detallado sobre el resultado de la prueba, ya sea de éxito o el error encontrado.
func EjecutarPrueba(test models.TestCase) (bool, string) {
	// Construir la URL completa concatenando la base de la URL con el endpoint especificado
	fullURL := test.URL + test.Endpoint

	// Realizar la solicitud HTTP utilizando los parámetros proporcionados en el caso de prueba
	statusCode, response, err := api.RealizarSolicitud(test.Method, fullURL, test.Body, test.Headers, test.Authorization, test.User, test.Password)
	if err != nil {
		// Si ocurre un error en la solicitud, retornamos un mensaje con el error
		return false, fmt.Sprintf("Error en la solicitud: %v", err)
	}

	// Comprobar si el código de estado de la respuesta coincide con el esperado
	if statusCode != test.ExpectedStatusCode {
		return false, fmt.Sprintf("Código de estado esperado: %d, obtenido: %d", test.ExpectedStatusCode, statusCode)
	}

	// Si hay una respuesta esperada, comparamos la respuesta obtenida con la esperada
	if test.ExpectedResponse != "" {
		var expected, actual map[string]interface{}

		// Deserializar la respuesta esperada en una estructura de tipo mapa
		if err := json.Unmarshal([]byte(test.ExpectedResponse), &expected); err != nil {
			return false, fmt.Sprintf("Error al deserializar la respuesta esperada: %v", err)
		}

		// Deserializar la respuesta obtenida de la API en una estructura de tipo mapa
		if err := json.Unmarshal([]byte(response), &actual); err != nil {
			return false, fmt.Sprintf("Error al deserializar la respuesta obtenida: %v", err)
		}

		// Comparar las dos estructuras deserializadas para ver si son iguales
		if !reflect.DeepEqual(expected, actual) {
			// Si las respuestas no coinciden, las serializamos nuevamente como JSON para mostrar la diferencia
			expectedJSON, _ := json.Marshal(expected)
			actualJSON, _ := json.Marshal(actual)

			// Convertir las respuestas serializadas en cadenas JSON
			expectedJSONStr := string(expectedJSON)
			actualJSONStr := string(actualJSON)

			// Retornar el mensaje de error con las respuestas esperada y obtenida
			return false, fmt.Sprintf("Las respuestas no coinciden: Esperada: %s Obtenida: %s", expectedJSONStr, actualJSONStr)
		}
	}

	// Si el código de estado y la respuesta son correctos, la prueba es exitosa
	return true, "Prueba exitosa"
}
