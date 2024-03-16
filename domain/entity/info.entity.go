package entity

import "github.com/google/uuid"

type Cabelo int

const (
	Crespo Cabelo = iota
	Liso
	Cacheado
)

func (c Cabelo) String() string {
	return [...]string{"Crespo", "Liso", "Cacheado"}[c]
}

type Olhos int

const (
	Castanho Olhos = iota
	Azul
	Verde
	Preto
)

func (o Olhos) String() string {
	return [...]string{"Castanho", "Azul", "Verde", "Preto"}[o]
}

type Pele int

const (
	Negra Pele = iota
	Parda
	Branca
)

func (p Pele) String() string {
	return [...]string{"Negra", "Parda", "Branca"}[p]
}

type Corpo int

const (
	Magra Corpo = iota
	Atletica
	Gorda
)

func (cp Corpo) String() string {
	return [...]string{"Magra", "Atletica", "Gorda"}[cp]
}

type Info struct {
	ID     string `json:"id"`
	Id_user     string `json:"id_user"`
	Cabelo Cabelo `json:"cabelo"`
	Olhos Olhos `json:"olhos"`
	Pele Pele `json:"pele"`
	Corpo Corpo `json:"corpo"`
}

type InfoDTO struct {	
	Id_user     string `json:"id_user"`
	Cabelo string `json:"cabelo"`
	Olhos string `json:"olhos"`
	Pele string `json:"pele"`
	Corpo string `json:"corpo"`
}

func NewInfo(id_user string, cabelo string, olhos string, pele string, corpo string) (*Info, error) {
	var newcabelo Cabelo
	switch cabelo {
	case "Crespo":
		newcabelo = Crespo
	case "Liso":
		newcabelo = Liso
	case "Cacheado":
		newcabelo = Cacheado
	default:
		newcabelo = Cacheado // Valor padrão ou tratamento de erro
	}

	var newolhos Olhos
	switch olhos {
	case "Castanho":
		newolhos = Castanho
	case "Azul":
		newolhos = Azul
	case "Verde":
		newolhos = Verde
	case "Preto":
		newolhos = Preto
	default:
		newolhos = Castanho // Valor padrão ou tratamento de erro
	}

	var newpele Pele
	switch pele {
	case "Parda":
		newpele = Parda
	case "Negra":
		newpele = Negra
	case "Branca":
		newpele = Branca
	default:
		newpele = Parda // Valor padrão ou tratamento de erro
	}

	var newcorpo Corpo
	switch corpo {
	case "Magra":
		newcorpo = Magra
	case "Atletica":
		newcorpo = Atletica
	case "Gorda":
		newcorpo = Gorda
	default:
		newcorpo = Magra // Valor padrão ou tratamento de erro
	}

	info := &Info{
		ID:     uuid.NewString(),
		Id_user: id_user,
		Cabelo: newcabelo,
		Olhos: newolhos,
		Pele: newpele,
		Corpo: newcorpo,

	}

	return info, nil

}

type InfoRepository interface {
	Create(info Info) (Info, error)
}