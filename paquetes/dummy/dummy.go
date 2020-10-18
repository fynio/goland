package dummy

var variable string

//HolaMundo aqui va la descripcion de esta funcion se me hace muy buena practica
func HolaMundo() string {
	return " " + variable
	// hola mundo
}

func minusculas() string {
	return "no se pueden ejecutar desde un paquete externo"
}

func init() {
	variable = "hola mundo =)"
}
