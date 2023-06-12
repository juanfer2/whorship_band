package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresClient struct {
	ConfigClient    *ConfigClient
	MigrationClient *MigrationClient
	DB              *gorm.DB
}

func NewPostgresClient(
	configClient *ConfigClient,
	migrationClient *MigrationClient,
) PostgresClient {
	return PostgresClient{
		ConfigClient:    configClient,
		MigrationClient: migrationClient,
		DB: Conn(
			configClient.Username,
			configClient.Password,
			configClient.Database,
			configClient.Port,
			configClient.Host,
		),
	}
}

func (p *PostgresClient) Conn() *gorm.DB {
	return p.DB
}

func (p *PostgresClient) CreateDataBase() {
	url := fmt.Sprintf(
		"user=%s password=%s port=%s host=%s sslmode=disable",
		p.ConfigClient.Username,
		p.ConfigClient.Password,
		p.ConfigClient.Port,
		p.ConfigClient.Host,
	)

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		query := fmt.Sprintf("CREATE DATABASE %s", p.ConfigClient.Database)

		_, err = db.Exec(query)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Successfully created database..")
		}
	}

	db.Close()
}

func Conn(
	Username string,
	Password string,
	Database string,
	Port string,
	Host string,
) *gorm.DB {
	loggerConfig := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			//IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful: true, // Disable color
		},
	)

	DSN := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable port=%s host=%s",
		Username,
		Password,
		Database,
		Port,
		Host,
	)

	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  DSN,  // data source name, refer https://github.com/jackc/pgx
			PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
		}), &gorm.Config{Logger: loggerConfig})

	if err != nil {
		log.Fatalf("💥 Error while reading config file %s", err)
	}

	log.Println("✨ Database connected")

	return db
}
