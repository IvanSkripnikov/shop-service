package models

const Success = "success"
const Failure = "failure"

const ServiceDatabase = "ShopService"

type Database struct {
	Address  string
	Port     int
	User     string
	Password string
	DB       string
}

type Redis struct {
	Address  string
	Port     string
	Password string
	DB       int
	Stream   string
}
