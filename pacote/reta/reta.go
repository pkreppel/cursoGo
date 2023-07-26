package main

import "math"

//Inicializando com letra maíuscula, é público dentro e fora do pacote.
//Minúsculo é privado dentro do package

//Ponto é publico
//ponto ou _ponto é privado

//Ponto representa uma coordenada no pano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

//Distancia é responsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
