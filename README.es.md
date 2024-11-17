<!-- omit from toc -->
# 🧪 **Framework de pruebas de API en Go**

[![Español](https://img.shields.io/badge/Language-Spanish-red)](README.es.md) [![English](https://img.shields.io/badge/Language-English-blue)](README.md)

Bienvenido! Escoge tu lenguaje preferido.

<!-- omit from toc -->
## 🌟 **Descripción del proyecto**

Este repositorio contiene un **marco de pruebas de API basado en Go** diseñado para automatizar las pruebas de APIs utilizando casos de prueba predefinidos almacenados en archivos CSV. Soporta la lectura y escritura de resultados en CSV, la generación de informes HTML y la ejecución de casos de prueba mediante solicitudes API configurables.

<!-- omit from toc -->
## 📑 Tabla de Contenidos

- [🧑‍💻 **Visión general**](#-visión-general)
- [🗂️ **Estructura del proyecto**](#️-estructura-del-proyecto)
- [🛠️ **Configuración e instalación**](#️-configuración-e-instalación)
- [🚀 **Uso**](#-uso)

---

## 🧑‍💻 **Visión general**

Este proyecto automatiza las pruebas de APIs mediante:

- 🔄 Carga de casos de prueba desde un archivo CSV.
- 📡 Realización de solicitudes HTTP basadas en los casos de prueba.
- ✅ Verificación del estado y contenido de la respuesta.
- 💾 Guardado de resultados en archivos CSV.
- 📊 Generación de informes HTML con los resultados de las pruebas.

---

## 🗂️ **Estructura del proyecto**

```bash
go-api-testing/
│
├── cmd/                        # Código principal de ejecución
│   └── main.go                 # Punto de entrada para ejecutar las pruebas
│
├── config/                     # Archivos de configuración
│   └── config.go               # Configuraciones globales (por ejemplo, lectura de CSV)
├── data/
│   └── results.csv             # Resultados de las pruebas en formato CSV
│   └── test_cases.csv          # Archivo CSV con los casos de prueba
│   └── report.html             # Informe HTML con los resultados de las pruebas
│
├── internal/                   # Lógica interna para la ejecución de pruebas
│   ├── api/                    # Lógica para hacer solicitudes API
│   │   └── client.go           # Funciones para enviar solicitudes HTTP
│   ├── csv/                    # Lógica para leer/escribir archivos CSV
│   │   └── reader.go           # Funciones para leer archivos CSV
│   │   └── writer.go           # Funciones para escribir resultados en CSV
│   ├── report/                 # Lógica para generar informes HTML
│   │   └── html.go             # Funciones para generar informes HTML
│   └── test/                   # Lógica de ejecución de pruebas
│       └── executor.go         # Funciones para ejecutar los casos de prueba
│
├── models/                     # Estructuras de datos para los casos de prueba
│   └── test_case.go            # Estructura para los datos del caso de prueba
│
├── scripts/                    # Scripts de configuración y utilidades
│   └── setup.bat               # Script de configuración para Windows
│   └── setup.sh                # Script de configuración para Linux/Mac
│
├── go.mod                      # Archivo del módulo Go con dependencias
├── go.sum                      # Suma de verificación de dependencias Go
└── README.md                   # Este archivo
```

---

## 🛠️ **Configuración e instalación**

Antes de ejecutar las pruebas, debes configurar el entorno necesario y las dependencias. Puedes usar el script `setup.sh` para Linux/Mac o el script `setup.bat` para Windows.

<!-- omit from toc -->
### 🌍 **Configuración en Linux/Mac (usando `setup.sh`)**

#### **Clonar el repositorio**

   ```bash
   git clone https://github.com/anxoportela/go-api-testing.git
   cd go-api-testing
   ```

#### **Hacer ejecutable el script `setup.sh`**

   ```bash
   chmod +x scripts/setup.sh
   ```

#### **Ejecutar el script de configuración**

   Este script instalará todas las dependencias necesarias, creará el archivo `.env` y generará el archivo `data/test_cases.csv`.

   ```bash
   ./scripts/setup.sh
   ```

   El script `setup.sh` hará lo siguiente:

- Instalar las dependencias necesarias de Go.
- Crear el archivo `.env` con las variables de entorno necesarias.
- Generar un archivo de ejemplo `test_cases.csv` en el directorio `data/`.

#### **Verificar la configuración**

   Una vez que el script de configuración haya terminado, deberías tener los siguientes archivos:

- `.env` (con las rutas para los archivos de casos de prueba, resultados e informes).
- `data/test_cases.csv` (casos de prueba predefinidos para ejecutar).
- Todas las dependencias necesarias de Go deberían estar instaladas.

---

<!-- omit from toc -->
### 🖥️ **Configuración en Windows (usando `setup.bat`)**

#### **Clonar el repositorio**

   ```cmd
   git clone https://github.com/anxoportela/go-api-testing.git
   cd go-api-testing
   ```

#### **Ejecutar el script `setup.bat`**

   Este script instalará todas las dependencias necesarias, creará el archivo `.env` y generará el archivo `data/test_cases.csv`.

   ```cmd
   scripts\setup.bat
   ```

   El script `setup.bat` hará lo siguiente:

- Instalar las dependencias necesarias de Go.
- Crear el archivo `.env` con las variables de entorno necesarias.
- Generar un archivo de ejemplo `test_cases.csv` en el directorio `data/`.

#### **Verificar la configuración**

   Después de que el script termine, confirma que se hayan creado los siguientes archivos:

- `.env` (con las rutas para los archivos de casos de prueba, resultados e informes).
- `data/test_cases.csv` (contiene los casos de prueba de ejemplo para las pruebas de API).

---

## 🚀 **Uso**

<!-- omit from toc -->
### **Preparar los casos de prueba**

   Edita el archivo `data/test_cases.csv` para incluir los casos de prueba que deseas ejecutar. Cada fila debe representar un caso de prueba con las siguientes columnas:

- `TestId`: Identificador único del caso de prueba.
- `TestCase`: Descripción del caso de prueba.
- `Run`: "Y" para ejecutar la prueba, "N" para saltarla.
- `Method`: Método HTTP (GET, POST, PUT, DELETE).
- `URL`: URL base para la API.
- `Endpoint`: El endpoint específico a probar.
- `Authorization`: Tipo de autorización (por ejemplo, token Bearer).
- `User`: Nombre de usuario para la autenticación (si es necesario).
- `Password`: Contraseña para la autenticación (si es necesario).
- `Headers`: Cabeceras en formato JSON (si las hay).
- `Body`: Cuerpo de la solicitud (para POST/PUT).
- `ExpectedStatusCode`: El código de estado HTTP esperado (por ejemplo, 200).
- `ExpectedResponse`: Cuerpo de la respuesta JSON esperado (si aplica).

<!-- omit from toc -->
### **Ejecutar las pruebas**

   Después de configurar el archivo `.env` y `test_cases.csv`, puedes ejecutar las pruebas ejecutando el archivo `main.go`.

   ```bash
   go run cmd/main.go
   ```

   Las pruebas se ejecutarán y los resultados se mostrarán en la terminal y se guardarán en:

- `data/results.csv` (archivo CSV con los resultados de las pruebas).
- `data/report.html` (informe HTML con los resultados de las pruebas).

---

<!-- omit from toc -->
## 📍 **Contribuciones**

¡Las contribuciones son bienvenidas! Si deseas contribuir a este proyecto, sigue estos pasos:

1. Haz un fork del proyecto.
2. Crea una nueva rama para tu característica o corrección de errores (`git checkout -b feature/nueva-característica`).
3. Realiza tus cambios y haz commit de ellos (`git commit -am 'Añadir nueva característica'`).
4. Empuja tus cambios a tu fork (`git push origin feature/nueva-característica`).
5. Crea una solicitud de pull.

---

<!-- omit from toc -->
## 📄 **Licencia**

Este proyecto está licenciado bajo la **Licencia MIT**. Para más detalles, consulta el archivo [LICENSE](LICENSE).

---

<!-- omit from toc -->
## 📧 **Contacto**

Si tienes problemas, preguntas o sugerencias, no dudes en ponerte en contacto con los mantenedores del proyecto:

**Correo electrónico**: [hello@anxoportela.dev](mailto:hello@anxoportela.dev)

---

<!-- omit from toc -->
### 🎉 **¡Disfruta utilizando el Framework de pruebas de API en Go!** 🎉
