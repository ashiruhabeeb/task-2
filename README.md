-- SIMPLE CRUD API

A simple RESTFUL API with Go uing Gin Framework and Gorm; the Gin frameword will be responsible for the api router while Gorm will will be responsible for the POSTGRES database connection and operations. 

To download Gin Framework: go get github.com/gin-gonic/gin
To download Gorm: go get gorm.io/gorm, & go get gorm.io/driver/postgres

The api performs the following CRUD operations:
- CREATE: add a new person.
- READ: fetches record of a registered person by query-name.
- UPDATE: modifies record of existing person reecord.
- DELETE: deletes the record of an existing person.

- API STRUCTURE
  |--- app
  |     |--- models
  |     |      |--- models.go
  |     |--- db.go
  |     |--- router_test.go
  |     |--- router.go
  |     |--- server.go
  |
  |--- config
  |     |--- config.go
  |--- .env
  |--- .gitignore
  |--- go.mod
  |--- go.sum
  |--- main.go
  |--- README.md
  
    - App Directory

The app directory houses the models package, as well as db.go, router.go, router_test.go and server.go file. 

The models package contains the models.go file where "person" entity data structure is defined; consisting of fields such as the ID, Name, CreatedAt, UpdatedAt and DeletedAt fileds. 
The ID field is is of integer data type, it is the primary key (database purpose) and unique identifier for all records to be manipulated as deem fit.
The Name fiels is of string data type, this generally describes the specified name provided for records to be consumed ny the api.
While, the CreatedAt, UpdatedAt and DeletedAt esentially are of the time data type, logs of timestamp are entered into respective fields accordingly i.e when a record is created, updated or deleted.

 - server.go file
This file houses the "App" struct and the receiver method "StartServer(port string)". The "App" struct holds the the Gin framework engine (*gin.Engine) and Gorm engine (*gorm.DB) components. THE app struct will serve as base for the various receiver methods to be implemented for adequate functionality of the api.

While the "StartServer(port string)" receiver method is essentially responsible for spinning off the api server into action by calling the the Gin.Run() method which takes the port as paramater. This function is exported for use in the application entry program file - main.go file.

  - db.go file
This file houses the SetUpDB(dbURL string) receiver method function, this function is responsible for establishing the postgres database connection where the api data/resources will be stored or persisted. This was achived with help of the GORM package. Once the the database connection was been successfully established, the user defined "Person" data structure was also migrated to the databse using the "AutoMigrate" method function of the GORM package. This function is exported for use in the application entry program file - main.go file.

  - router.go
This file basically contains the CRUD endpoints implementation i.e POST, GET, PUT and DELETE verbs serving the CREATE, READ, UPDATE and DELETE operations respectively.


  - Config Directory

This package houses the config.go file, that handles the configuration abstraction of environment variables defined in the shell file (.env) without necessarilly exposing the stored sensitive data for our API to the outside world. This was achieved by importing the viper package, defining a config (cfg) struct to hold the defined environment variables and the LoadEnvVariables(path string)(c Cfg, err error); it takes the path parameters of string type and returns the config value and error if there is any.


  - Utils Directory

This package houses the validate.go file. It facilitates the validation of client side input to meet set condittions. This will ensure consistency in data supplied to the API. InputValidator() function is implemented in this file and exported for use to validate defined copy of the "Person" struct anyhwere in API.

  - .env file
This serves as an abstraction of my bash shell, it holds environment variables in key:pair variables. This file will not be pushed to github rrepository.

  - .gitignoe
This is a Github operational file that basically holds, monitor and ensure named files in it to be prevented from pushing up to remote branch. IN this case, it hold the (.env) name, so as to prevent exposing the contents of the environment file.

  - go.mod and go.sum files
Both of this files are responsible for managing application packages i.e downloaded external dependencies.

  - main.go file
This file essentially serves as the entry point for api application. Here the API server component base "APP" is initialized a variable and necessary components are called so as to spurn off the server operation.

  - README.md file
This file basically describes the composition and manner of functionality of this application.