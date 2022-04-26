# Drivers Challenge

# Objetivo
Realizar una rest API en go que ofrezca los siguientes servicios:
- Guardado de nuevos conductores
- Autenticación de conductores a través de token
- Consulta de conductores por paginado y cantidad configurable (requiere autenticación)
- Consulta de conductores disponibles (aquellos que no se encuentran haciendo un viaje)

# Detalle de punto de entrada
El punto de entrada se encuentra en el archivo root.go. Dicho archivo contiene la inicialización de una base de datos tipo SQL local. Además crea las instancias de los componentes necesarios de todas las capas de arquitectura de capas.

## Inyección de dependencias
Todas las dependencias se encuentran dentro del paquete di. Ahí se componen los objetos respetando las interfaces señaladas en cada carpeta de cada capa.

## Autenticación y autorización
Se utilizó la librería JWT correspondiente al middleware gin-gonic para tratar este tema. Dentro del archivo middlewares/auth.go está el detalle de la autorización del token. Esta autorización se utiliza para los endpoints detallados anteriormente, en el archivo root.go.
Además, la lógica de generación y validación de tokens se están en el archivo auth/auth.go.

# Presentación
La capa de presentación corresponde a los siguiente paquetes:
- Controllers
- Bodies

## Controllers
En dicho paquete se encuentran los controladores para generar token y para realizar las operaciones con respecto a conductores. Dentro de la carpeta se encuentran las interfaces que usan los controladores por lo que los mismos no conocen ningún detalle de implementación de los servicios.
Todos los controladores llaman a los servicios y devuelven el resultado junto a su código HTTP. En caso de error, los controladores toman el error y cambian el código de respuesta.

## Bodies
En este paquete solo se encuentran las estructuras que aceptarán los controladores para las operaciones de tipo POST.

# Servicio
La capa de servicio solo está representada por el paquete services. 
Dentro del mismo se encuentran los servicios la recepción por parte del controlador y el traspaso. En esta capa estarían las validaciones de lógica de negocio, que en este caso no hay. 
Dentro de esta carpeta también se encuentran las interfaces que utilizan los servicios, por lo que no conocen los detalles de implementación de los repositorios. Cabe destacar que el paquete models es conocida tanto por la capa de servicio como la de persistencia.

# Persistencia
La capa de persistencia está compuesta por:
- Repositories
- Config
- Migrations

## Repositories
En este paquete se encuentran las clases que reciben un modelo y ejecutan la query correspondiente en la base de datos.

## Config
En este paquete solamente se encuentra la configuración de la base de datos.

## Migrations
En este paquete se encuentran todas las querys que se ejecutan cuando se inicia la aplicación.
