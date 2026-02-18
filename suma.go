package proyectoModuloGit

import (
	"math/rand"
)

// Sumar suma dos números
func Sumar(a, b int) int {
	return a + b
}

// Saludar retorna un saludo personalizado
func Saludar(nombre string) string {
	return "Hola, " + nombre + "!"
}

func SaludarRandom() string {
	randomValues := []string{"Hola", "Hi", "Hey", "Saludos", "Greatings"}
	return randomValues[rand.Intn(len(randomValues))]
}
