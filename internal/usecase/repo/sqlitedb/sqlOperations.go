package sqlitedb

import "project/internal/entity"

func (database *SqliteDB) CreateTableOperations() error {
	createOperationsTable := `CREATE TABLE IF NOT EXISTS operations (       
    id INT PRIMARY KEY,
    account_id INT,    
    amount INT NOT NULL,      
    type VARCHAR(6),       
    comment VARCHAR(100),       
    date VARCHAR(50),
    FOREIGN KEY (account_id) REFERENCES accounts(id));` //date -  datetime, id - AUTO_INCREMENT
	var _, err = database.db.Exec(createOperationsTable)
	if err != nil {
		return err
	}
	return nil
}

func (database *SqliteDB) InsertTableOperations(operation entity.Operation) error {
	insertOperation := `INSERT INTO operations(id, account_id, amount, type, comment, date) VALUES (?, ?, ?, ?, ?, ?);`
	_, err := database.db.Exec(insertOperation, operation.ID, operation.AccountID, operation.Amount, operation.Type, operation.Comment, operation.Date)
	if err != nil {
		return err
	}
	//добавить логгер
	return nil
}

func (database *SqliteDB) SelectTableOperations(account entity.Account) ([]entity.Operation, error) {
	selectOperations := `SELECT * FROM operations
where account_id = ?`
	rows, err := database.db.Query(selectOperations, account.ID)
	if err != nil {
		return nil, err
	}
	operations := []entity.Operation{}
	for rows.Next() {
		var operation entity.Operation
		err = rows.Scan(&operation.ID, &operation.AccountID, &operation.Amount, &operation.Type, &operation.Comment, &operation.Date)
		if err != nil {
			return nil, err
		}
		operations = append(operations, operation)
	}
	//добавить логгер
	return operations, nil
}
