package sqlitedb

import "project/internal/entity"

func (database *SqliteDB) CreateTableAccounts() error {
	createAccountsTable := `CREATE TABLE IF NOT EXISTS accounts (       
    id INT PRIMARY KEY,       
    balance INT NOT NULL
                                    
    );`
	_, err := database.db.Exec(createAccountsTable)
	if err != nil {
		return err
	}
	return nil
}
func (database *SqliteDB) InsertTableAccounts(account entity.Account) error {
	InsertAccountsTable := `
        INSERT INTO accounts (id, balance) VALUES (?, ?)`
	//amount, accountId
	_, err := database.db.Exec(InsertAccountsTable, account.ID, account.Balance)
	if err != nil {
		return err
	}
	return nil
}

func (database *SqliteDB) UpdateTableAccounts(account entity.Account) error {
	UpdateAccountsTable := `
        UPDATE accounts 
        SET balance = ? 
        WHERE id = ?`
	//amount, accountId
	_, err := database.db.Exec(UpdateAccountsTable, account.ID, account.Balance)
	if err != nil {
		return err
	}
	return nil
}
