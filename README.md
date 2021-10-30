SEMINARIO GO 2021
TUDAI RAUCH

Trabajo realizado por Esains Sebastian y Esains Leonardo

El trabajo Practico consiste en dada una cadena decir si es valida o no
ej: la cadena "TX03SYL" es valida ya que el tipo TX es texto, 03 es la longitud de SYL
ej: la cadena "NN03SYL" no es valida ya que el tipo NN es numerico y SYL es texto

Se realizo toda la estructura necesaria para retornar si la cadena es valida o no. 
La estructura de la cadena a analizar es la siguiente: los primeros dos caracteres corresponden al tipo (NN para numeros y TX para texto), los dos caracteres siguientes
corresponden a la longitud de la cadena y lo que sigue para la derecha es el valor.
Para comenzar a comprobar si una cadena es valida o no nos tenemos que asegurar que dicha cadena sea mayor de 4, si no es asi retorna un mensaje de error.
Si la cadena es mayor a 4 se la separa por partes, longitud, tipo y valor para realizar los chequeos correspondientes.
Se generan los chequeos correspondientes para validar o no la cadena mediante main.test.go y entities.test.go
