package repository

import (
	"database/sql"
	"fmt"
	"meditrack/model"
)

type NCMRepository struct {
	connection *sql.DB
}

func NewNCMRepository(db *sql.DB) NCMRepository {
	return NCMRepository{
		connection: db,
	}
}

func (n *NCMRepository) GetAllNCM() ([]model.NCM, error) {
	query := `SELECT * FROM ncm`

	rows, err := n.connection.Query(query)
	if err != nil {
		return []model.NCM{}, err
	}
	defer rows.Close()

	var ncmList []model.NCM

	for rows.Next() {
		var ncm model.NCM

		if err := rows.Scan(
			&ncm.Code,
			&ncm.Description,
			&ncm.FullDescription,
			&ncm.Ex,
		); err != nil {
			return []model.NCM{}, err
		}

		ncmList = append(ncmList, ncm)
	}

	return ncmList, nil
}

func (n *NCMRepository) GetNCMByCode(code string) (*model.NCM, error) {
	query := `SELECT * FROM ncm WHERE code = ?`

	var ncm model.NCM

	if err := n.connection.QueryRow(query, code).Scan(
		&ncm.Code,
		&ncm.Description,
		&ncm.FullDescription,
		&ncm.Ex,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &ncm, nil
}

func (n *NCMRepository) CreateNCM(ncm model.NCM) (model.NCM, error) {
	query := `INSERT INTO ncm (code, description, full_description, ex) VALUES (?, ?, ?, ?)`

	_, err := n.connection.Exec(query, ncm.Code, ncm.Description, ncm.FullDescription, ncm.Ex)
	if err != nil {
		fmt.Println(err)
		return model.NCM{}, err
	}

	return ncm, nil
}
