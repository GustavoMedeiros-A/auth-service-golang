# 1 - Install Gin to framerork

- **github.com/gin-gonic/gin**

# 2 - install all packages

- We gonna use GORM (a ORM libary for Golang), BCrypt (to password hash), jwt-go to generate JWT token and Compile Daemon to
  update the code when the file was change

- **go get gorm.io/driver/postgres**
- **go get golang.org/x/crypto/bcrypt**
- **go get github.com/golang-jwt/jwt/v4**
- **go get github.com/joho/godotenv**
- **go get github.com/githubnemo/CompileDaemon**
- **go install github.com/githubnemo/CompileDaemon**

- After all that, config the project as you can see in the files and run **compiledaemon --command="./auth-service"** to
  run in a compileDaemon mode
