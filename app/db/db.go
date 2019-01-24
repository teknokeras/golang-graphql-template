package db

import (
	"os"
	"fmt"
	"strings"

    "github.com/go-pg/pg"
    "github.com/go-pg/pg/orm"

	"github.com/teknokeras/golang-graphql-template/app/modules"
)


func NewDB() *pg.DB{

	var dbHost strings.Builder

	dbHost.WriteString(os.Getenv("POSTGRES_HOST"))
	dbHost.WriteString(":")
	dbHost.WriteString(os.Getenv("POSTGRES_DB_PORT"))
	
	fmt.Println(dbHost.String())

	db := pg.Connect(&pg.Options{
        User: os.Getenv("POSTGRES_USER"),
        Password: os.Getenv("POSTGRES_PASSWORD"),
        Database: os.Getenv("POSTGRES_DB"),
        Addr: dbHost.String(),
    })

    for _, model := range modules.ModelList {
        err := db.CreateTable(model, &orm.CreateTableOptions{
            IfNotExists: true,
        })
        if err != nil {
            return nil
        }
    }

    modules.InitFixtures(db)

    return db

}

/*
import (
	//"os"
	//"fmt"
	//"strings"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/role"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/user"

)

type DB struct {
	DB *gorm.DB
}

func NewDB() (*DB, error) {
	// connect to the example db, create if it doesn't exist
	db, err := gorm.Open("postgres", "host=db port=5432 user=appdb dbname=golangdb password=letmein sslmode=disable")
  
	//db, err := gorm.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// drop tables and all data, and recreate them fresh for this run
	db.DropTableIfExists(&user.User{}, &role.Role{})
	db.AutoMigrate(&user.User{}, &role.Role{})

	return &DB{db}, nil
}

//Every models need to change this file so that the corresponding table can be created
func createTables(db *pg.DB){
	//User
	fmt.Println("Create User Table")
	err := db.CreateTable(&user.User{}, &orm.CreateTableOptions{
    	Temp: true,
	})

	if err != nil {
		panic(err)
	}

	//Role
	fmt.Println("Create Role Table")
	err = db.CreateTable(&role.Role{}, &orm.CreateTableOptions{
    	Temp: true,
	})

	if err != nil {
		panic(err)
	}	
}

func initFixtures(db *pg.DB){
	r := role.Role{Name: os.Getenv("DEFAULT_ADMINISTRATOR_ROLE")}

	err := db.Insert(&r)

	if err != nil{
		fmt.Println("Cannot create Default Role")
		panic(err)
		return
	}

	u := user.User{
		Name: os.Getenv("DEFAULT_ADMIN_FULL_NAME"), 
		Email: os.Getenv("DEFAULT_ADMIN_EMAIL"), 
		Password: os.Getenv("DEFAULT_ADMIN_PASSWORD"), 
		UserRole: &r,
	}

	err = db.Insert(&u)

	if err != nil{
		fmt.Println("Cannot create Default User")
		panic(err)
	}	
}


func NewDb() *pg.DB{
	var dbHost strings.Builder

	dbHost.WriteString(os.Getenv("POSTGRES_HOST"))
	dbHost.WriteString(":")
	dbHost.WriteString(os.Getenv("POSTGRES_DB_PORT"))
	
	fmt.Println(dbHost.String())

	db := pg.Connect(&pg.Options{
        User: os.Getenv("POSTGRES_USER"),
        Password: os.Getenv("POSTGRES_PASSWORD"),
        Database: os.Getenv("POSTGRES_DB"),
        Addr: dbHost.String(),
    })

    //createTables(db)

		//User
	fmt.Println("Create User Table")
	err := db.CreateTable(&user.User{}, &orm.CreateTableOptions{
    	Temp: true,
	})

	if err != nil {
		panic(err)
	}

	//Role
	fmt.Println("Create Role Table")
	err = db.CreateTable(&role.Role{}, &orm.CreateTableOptions{
    	Temp: true,
	})

	if err != nil {
		panic(err)
	}	

    initFixtures(db)

    return db
}
*/