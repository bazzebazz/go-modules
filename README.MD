Lista de comandos de GO para versionamiento

Si necesitamos importar un paquete que tenemos en los imports podemos hacer el comando

`go mod tidy` Este nos descargara la version compatible pero tambien limpiara los paquetes que no utilizemos

1. Para ver las versiones de un paquete es el siguiente comando `go list -m -versions rsc.io/sampler`
2. Para ver nuestros paquetes de dependencias en nuestro proyecto `go list -m all`
3. Para ver nuestras paquetes y sus dependencias `go mod graph`
4. Para actualizar un paquete a una version especifica `go get rsc.io/sampler@v1.3.1`
5. Para actualizar un paquete a la ultima version menor `go get golang.org/x/text`