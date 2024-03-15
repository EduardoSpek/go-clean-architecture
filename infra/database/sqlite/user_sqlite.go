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
	ErrUserExists = errors.New("Usuário já cadastrado com este nome")	
    ErrUserNotExistsWithID = errors.New("Não existe usuário com este ID")
)

type UserSQLiteRepository struct {}

func NewUserSQLiteRepository() *UserSQLiteRepository {
	userRepo := &UserSQLiteRepository{}
	userRepo.CreateUserTable()
	return &UserSQLiteRepository{}
}

func (repo *UserSQLiteRepository) Connect() (*sql.DB, error) {    
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
func (repo *UserSQLiteRepository) Create(user entity.User) (entity.User, error) {    
    db, _ := repo.Connect()
	defer db.Close()

    UserExists := repo.UserExists(user.Name)
	if UserExists {
		return entity.User{}, ErrUserExists
	}
 
    insertQuery := "INSERT INTO users (id,  name, zap) VALUES (?, ?, ?)"
    _, err := db.Exec(insertQuery, user.ID, user.Name, user.Zap)

    if err != nil {
		return entity.User{}, err
	}     
    
    return user, err
}

func (repo *UserSQLiteRepository) Update(user entity.User) (entity.User, error)  {    
    db, _ := repo.Connect()
	defer db.Close()

    UserExists := repo.UserExists(user.Name)
	if UserExists {
		return entity.User{}, ErrUserExists
	}
    
    _, err := repo.GetById(user.ID)
    if err != nil {
        fmt.Println(err)
		return entity.User{}, err
	}    

    query := "UPDATE users SET"
    if user.Name != "" {
		query += " name = '" + user.Name + "'"
	}
    if user.Zap != "" {
		if user.Name != "" {
			query += ","
		}
		query += " zap = '" + user.Zap + "'"
	}
	query += " WHERE id = '" + fmt.Sprint(user.ID) + "'"    

    if user.Name != "" || user.Zap != "" {
		_, err := db.Exec(query)
		if err != nil {
			fmt.Println(err)
			return entity.User{}, err
		}	
	}

    updateuser, err := repo.GetById(user.ID)

    if err != nil {
        fmt.Println(err)
		return entity.User{}, err
	}

    return updateuser, err
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
            return entity.User{}, ErrUserNotExistsWithID
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
		return nil, err
	}
    
    rows, err := db.Query("SELECT * FROM users ORDER BY name ASC")
    if err != nil {        
        return nil, err
    }    

    defer rows.Close()

    var users []entity.User
    users = []entity.User{}
    
    for rows.Next() {
        var user entity.User
        err := rows.Scan(&user.ID, &user.Name, &user.Zap)
        if err != nil {            
            return nil, err
        }
        users = append(users, user)
    }
    
    return users, nil
}

func (repo *UserSQLiteRepository) Delete(id string) (error) {
	
	db, err := repo.Connect()
	defer db.Close()

    if err != nil {        
		return err
	}

    _, err = repo.GetById(id)

    if err != nil {        
		return err
	}

    _ , err = db.Exec("DELETE FROM users WHERE id = ?", id)

    if err != nil {        
		return err
	}

    return nil

}

//VALIDATIONS
func (repo *UserSQLiteRepository) UserExists(name string) bool {
    db, _ := repo.Connect()
	defer db.Close()

    userQuery := "SELECT name FROM users WHERE name = ?"
    row := db.QueryRow(userQuery, name)    

    // Recuperando os valores do banco de dados
    err := row.Scan(&name)
    if err != nil {        
        if err == sql.ErrNoRows {            
            return false
        }
    }
  
    return true
}