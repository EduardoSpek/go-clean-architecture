package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/eduardospek/go-clean-arquiteture/domain/entity"
	_ "github.com/mattn/go-sqlite3"
)

type UserSQLiteRepository struct {}

func NewUserSQLiteRepository() *UserSQLiteRepository {
	userRepo := &UserSQLiteRepository{}
	userRepo.CreateUserTable()
	return &UserSQLiteRepository{}
}

func (repo *UserSQLiteRepository) Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./agenda.db")
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

func (repo *UserSQLiteRepository) CreateUserTable() error {    
    db, err := repo.Connect()
	defer db.Close()

	if err != nil {
        fmt.Println(err)
		return err
	}
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id VARCHAR(36) PRIMARY KEY NOT NULL,
        name VARCHAR(50) NOT NULL,
        zap VARCHAR(100) NOT NULL
    )`)
    return err
}

// insertUser insere um novo usuário no banco de dados
func (repo *UserSQLiteRepository) Create(user entity.User) error {    
    db, err := repo.Connect()
	defer db.Close()

	if err != nil {
		return err
	}    
    insertQuery := "INSERT INTO users (id,  name, zap) VALUES (?, ?, ?)"
    _, err = db.Exec(insertQuery, user.ID, user.Name, user.Zap)
    return err
}

func (repo *UserSQLiteRepository) GetById(id string) (entity.User, error) {
	db, err := repo.Connect()
	defer db.Close()
	
	if err != nil {
        fmt.Println("Erro ao conectar ao DB")
		return entity.User{}, err
	}    

    userQuery := "SELECT name, zap FROM users WHERE id = ?"
    row := db.QueryRow(userQuery, id)    

    // Variáveis para armazenar os dados do usuário
    var name, zap string

    // Recuperando os valores do banco de dados
    err = row.Scan(&name, &zap)
    if err != nil {        
        // Se não houver usuário correspondente ao ID fornecido, retornar nil
        if err == sql.ErrNoRows {            
            return entity.User{}, err
        }
        // Se ocorrer outro erro, retornar o erro        
        return entity.User{}, err
    }

    // Criando um objeto models.User com os dados recuperados
    user := &entity.User{
        ID: id,
        Name: name,
        Zap:    zap,
    }
    
    return *user, err
}

func (repo *UserSQLiteRepository) List() ([]entity.User, error) {
	
	db, err := repo.Connect()
	defer db.Close()

	if err != nil {
        fmt.Printf("Erro ao conectar com o banco de dados")
		return nil, err
	}
    
    rows, err := db.Query("SELECT * FROM users ORDER BY name ASC")
    if err != nil {
        fmt.Println("Erro ao selecionar usuarios")
        return nil, err
    }    

    defer rows.Close()

    var users []entity.User
    users = []entity.User{}
    
    for rows.Next() {
        var user entity.User
        err := rows.Scan(&user.ID, &user.Name, &user.Zap)
        if err != nil {
            fmt.Println("Erro ao listar usuarios")
            return nil, err
        }
        users = append(users, user)
    }
    
    return users, nil
}
