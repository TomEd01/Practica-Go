# 🧱 Fundamentos en Go

Esta carpeta contiene los fundamentos básicos del lenguaje de programación Go, cada uno implementado en su propia subcarpeta para facilitar la ejecución individual y evitar conflictos de compilación.

---

## 📚 Contenidos

| Tema            | Descripción                                          |
|-----------------|------------------------------------------------------|
| `hola/`         | Primer programa en Go: impresión básica por consola. |
| `datos/`        | Declaración de variables, constantes y tipos.        |
| `entrada_salida/` | Lectura de datos desde consola.                     |
| `condicionales/`| Uso de estructuras `if`, `else if`, `else`.         |
| `operadores/`   | Operadores aritméticos, lógicos y de comparación.    |

---

## ⚠️ Nota sobre la estructura

Cada tema se encuentra en una **subcarpeta separada** porque todos los archivos contienen una función `main()`, y **Go no permite múltiples funciones `main()` en el mismo paquete al compilar en conjunto**.

Por ello, para evitar errores como `main redeclared in this block`, se optó por esta organización modular. Esto permite ejecutar cada ejemplo de forma independiente, lo que además promueve buenas prácticas de desarrollo.

```bash
# Ejemplo de ejecución:
go run Fundamentos/condicionales
