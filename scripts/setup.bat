@echo off
REM --------------------------------------------------------
REM setup.bat: Script para preparar el entorno de pruebas API en Windows
REM --------------------------------------------------------

REM Descripción:
REM   Este script realiza la configuración inicial para ejecutar las pruebas API:
REM   - Verifica si Go está instalado.
REM   - Inicializa el módulo Go si no se ha hecho anteriormente.
REM   - Descarga las dependencias necesarias para el proyecto.
REM   - Crea un archivo CSV con casos de prueba básicos si no existe.
REM   - Elimina los archivos de resultados anteriores (data/results.csv y data/report.html) si existen.
REM   - Crea un archivo .env con las variables de entorno necesarias si no existe.
REM
REM Uso:
REM   Ejecuta el script desde la terminal con el siguiente comando:
REM   setup.bat
REM
REM Requisitos:
REM   - Go debe estar instalado en el sistema.
REM   - Debes tener permisos de escritura en el directorio donde se ejecuta el script.

REM --------------------------------------------------------

REM Imprimir mensaje de inicio
echo Iniciando la configuración del entorno de pruebas...

REM --------------------------------------------------------
REM 1. Asegurarse de que Go esté instalado
REM --------------------------------------------------------
where go >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo Error: Go no esta instalado. Por favor, instala Go.
    exit /b 1
) else (
    echo Go esta instalado: %go_version%
)

REM --------------------------------------------------------
REM 2. Inicializar el modulo Go (si no se ha hecho antes)
REM --------------------------------------------------------
if not exist go.mod (
    echo Inicializando el modulo Go...
    go mod init go-api-testing
) else (
    echo El modulo Go ya esta inicializado.
)

REM --------------------------------------------------------
REM 3. Descargar las dependencias necesarias
REM --------------------------------------------------------
echo Descargando dependencias de Go...
go mod tidy

REM --------------------------------------------------------
REM 4. Verificar y crear la carpeta 'data' si no existe
REM --------------------------------------------------------
if not exist "data" (
    echo La carpeta 'data' no existe. Creando...
    mkdir data
) else (
    echo La carpeta 'data' ya existe.
)

REM --------------------------------------------------------
REM 5. Crear el archivo CSV de casos de prueba (si no existe)
REM --------------------------------------------------------
set TEST_CASES_FILE=data\test_cases.csv
if not exist "%TEST_CASES_FILE%" (
    echo El archivo "%TEST_CASES_FILE%" no existe. Creando archivo con algunos casos de prueba básicos...

    REM Crear archivo CSV con ejemplos de pruebas
    echo TestId^,TestCase^,Run^,Method^,URL^,Endpoint^,Authorization^,User^,Password^,Headers^,Body^,ExpectedStatusCode^,ExpectedResponse > %TEST_CASES_FILE%
    echo TC-001^,Obtener primer usuario OK^,Y^,GET^,https://jsonplaceholder.typicode.com^,/users/1^,,,,,,200^,"{""id"":1,""name"":""Leanne Graham"",""username"":""Bret"",""email"":""Sincere@april.biz"}" >> %TEST_CASES_FILE%
    echo TC-002^,Obtener primer usuario KO^,Y^,GET^,https://jsonplaceholder.typicode.com^,/users/1^,,,,,,200^,"{}" >> %TEST_CASES_FILE%
    echo TC-003^,Obtener + StatusCode KO^,Y^,GET^,https://jsonplaceholder.typicode.com^,/users/1^,,,,,,404^,"{}" >> %TEST_CASES_FILE%
    echo TC-004^,Invalid endpoint^,Y^,GET^,https://jsonplaceholder.typicode.com^,/invalid-endpoint^,,,,,,404^,"{}" >> %TEST_CASES_FILE%
    echo TC-005^,Invalid Status Code^,Y^,GET^,https://jsonplaceholder.typicode.com^,/invalid-endpoint^,,,,,,200^,"{}" >> %TEST_CASES_FILE%
    echo TC-006^,Skipped test^,N^,GET^,https://jsonplaceholder.typicode.com^,/invalid-endpoint^,,,,,,404^,"{}" >> %TEST_CASES_FILE%
    echo TC-007^,Token Auth^,Y^,GET^,https://httpbin.org^,/bearer^,Bearer test^,,,,200^,"{""authenticated"":true,""token"":""test :""}" >> %TEST_CASES_FILE%
    echo TC-008^,Ejemplo de POST^,Y^,POST^,https://jsonplaceholder.typicode.com^,/users^,,,,,"{""name"":""Juan Pérez"",""email"":""juan@example.com"",""phone"":""123-456-7890""}"^,201^,"{""id"":11}" >> %TEST_CASES_FILE%
    echo TC-009^,Ejemplo de POST Error^,Y^,POST^,https://jsonplaceholder.typicode.com^,/users^,,,,,"{""name"":""Juan Pérez"",""email"":""juan@example.com"",""phone"":""123-456-7890""}"^,201^,"{}" >> %TEST_CASES_FILE%

    echo Archivo "%TEST_CASES_FILE%" creado con ejemplos de pruebas.
) else (
    echo El archivo "%TEST_CASES_FILE%" ya existe.
)

REM --------------------------------------------------------
REM 6. Verificar si los archivos de resultados existen, si es así, eliminarlos
REM --------------------------------------------------------
set RESULTS_FILE=data\results.csv
if exist "%RESULTS_FILE%" (
    echo El archivo "%RESULTS_FILE%" ya existe, eliminando archivo anterior...
    del "%RESULTS_FILE%"
)

set REPORT_FILE=data\report.html
if exist "%REPORT_FILE%" (
    echo El archivo "%REPORT_FILE%" ya existe, eliminando archivo anterior...
    del "%REPORT_FILE%"
)

REM --------------------------------------------------------
REM 7. Crear archivo .env si no existe
REM --------------------------------------------------------
set ENV_FILE=.env
if not exist "%ENV_FILE%" (
    echo El archivo "%ENV_FILE%" no existe. Creando archivo con las variables de entorno...

    REM Crear archivo .env con las variables de entorno
    echo TEST_CASES_FILE=data\test_cases.csv >> %ENV_FILE%
    echo RESULTS_FILE=data\results.csv >> %ENV_FILE%
    echo REPORT_FILE=data\report.html >> %ENV_FILE%

    echo Archivo "%ENV_FILE%" creado con las variables de entorno necesarias.
) else (
    echo El archivo "%ENV_FILE%" ya existe.
)

REM --------------------------------------------------------
REM 8. Mostrar mensaje final
REM --------------------------------------------------------
echo El entorno está listo. Puedes ejecutar las pruebas ahora con 'go run cmd\main.go'.
pause
