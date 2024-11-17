#!/bin/bash

# --------------------------------------------------------
# setup.sh: Script para preparar el entorno de pruebas API
# --------------------------------------------------------

# Descripción:
#   Este script realiza la configuración inicial para ejecutar las pruebas API:
#   - Verifica si Go está instalado.
#   - Inicializa el módulo Go si no se ha hecho anteriormente.
#   - Descarga las dependencias necesarias para el proyecto.
#   - Crea un archivo CSV con casos de prueba básicos si no existe.
#   - Elimina los archivos de resultados anteriores (data/results.csv y data/report.html) si existen.
#   - Crea un archivo .env con las variables de entorno necesarias si no existe.
#
# Uso:
#   Ejecuta el script desde la terminal con el siguiente comando:
#   ./setup.sh
#
# Requisitos:
#   - Go debe estar instalado en el sistema.
#   - Debes tener permisos de escritura en el directorio donde se ejecuta el script.

# --------------------------------------------------------

# Imprimir mensaje de inicio
echo "Iniciando la configuración del entorno de pruebas..."

# --------------------------------------------------------
# 1. Asegurarse de que Go esté instalado
# --------------------------------------------------------
if ! command -v go &> /dev/null
then
    echo "Error: Go no está instalado. Por favor, instala Go."
    exit 1
else
    echo "Go está instalado: $(go version)"
fi

# --------------------------------------------------------
# 2. Inicializar el módulo Go (si no se ha hecho antes)
# --------------------------------------------------------
if [ ! -f go.mod ]; then
    echo "Inicializando el módulo Go..."
    go mod init go-api-testing
else
    echo "El módulo Go ya está inicializado."
fi

# --------------------------------------------------------
# 3. Descargar las dependencias necesarias
# --------------------------------------------------------
echo "Descargando dependencias de Go..."
go mod tidy

# --------------------------------------------------------
# 4. Verificar y crear la carpeta 'data' si no existe
# --------------------------------------------------------
# Verificamos si la carpeta 'data' existe. Si no, la creamos.
if [ ! -d "data" ]; then
    echo "La carpeta 'data' no existe. Creándola..."
    mkdir -p data
else
    echo "La carpeta 'data' ya existe."
fi

# --------------------------------------------------------
# 5. Crear el archivo CSV de casos de prueba (si no existe)
# --------------------------------------------------------
TEST_CASES_FILE="data/test_cases.csv"
if [ ! -f "$TEST_CASES_FILE" ]; then
    echo "El archivo '$TEST_CASES_FILE' no existe. Creando archivo con algunos casos de prueba básicos..."
    
    # Crear archivo CSV con ejemplos de pruebas
    cat <<EOL > $TEST_CASES_FILE
TestId,TestCase,Run,Method,URL,Endpoint,Authorization,User,Password,Headers,Body,ExpectedStatusCode,ExpectedResponse
TC-001,Obtener primer usuario OK,Y,GET,https://jsonplaceholder.typicode.com,/users/1,,,,,,200,"{""id"":1,""name"":""Leanne Graham"",""username"":""Bret"",""email"":""Sincere@april.biz"",""address"":{""street"":""Kulas Light"",""suite"":""Apt. 556"",""city"":""Gwenborough"",""zipcode"":""92998-3874"",""geo"":{""lat"":""-37.3159"",""lng"":""81.1496""}},""phone"":""1-770-736-8031 x56442"",""website"":""hildegard.org"",""company"":{""name"":""Romaguera-Crona"",""catchPhrase"":""Multi-layered client-server neural-net"",""bs"":""harness real-time e-markets""}}"
TC-002,Obtener primer usuario KO,Y,GET,https://jsonplaceholder.typicode.com,/users/1,,,,,,200,"{}"
TC-003,Obtener + StatusCode KO,Y,GET,https://jsonplaceholder.typicode.com,/users/1,,,,,,404,"{}"
TC-004,Invalid endpoint,Y,GET,https://jsonplaceholder.typicode.com,/invalid-endpoint,,,,,,404,"{}"
TC-005,Invalid Status Code,Y,GET,https://jsonplaceholder.typicode.com,/invalid-endpoint,,,,,,200,"{}"
TC-006,Skipped test,N,GET,https://jsonplaceholder.typicode.com,/invalid-endpoint,,,,,,404,"{}"
TC-007,Token Auth,Y,GET,https://httpbin.org,/bearer,Bearer test,,,,,200,"{""authenticated"":true,""token"":""test :""}"
TC-008,Ejemplo de POST,Y,POST,https://jsonplaceholder.typicode.com,/users,,,,,"{""name"":""Juan Pérez"",""email"":""juan@example.com"",""phone"":""123-456-7890""}",201,"{""id"":11}"
TC-009,Ejemplo de POST Error,Y,POST,https://jsonplaceholder.typicode.com,/users,,,,,"{""name"":""Juan Pérez"",""email"":""juan@example.com"",""phone"":""123-456-7890""}",201,"{}"
EOL

    echo "Archivo '$TEST_CASES_FILE' creado con ejemplos de pruebas."
else
    echo "El archivo '$TEST_CASES_FILE' ya existe."
fi

# --------------------------------------------------------
# 6. Verificar si los archivos de resultados existen, si es así, eliminarlos
# --------------------------------------------------------
RESULTS_FILE="data/results.csv"
if [ -f "$RESULTS_FILE" ]; then
    echo "El archivo '$RESULTS_FILE' ya existe, eliminando archivo anterior..."
    rm "$RESULTS_FILE"
fi

REPORT_FILE="data/report.html"
if [ -f "$REPORT_FILE" ]; then
    echo "El archivo '$REPORT_FILE' ya existe, eliminando archivo anterior..."
    rm "$REPORT_FILE"
fi

# --------------------------------------------------------
# 7. Crear archivo .env si no existe
# --------------------------------------------------------
ENV_FILE=".env"
if [ ! -f "$ENV_FILE" ]; then
    echo "El archivo '$ENV_FILE' no existe. Creando archivo con las variables de entorno..."
    
    # Crear archivo .env con las variables de entorno
    cat <<EOL > $ENV_FILE
TEST_CASES_FILE=data/test_cases.csv
RESULTS_FILE=data/results.csv
REPORT_FILE=data/report.html
EOL

    echo "Archivo '$ENV_FILE' creado con las variables de entorno necesarias."
else
    echo "El archivo '$ENV_FILE' ya existe."
fi

# --------------------------------------------------------
# 8. Mostrar mensaje final
# --------------------------------------------------------
echo "El entorno está listo. Puedes ejecutar las pruebas ahora con 'go run cmd/main.go'."
