package domain

type Config interface {
	InitDB() (string, error)
	InitMigrate() interface{}
	InitSeeder() interface{}
}
