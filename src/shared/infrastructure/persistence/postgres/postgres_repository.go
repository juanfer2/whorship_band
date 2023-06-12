package postgres

import (
	"errors"
	"fmt"
	"log"

	persistence_domain "github.com/juanfer2/whorship_band/src/shared/infrastructure/persistence/domain"
	"github.com/juanfer2/whorship_band/src/shared/utilities"
	"gorm.io/gorm"
)

type PostgresRepository[T any] struct {
	PostgresClient *PostgresClient
}

func NewPostgresRepository[T any](
	postgresClient *PostgresClient,
) persistence_domain.Repository[T] {
	return &PostgresRepository[T]{
		PostgresClient: postgresClient,
	}
}

func (pr *PostgresRepository[T]) Create(newRecord T) (T, error) {
	fmt.Println("newRecord")
	utilities.PrintJson(newRecord)

	err := pr.PostgresClient.Conn().Create(&newRecord)

	if err != nil {
		return newRecord, err.Error
	}

	return newRecord, nil
}

/*
	func (pr *PostgresRepository[T]) Create(newRecord T) (T, error) {
		err := pr.PostgresClient.Conn().Create(&newRecord)
		if err != nil {
			return newRecord, err.Error
		}

		return newRecord, nil
	}
*/
func (pr *PostgresRepository[T]) FindBy(id int) T {
	var record T
	pr.PostgresClient.Conn().First(&record, id)

	return record
}

func (pr *PostgresRepository[T]) FindByUuid(id string) (T, error) {
	var record T
	err := pr.PostgresClient.Conn().First(&record, "id = ?", id).Error

	fmt.Println("errors.Is(err, gorm.ErrRecordNotFound)")
	fmt.Println(errors.Is(err, gorm.ErrRecordNotFound))

	fmt.Println("---------_>err")
	fmt.Println(err)

	if err != nil {
		fmt.Println(err)
		return record, err
	}

	return record, nil
}

func (pr *PostgresRepository[T]) All() []T {
	var records []T
	defer pr.PostgresClient.Conn()
	pr.PostgresClient.Conn().Find(&records)

	return records
}

func (pr *PostgresRepository[T]) Delete(record T) T {
	pr.PostgresClient.Conn().Delete(&record)

	return record
}

func (pr *PostgresRepository[T]) Save(record T) T {
	pr.PostgresClient.Conn().Save(&record)

	return record
}

func (pr *PostgresRepository[T]) UpdateByUuid(id string, updateRecord T) (T, error) {
	record, err := pr.FindByUuid(id)

	fmt.Println("---------_>err")
	fmt.Println(err)

	if err != nil {
		return record, err
	}

	resultError := pr.PostgresClient.Conn().Model(&record).Updates(updateRecord).Error
	if resultError != nil {
		return record, resultError
	}

	return record, nil
}

func (pr *PostgresRepository[T]) Update(updateRecord T) (T, error) {
	err := pr.PostgresClient.Conn().Model(&updateRecord).Updates(updateRecord)
	if err != nil {
		return updateRecord, err.Error
	}

	return updateRecord, nil
}

func (pr *PostgresRepository[T]) Where(query any, arg ...any) *gorm.DB {
	var table T
	return pr.PostgresClient.Conn().Model(&table).Where(query, arg)
}

func (pr *PostgresRepository[T]) WhereBy(query any, arg ...any) T {
	var table T
	pr.PostgresClient.Conn().Where(query, arg).First(&table)

	return table
}

func (pr *PostgresRepository[T]) Not(query any, arg ...any) []T {
	var table []T
	pr.PostgresClient.Conn().Not(query, arg).Find(&table)

	return table
}

func (pr *PostgresRepository[T]) CreateInBatches(data []T) {
	pr.PostgresClient.Conn().CreateInBatches(data, 100)
}

func (pr *PostgresRepository[T]) Join(relations ...string) *gorm.DB {
	var table T
	model := pr.PostgresClient.Conn().Model(&table)

	for _, relationName := range relations {
		joinRelation()(model, relationName)
	}

	if model.Error != nil {
		log.Fatalln(model.Error)
	}

	return model
}

type JoinRelation = func(g *gorm.DB, name string) *gorm.DB

func joinRelation() JoinRelation {
	return func(db *gorm.DB, name string) *gorm.DB {
		return db.Joins(name)
	}
}

func (pr *PostgresRepository[T]) Preload(relations ...string) *gorm.DB {
	pr.PostgresClient.Conn()
	var table T
	model := pr.PostgresClient.Conn().Model(&table)

	for _, relationName := range relations {
		preloadRelation()(model, relationName)
	}

	if model.Error != nil {
		log.Fatalln(model.Error)
	}

	return model
}

type RelationPreload = func(g *gorm.DB, name string) *gorm.DB

func preloadRelation() RelationPreload {
	return func(db *gorm.DB, name string) *gorm.DB {
		return db.Preload(name)
	}
}

func (pr *PostgresRepository[T]) FindWithAssociation(
	associations []string, query any, arg ...any,
) T {
	var record T
	var table T
	model := pr.PostgresClient.Conn().Model(&table).Where(query, arg)

	for _, association := range associations {
		fmt.Println(association)
		associationAppend()(model, association)
	}

	fmt.Println(model.Statement.Vars)

	err := model.First(&record).Error
	if err != nil {
		fmt.Println("***Error***")
		fmt.Println(err)
	}

	return record
}

func (pr *PostgresRepository[T]) WhereWithAssociation(
	associations []string, query any, arg ...any,
) []T {
	var records []T
	var table T
	model := pr.PostgresClient.Conn().Model(&table).Where(query, arg)

	for _, association := range associations {
		associationAppend()(model, association)
	}

	model.Find(&records)

	return records
}

type AssociationFunction = func(g *gorm.DB, name string) *gorm.Association

func associationAppend() AssociationFunction {
	return func(db *gorm.DB, name string) *gorm.Association {
		return db.Association(name)
	}
}
