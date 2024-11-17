// Paquete models define las estructuras utilizadas en las pruebas, como los casos de prueba en formato CSV.
package models

// TestCase representa un caso de prueba que contiene la información necesaria para ejecutar una prueba de API.
// La estructura se mapea directamente a un archivo CSV donde cada columna corresponde a un campo de la prueba.
//
// La estructura se utiliza para leer los casos de prueba desde un archivo CSV y luego ejecutar las solicitudes HTTP
// según los datos proporcionados.
//
// Campos:
//   - TestId: El identificador único del caso de prueba. (Por ejemplo, "TC01").
//   - TestCase: El nombre o descripción del caso de prueba.
//   - Run: Indica si el caso de prueba debe ejecutarse ("Y" para sí, "N" para no).
//   - Method: El método HTTP que se debe usar (por ejemplo, "GET", "POST").
//   - URL: La URL base a la que se le añadirá el endpoint.
//   - Endpoint: El endpoint específico a anexar a la URL base.
//   - Authorization: El tipo de autenticación que se utilizará para la solicitud (por ejemplo, "Bearer").
//   - User: El nombre de usuario para la autenticación, si se requiere.
//   - Password: La contraseña asociada al nombre de usuario, si se requiere.
//   - Headers: Las cabeceras HTTP adicionales para la solicitud, en formato JSON.
//   - Body: El cuerpo de la solicitud, que será enviado en el caso de solicitudes como "POST" o "PUT".
//   - ExpectedStatusCode: El código de estado HTTP que se espera recibir en la respuesta.
//   - ExpectedResponse: La respuesta esperada de la API, en formato JSON, que se compara con la respuesta real.
type TestCase struct {
	TestId             string `json:"TestId"`             // Identificador del caso de prueba.
	TestCase           string `json:"TestCase"`           // Nombre o descripción del caso de prueba.
	Run                string `json:"Run"`                // Indica si el caso de prueba debe ejecutarse ("Y" o "N").
	Method             string `json:"Method"`             // Método HTTP (GET, POST, PUT, DELETE).
	URL                string `json:"URL"`                // URL base para la solicitud.
	Endpoint           string `json:"Endpoint"`           // Endpoint que se añadirá a la URL base.
	Authorization      string `json:"Authorization"`      // Tipo de autorización (por ejemplo, "Bearer").
	User               string `json:"User"`               // Nombre de usuario para la autenticación.
	Password           string `json:"Password"`           // Contraseña para la autenticación.
	Headers            string `json:"Headers"`            // Cabeceras HTTP en formato JSON.
	Body               string `json:"Body"`               // Cuerpo de la solicitud (para POST, PUT).
	ExpectedStatusCode int    `json:"ExpectedStatusCode"` // Código de estado esperado en la respuesta.
	ExpectedResponse   string `json:"ExpectedResponse"`   // Respuesta esperada en formato JSON.
}
