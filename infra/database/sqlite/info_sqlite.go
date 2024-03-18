package sqlite

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/eduardospek/go-clean-architecture/domain/entity"
	_ "github.com/mattn/go-sqlite3"
)

type UserRepository interface {
	GetById(id string) (entity.User, error)
}

var (
	ErrUserNotFound = errors.New("erro: Usuário não encontrado")	    
	ErrInfoExists = errors.New("informações já cadastradas")	    
)

type InfoSQLiteRepository struct {
	UserRepository UserRepository
}

func NewInfoSQLiteRepository(repository UserRepository) *InfoSQLiteRepository {	
	infoRepo := &InfoSQLiteRepository{ UserRepository: repository }
	infoRepo.CreateInfoTable()
	return infoRepo
}

func (repo *InfoSQLiteRepository) CreateInfoTable() error {    
    db, err := conn.Connect()
	if err != nil {
        fmt.Println(err)
		return err
	}
	defer db.Close()

    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS info (
        id VARCHAR(36) PRIMARY KEY NOT NULL,
		id_user VARCHAR(36) NOT NULL,
        cabelo VARCHAR(50) NOT NULL,
        olhos VARCHAR(50) NOT NULL,
		pele VARCHAR(50) NOT NULL,
		corpo VARCHAR(50) NOT NULL
    )`)
    return err
}

// insertInfo insere um novo usuário no banco de dados
func (repo *InfoSQLiteRepository) Create(info entity.Info) (entity.InfoOutput, error) {    
    db, _ := conn.Connect()
	defer db.Close()

	cabelo := info.Cabelo.String()
	olhos := info.Olhos.String()
	pele := info.Pele.String()
	corpo := info.Corpo.String()
 
    insertQuery := "INSERT INTO info (id, id_user, cabelo, olhos, pele, corpo) VALUES (?, ?, ?, ?, ?, ?)"
    _, err := db.Exec(insertQuery, info.ID, info.Id_user, cabelo, olhos, pele, corpo)

    if err != nil {
		return entity.InfoOutput{}, err
	}   
	
	infoOutput := entity.InfoOutput {
		ID: info.ID,
		Id_user: info.Id_user,
		Cabelo: info.Cabelo.String(),
		Olhos: info.Olhos.String(),
		Pele: info.Pele.String(),
		Corpo: info.Corpo.String(),
	}
    
    return infoOutput, err
}

//VALIDATIONS
func (repo *InfoSQLiteRepository) InfoExists(id_user string) error {
    db, _ := conn.Connect()
	defer db.Close()

    userQuery := "SELECT id_user FROM info WHERE id_user = ?"
    row := db.QueryRow(userQuery, id_user)    

    // Recuperando os valores do banco de dados
    err := row.Scan(&id_user)
    if err != nil {        
        if err == sql.ErrNoRows {            
            return err
        }
    }
  
    return nil
}