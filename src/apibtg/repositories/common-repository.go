package repositories

import (
	database "BTG-Test/src/apibtg/database"
	"database/sql"

	logger "github.com/sirupsen/logrus"
)

type (
	CommonRepository interface {
		FindAll(query string) (*sql.Rows, error)
		FindOne(query string, args ...interface{}) (*sql.Row, error)
		// StoreOne(data Customer) (Customer, error)
		// Update(data Customer) (Customer, error)
		// DeleteById(id int) error
	}
)

type CommonRepositoryImpl struct{}

func CreateCommonRepositoryImpl() CommonRepository {
	return &CommonRepositoryImpl{}

}

func (c *CommonRepositoryImpl) FindAll(query string) (*sql.Rows, error) {
	db, errDb := database.GetConnectionBTG()
	if errDb != nil {
		logger.Error(errDb)
		return nil, errDb
	}

	rows, errRows := db.Query(query)
	if errRows != nil {
		if errRows.Error() == sql.ErrNoRows.Error() {
			return nil, errRows
		}
		return nil, errRows
	}

	return rows, nil
}

func (c *CommonRepositoryImpl) FindOne(query string, args ...interface{}) (*sql.Row, error) {
	db, errDb := database.GetConnectionBTG()
	if errDb != nil {
		logger.Error(errDb)
		return nil, errDb
	}

	row := db.QueryRow(query, args)
	if row.Err() != nil {
		logger.Error(row.Err())
		return nil, row.Err()
	}
	return row, nil
}
