module vulnerable-app

go 1.19

require github.com/dgrijalva/jwt-go v3.2.0+incompatible // Known vulnerability in JWT handling
