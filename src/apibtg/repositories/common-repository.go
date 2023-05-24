package repositories

import (
	database "BTG-Test/src/apibtg/database"
	"database/sql"
	"fmt"
	"strings"

	logger "github.com/sirupsen/logrus"
)

type (
	CommonRepository interface {
		FindAll(query string) (*sql.Rows, error)
		FindById(query string, id int) (*sql.Row, error)
		StoreOne(query, colName string, data map[int]interface{}) (*sql.Row, error)
		// Update(data Customer) (Customer, error)
		DeleteById(table, column string, id int) error
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

func (c *CommonRepositoryImpl) FindById(query string, id int) (*sql.Row, error) {
	db, errDb := database.GetConnectionBTG()
	if errDb != nil {
		logger.Error(errDb)
		return nil, errDb
	}
	row := db.QueryRow(query, id)
	if row.Err() != nil {
		logger.Error(row.Err())
		return nil, row.Err()
	}
	return row, nil

}

func (c *CommonRepositoryImpl) StoreOne(query, colName string, data map[int]interface{}) (*sql.Row, error) {
	db, errDb := database.GetConnectionBTG()
	if errDb != nil {
		logger.Error(errDb)
		return nil, errDb
	}

	values := createQuery(colName, data)
	row := db.QueryRow(query + values)
	if row.Err() != nil {
		logger.Error(row.Err())
		return nil, row.Err()
	}
	return row, nil
}

func (c *CommonRepositoryImpl) DeleteById(table, column string, id int) error {
	db, errDb := database.GetConnectionBTG()
	if errDb != nil {
		logger.Error(errDb)
		return errDb
	}

	query := fmt.Sprintf("DELETE FROM \"BTG_Schema\".\"%v\" WHERE \"%v\" = $1", table, column)

	_, errQuery := db.Exec(query, id)
	if errQuery != nil {
		return errQuery
	}
	return nil
}

func createQuery(colName string, params map[int]interface{}) string {
	valSlice := make([]string, 0, len(params))

	for i := 1; i <= len(params); i++ {
		valSlice = append(valSlice, fmt.Sprintf("'%v'", params[i]))
	}
	values := strings.Join(valSlice, ",")
	return fmt.Sprintf("VALUES(%v) RETURNING \"%v\"", values, colName)
}
