// Paquete api contiene funciones para realizar solicitudes HTTP y manejar sus respuestas.
package api

import (
	"bufio"
	"encoding/json"
	"io"
	"net/http"
)

// RealizarSolicitud realiza una solicitud HTTP utilizando el método, URL, y cuerpo especificados.
// Además, agrega cabeceras y credenciales de autenticación si se proporcionan.
// La función devuelve el código de estado HTTP de la respuesta, el cuerpo de la respuesta como cadena,
// y un posible error si la solicitud falla.
//
// Parameters:
//   - method (string): El método HTTP a usar para la solicitud (e.g., "GET", "POST").
//   - url (string): La URL a la que se debe hacer la solicitud.
//   - body (string): El cuerpo de la solicitud en formato JSON, si corresponde.
//   - headers (string): Un string JSON que representa las cabeceras adicionales a agregar a la solicitud.
//   - auth (string): El tipo de autenticación (por ejemplo, "Bearer" o "Basic").
//   - user (string): El nombre de usuario utilizado en la autenticación, si corresponde.
//   - password (string): La contraseña asociada al usuario, si corresponde.
//
// Returns:
//   - int: El código de estado HTTP de la respuesta (e.g., 200, 404).
//   - string: El cuerpo de la respuesta como una cadena de texto.
//   - error: Un error, si se produjo alguno durante la solicitud o procesamiento de la respuesta.
func RealizarSolicitud(method, url, body, headers, auth, user, password string) (int, string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return 0, "", err
	}

	// Agregar cabeceras si existen
	if headers != "" {
		var headerMap map[string]string
		if err := json.Unmarshal([]byte(headers), &headerMap); err != nil {
			return 0, "", err
		}
		for key, value := range headerMap {
			req.Header.Add(key, value)
		}
	}

	// Autorización: Si se proporciona un tipo de autenticación, agregar el encabezado de autorización
	if auth != "None" {
		req.Header.Add("Authorization", auth+" "+user+":"+password)
	}

	// Ejecutar la solicitud HTTP
	resp, err := client.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()

	// Leer el cuerpo de la respuesta
	var responseBody string
	buf := bufio.NewReader(resp.Body) // Usamos bufio para leer el cuerpo de manera más eficiente
	content := make([]byte, 0)

	for {
		chunk, err := buf.ReadByte() // Leemos byte a byte del cuerpo de la respuesta
		if err == io.EOF {
			break // Llegamos al final del cuerpo
		}
		if err != nil {
			return 0, "", err
		}
		content = append(content, chunk) // Añadimos cada byte al contenido
	}

	// Convertir el contenido leído en una cadena de texto
	responseBody = string(content)

	// Retornar el código de estado HTTP, el cuerpo de la respuesta y ningún error
	return resp.StatusCode, responseBody, nil
}
