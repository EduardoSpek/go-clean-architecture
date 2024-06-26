package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

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
		corpo VARCHAR(50) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    )`)
    return err
}

// insertInfo insere um novo usuário no banco de dados
func (repo *InfoMysqlRepository) Create(info entity.Info) (entity.InfoOutput, error) {    
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

func (repo *InfoMysqlRepository) Update(info entity.Info) (entity.InfoOutput, error) {    
    db, _ := conn.Connect()
	defer db.Close()

	cabelo := info.Cabelo.String()
	olhos := info.Olhos.String()
	pele := info.Pele.String()
	corpo := info.Corpo.String()
 
    insertQuery := "UPDATE info SET cabelo=?, olhos=?, pele=?, corpo=? WHERE id=?"
    _, err := db.Exec(insertQuery, cabelo, olhos, pele, corpo, info.ID)

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

func (repo *InfoMysqlRepository) GetById(id string) (entity.InfoOutput, error) {
	db, err := conn.Connect()	
	
	if err != nil {
        fmt.Println("Erro ao conectar ao DB")
		return entity.InfoOutput{}, err
	}   
    
    defer db.Close()    

    userQuery := "SELECT id_user, cabelo, olhos, pele, corpo, created_at, updated_at FROM info WHERE id = ?"
    row := db.QueryRow(userQuery, id)     

    // Variáveis para armazenar os dados do usuário
    var id_user, cabelo, olhos, pele, corpo string
	var created_at, updated_at time.Time

    // Recuperando os valores do banco de dados
    err = row.Scan(&id_user, &cabelo, &olhos, &pele, &corpo, &created_at, &updated_at)
    if err != nil {        
        // Se não houver usuário correspondente ao ID fornecido, retornar nil
        if err == sql.ErrNoRows {            
            return entity.InfoOutput{}, ErrUserNotExistsWithID
        }
        // Se ocorrer outro erro, retornar o erro        
        return entity.InfoOutput{}, err
    }

    // Criando a entidade Info com os dados recuperados
    user := &entity.InfoOutput{
        ID: id,
		Id_user: id_user,
        Cabelo: cabelo,
		Olhos: olhos,
		Pele: pele,
		Corpo: corpo,
		CreatedAt: created_at.Local(),
		UpdatedAt: updated_at.Local(),
    }    
    
    return *user, err
}

func (repo *InfoMysqlRepository) GetByIdUser(id_user string) (entity.InfoOutput, error) {
	db, err := conn.Connect()	
	
	if err != nil {
        fmt.Println("Erro ao conectar ao DB")
		return entity.InfoOutput{}, err
	}   
    
    defer db.Close()    

    userQuery := "SELECT id, cabelo, olhos, pele, corpo, created_at, updated_at FROM info WHERE id_user = ?"
    row := db.QueryRow(userQuery, id_user)     

    // Variáveis para armazenar os dados do usuário
    var id, cabelo, olhos, pele, corpo string
	var created_at, updated_at time.Time

    // Recuperando os valores do banco de dados
    err = row.Scan(&id, &cabelo, &olhos, &pele, &corpo, &created_at, &updated_at)
    if err != nil {        
        // Se não houver usuário correspondente ao ID fornecido, retornar nil
        if err == sql.ErrNoRows {            
            return entity.InfoOutput{}, ErrUserNotExistsWithID
        }
        // Se ocorrer outro erro, retornar o erro        
        return entity.InfoOutput{}, err
    }

    // Criando a entidade Info com os dados recuperados
    user := &entity.InfoOutput{
        ID: id,
		Id_user: id_user,
        Cabelo: cabelo,
		Olhos: olhos,
		Pele: pele,
		Corpo: corpo,
		CreatedAt: created_at.Local(),
		UpdatedAt: updated_at.Local(),
    }    
    
    return *user, err
}

//VALIDATIONS
func (repo *InfoMysqlRepository) UserWithInfo(id_user string) error {
    db, _ := conn.Connect()
	defer db.Close()

    userQuery := "SELECT id_user FROM info WHERE id_user = ?"
    row := db.QueryRow(userQuery, id_user)    

    // Recuperando os valores do banco de dados
    err := row.Scan(&id_user)
    if err != nil {        
        if err == sql.ErrNoRows {            
            return nil
        }
    }
  
    return errors.New("erro: Usuário já tem informações")
}