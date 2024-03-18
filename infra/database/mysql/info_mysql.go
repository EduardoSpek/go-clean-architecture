package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/eduardospek/go-clean-architecture/domain/entity"
	_ "github.com/go-sql-driver/mysql"
)

type UserRepository interface {
	GetById(id string) (entity.User, error)
}

var (	
	ErrUserNotFound = errors.New("erro: Usuário não encontrado")	
	ErrInfoExists = errors.New("informações já cadastradas")		    
)

type InfoMysqlRepository struct {
	UserRepository UserRepository
}

func NewInfoMysqlRepository(repository UserRepository) *InfoMysqlRepository {
	infoRepo := &InfoMysqlRepository{ UserRepository: repository }
	infoRepo.CreateInfoTable()
	return infoRepo
}

func (repo *InfoMysqlRepository) CreateInfoTable() error {    	
    db, err := conn.Connect()
	defer db.Close()

	if err != nil {
        fmt.Println(err)
		return err
	}
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
func (repo *InfoMysqlRepository) Create(info entity.Info) (entity.InfoOutput, error) {    
    db, _ := conn.Connect()
	defer db.Close()

    InfoExists := repo.InfoExists(info.Id_user)
	if InfoExists {
		return entity.InfoOutput{}, ErrInfoExists
	}

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
func (repo *InfoMysqlRepository) InfoExists(id_user string) bool {
    db, _ := conn.Connect()
	defer db.Close()

    userQuery := "SELECT id_user FROM info WHERE id_user = ?"
    row := db.QueryRow(userQuery, id_user)    

    // Recuperando os valores do banco de dados
    err := row.Scan(&id_user)
    if err != nil {        
        if err == sql.ErrNoRows {            
            return false
        }
    }
  
    return true
}