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

func NewInfo(info Info) (*Info, error) {
	newinfo := &Info{
		ID:     uuid.NewString(),
		Id_user: info.Id_user,
		Cabelo: info.Cabelo,
		Olhos: info.Olhos,
		Pele: info.Pele,
		Corpo: info.Corpo,

	}

	return newinfo, nil

}

type InfoRepository interface {
	Create(info Info) (Info, error)
}