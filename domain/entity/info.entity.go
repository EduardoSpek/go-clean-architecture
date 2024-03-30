package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Cabelo int

const (
	Crespo Cabelo = iota
	Liso
	Cacheado
)

func (c Cabelo) String() string {
	return [...]string{"Crespo", "Liso", "Cacheado"}[c]
}

func ParseCabelo(str string) (Cabelo, error) {
    switch str {
    case "Crespo":
        return Crespo, nil
    case "Liso":
        return Liso, nil
    case "Cacheado":
        return Cacheado, nil    
    default:
        return -1, fmt.Errorf("tipo de cabelo inválido: %s", str)
    }
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

func ParseOlhos(str string) (Olhos, error) {
    switch str {
    case "Azul":
        return Azul, nil
    case "Verde":
        return Verde, nil
    case "Castanho":
        return Castanho, nil
    case "Preto":
        return Preto, nil
    default:
        return -1, fmt.Errorf("cor dos olhos inválida: %s", str)
    }
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

func ParsePele(str string) (Pele, error) {
    switch str {
    case "Negra":
        return Negra, nil
    case "Parda":
        return Parda, nil
    case "Branca":
        return Branca, nil    
    default:
        return -1, fmt.Errorf("cor de pele inválida: %s", str)
    }
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

func ParseCorpo(str string) (Corpo, error) {
    switch str {
    case "Magra":
        return Magra, nil
    case "Atletica":
        return Atletica, nil
    case "Gorda":
        return Gorda, nil    
    default:
        return -1, fmt.Errorf("tipo de corpo inválido: %s", str)
    }
}

type Info struct {
	ID     string `json:"id"`
	Id_user     string `json:"id_user"`
	Cabelo Cabelo `json:"cabelo"`
	Olhos Olhos `json:"olhos"`
	Pele Pele `json:"pele"`
	Corpo Corpo `json:"corpo"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type InfoInput struct {		
	ID     string `json:"id"`
	Id_user string `json:"id_user"`
	Cabelo string `json:"cabelo"`
	Olhos string `json:"olhos"`
	Pele string `json:"pele"`
	Corpo string `json:"corpo"`
}
type InfoOutput struct {	
	ID string `json:"id"`
	Id_user string `json:"-"`
	Cabelo string `json:"cabelo"`
	Olhos string `json:"olhos"`
	Pele string `json:"pele"`
	Corpo string `json:"corpo"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewInfo(info InfoInput) (*Info, error) {
	cabelo, err := ParseCabelo(info.Cabelo)
	
	if err != nil {
		return nil, err
	}

	olhos, err := ParseOlhos(info.Olhos)
	
	if err != nil {
		return nil, err
	}
	
	pele, err := ParsePele(info.Pele)

	if err != nil {
		return nil, err
	}

	corpo, err := ParseCorpo(info.Corpo)

	if err != nil {
		return nil, err
	}

	if info.ID == "" {
		info.ID = uuid.NewString()
	}
	
	newinfo := &Info{
		ID:     info.ID,
		Id_user: info.Id_user,
		Cabelo: cabelo,
		Olhos: olhos,
		Pele: pele,
		Corpo: corpo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

	}

	return newinfo, nil

}