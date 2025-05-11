# Tabla de Contenidos

- [Tabla de Contenidos](#tabla-de-contenidos)
- [GO](#go)
  - [Recursos](#recursos)
  - [Instalación](#instalación)
  - [Usos Principales](#usos-principales)
  - [Frameworks](#frameworks)
  - [Librería Estandar](#librería-estandar)
  - [Paquetes](#paquetes)
    - [Paquete `main`](#paquete-main)
    - [Importar Paquetes](#importar-paquetes)
  - [Módulos](#módulos)
  - [Punto de Entrada de un Módulo](#punto-de-entrada-de-un-módulo)
  - [Tipos de Datos Primitivos](#tipos-de-datos-primitivos)
  - [Variables](#variables)
  - [Condicionales](#condicionales)
    - [if, else if y else](#if-else-if-y-else)
    - [Switch](#switch)
  - [Arrays](#arrays)
  - [Maps](#maps)
  - [Bucles](#bucles)

# GO

GO es un lenguaje de prograación creado por Google en 2007 pensado para ser eficiente y fácil de aprender. Surge como una alternativa a C/C++ para ciertas tareas, sobre todo las relacionadas a la concurrencia.

GO es un lenguaje compilado, fuertemente tipado, concurrente, imperativo ([programación imperativa](https://es.wikipedia.org/wiki/Programaci%C3%B3n_imperativa)) y estructurado (no posee Clases, Objetos ni Enumerados, en su lugar si queremos representar "objetos" tendremos estructuras)

## Recursos

- [Mouredev - Primeros pasos en GO](https://www.youtube.com/watch?v=AGiayASyp2Q)

## Instalación

Instalar GO es muy sencillo (y similar a Python), debemos dirigirnos a la página oficial y allí descargar el instalador de la última versión estable de GO ([aquí](https://go.dev/doc/install)), tras tener el instalador lo ejecutamos y seguimos los pasos estandar de instalar un programa.

## Usos Principales

GO es increíble manejando concurrencia y trabajando con múltiples hilos por lo que algunos de los usos principales de GO son:

- Aplicaciones de servidor, APIs, BBDD, Microservicios.
- Aplicaciones de CLI, tanto simples como complejas.
- Análisis de datos.

## Frameworks

Hay muchos Frameworks para GO, principalmente orientados a la creación de aplicaciones de servidor. Algunos de los más conocidos son Gin, Fiber y Buffalo, pero hay muchos más.

## Librería Estandar

Hace referencia a todos los módulos y paquetes que trae GO por defecto (lo que en Java llamamos la API de Java). La documentación de esta se encuentra [aquí](https://pkg.go.dev/std)

## Paquetes

Un paquete es un conjunto de archivos que se encuentran en el mismo directorio y son compilados juntos, todas las funciones, contantes, tipos y variables creadas en un archivo son visibles para todo el resto del mismo paquete.

Para definir un paquete en GO es tan simple como que en nuestro archivo `.go` escribamos `package <packageName>`, hacer esto es obligatorio ya que todos los archivos de GO deben pertenecer a un paquete.

### Paquete `main`

Cuando nosotros queramos ejecutar nuestro código de GO por defecto intentará buscar el paquete llamado `main`, por lo que cuando queramos que nuestro código de GO sea ejecutado directamente tener un paquete llamado así, no obstante cuando publiquemos un paquete con la idea de que otras personas lo descarguen e implementen en sus propios proyectos NO debemos tener un paquete `main`.

```go
package main

// ...
```

### Importar Paquetes

GO está pensado para solo cargar el código estrictamente necesario, en consecuencia vamos a tener que importar paquetes constantemente, incluso los de la Librería Estándar, hacer esto es bastante simple, usamos la palabra reservada `import` seguida del nombre del paquete entre comillas dobles (`"`).

Uno de los paquetes más básicos de GO es `fmt` que permite dar formato a textos

```go
package main

import "fmt"

// ...
```

## Módulos

Un módulo es en esencia un conjunto de paquetes que se distribuyen en conjunto.

GO está basado en módulos, un programa de GO (sea el que sea) es en sí mismo un módulo que es capaz de importar código de otros módulos. Esto es importante porque, como vimos con anterioridad, en GO vamos a usar constantemente código de otros módulos, como la Librería Estándar, y si no inicializamos nuestro módulo no prodremos hacerlo.

Para inicializar nuetro modulo ejecutamos lo siguiente en nuestra terminal

```sh
mkdir myProject

cd myProject

go mod init example/myProject # Esto inicializa el módulo
```

A la hora de inicializar el módulo de GO tenemos que darle un nombre, generalmente este es la url del repositorio en donde se puse encocntrar el módulo (Ej: `github.com/account/styleConsole`), también puede ocurrir que el módulo en cuestión sea submódulo de otro, en dicho caso debemos especificar el subdirectorio, por ejemplo si en el repositorio el "sub-módulo" está en la carpeta `src_code` el nombre de este "sub-módulo" debería ser `github.com/account/styleConsole/src_code`.

En general el nombre del módulo debe describir dónde encontrarlo y qué hace.

Al ejecutar `go mod init <modulePath>` se creará un archivo `go.mod` que se encargará de gestionar todo lo referente al módulo, entre esto las dependencias.

## Punto de Entrada de un Módulo

Al ejecutar un módulo GO busca el paquete `main` y dentro de él ejecuta la función con el mismo nombre, que es el punto de entrada de la aplicación. Veremos funciones en GO más adelante pero de momento nos basta con entender que para poder ejecutar nuestro código es obligatorio tener por lo menos una.

```go
// path: ./hello_world.go
package main

import "fmt" // da formato a textos

func main() {
  fmt.printLn("Hola mundo") // `fmt.Println()` imprime texto en consola con un salto de linea
}
```

## Tipos de Datos Primitivos

- `string` --> cadenas de texto, se escriben siempre entre comillas dobles (`"`)
- `int` --> para enteros de, como mínimo, 32 bits, toma de referencia la arquitectura del OS
- `int8` --> para enteros de 8 bits
- `int16` --> para enteros de 16 bits
- `int32` --> para enteros de 32 bits
- `int64` --> para enteros de 64 bits
- `uint` --> para enteros sin signo de, como mínimo, 32 bits, también existen sus versiones de 8, 16, 32 y 64.
- `float32` --> para decimales de 32 bits
- `float64` --> para decimales de 64 bits
- `bool` --> `true` o `false`

## Variables

La forma más simple de definir variables en GO es con la palabra reservada `var` siguiendo la estructura `var <NAME> <TYPE>`, al igual que en leguajes como JavaScript podemos declarar una variable e inicializarla en otra linea o hacerlo en una sola.

```go
package main

import "fmt"

func main() {
  var myString string
  myString = "Soy una cadena"

  var otherString string = "Soy otra cadena"

  // myString = 10 --> Error pq GO es de tipado fuerte

  fmt.Println(myString, "-", otherString)
}
```

No obstante, GO tiene algo que pocos leguajes tipados poseen, y es inferencia de tipos, esto quiere decir que en ciertos casos GO es capaz de reconocer a qué tipo de dato pertenece una variable sin que se lo digamos.

```go
package main

import "fmt"

func main() {
  var myString = "Soy una cadena" // GO infiere que "myString" es de tipo "string"
}
```

Si queremos ahorrarnos usar la palabra reservada `var` a la hora de declarar e inicializar una variable al mismo tiempo usando la inferencia de tipos podemos usar el operador `:=`

```go
package main

import "fmt"

func main() {
  myString := "Soy una cadena"
}
```

## Condicionales

### if, else if y else

Como en muchos leguajes de programación tenemos condicionales que siguen la siguiente estructura

```txt
  // ...
  if <condicion> {
    // code
  } else if <condicion> {
    // code
  } else {
    // code
  }
```

```go
package main

import "fmt"

func main() {
  var myInt int16 = 10

  if myInt == 0 {
    fmt.Println("myInt es exactamente 0")
  } else if myInt < 0 {
    fmt.Println("myInt es negativo")
  } else {
    fmt.Println("myInt es positivo")
  }
}
```

### Switch

> [!WARNING]
> IN PROGRESS...

## Arrays

Al igual que en lenguajes como Java a la hora de crear un Array debemos aclarar su longitud y tipo de dato, esto lo hacemos con la sintaxis `var <arrayName> [<elementsQty>]<type>`. Ej:

```go
var myArray [3]int // un Array de 3 enteros
```

A la hora de asignarle valores a los elementos de un Array lo hacemos con `myArray[<index>] = <value>`.

```go
var myArray [3]int8
myArray[0] = 10
myArray[1] = 30
myArray[2] = 8

fmt.Println(myArray)
```

## Maps

Los Maps son estructuras de datos del tipo `clave-valor` como los diccionarios de Python o los Objetos de JavaScript. Para crearlos usamos la función `make`, y le pasamos como argumento `map[<keyType>]<valueType>`. Por ejemplo, si queremos crear un Map en el que la clave sea un nombre y el valor sea la edad de una persona lo hariamos de la siguiente manera

```go
ages := make(map[string]unit8)
```

Para asignar valores al Map lo que hacemos es pasarle la clave entre corchetes e igualarlo al valor

```go
ages["Paco"] = 24
ages["Carlos"] = 36
ages["Juan"] = 64
```

A la hora de mostrar simplemente accedemos a la clave especificandola entre corchetes.

```go
fmt.Println(ages) // imprime todo el Map
fmt.Println(ages["Paco"]) // Imprime el valor de la clave "Paco"
```

Si queremos inicializar el Map al mismo tiempo que lo creamos podemos hacerlo de la siguiente manera

```go
sonsQty := map[string]uint8{"Paco": 0, "Carlos": 1, "Juan": 4}
fmt.Println(sonsQty)
```

## Bucles

> [!WARNING]
> IN PROGRESS...
