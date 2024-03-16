package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/eduardospek/go-clean-arquiteture/domain/entity"
	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrInfoExists = errors.New("Informações já cadastradas")	    
)

type InfoSQLiteRepository struct {}

func NewInfoSQLiteRepository() *InfoSQLiteRepository {
	infoRepo := &InfoSQLiteRepository{}
	infoRepo.CreateInfoTable()
	return &InfoSQLiteRepository{}
}

func (repo *InfoSQLiteRepository) Connect() (*sql.DB, error) {    
	db, err := sql.Open("sqlite3", os.Getenv("PATCH_DB_SQLITE"))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (repo *InfoSQLiteRepository) CreateInfoTable() error {    
    db, err := repo.Connect()
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
func (repo *InfoSQLiteRepository) Create(info entity.Info) (entity.Info, error) {    
    db, _ := repo.Connect()
	defer db.Close()

    InfoExists := repo.InfoExists(info.Id_user)
	if InfoExists {
		return entity.Info{}, ErrInfoExists
	}

	cabelo := string(info.Cabelo.String())
 
    insertQuery := "INSERT INTO info (id, id_user, cabelo, olhos, pele, corpo) VALUES (?, ?, ?, ?, ?, ?)"
    _, err := db.Exec(insertQuery, info.ID, info.Id_user, cabelo, info.Olhos.String(), info.Pele.String(), info.Corpo)

    if err != nil {
		return entity.Info{}, err
	}     
    
    return info, err
}

//VALIDATIONS
func (repo *InfoSQLiteRepository) InfoExists(id_user string) bool {
    db, _ := repo.Connect()
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