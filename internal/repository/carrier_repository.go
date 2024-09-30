package repository

import (
	"database/sql"
	"fmt"

	"github.com/L3onix/frete-rapido-test/internal/model"
)

type CarrierRepository struct {
	connection *sql.DB
}

func NewCarrierRepository(connection *sql.DB) CarrierRepository {
	return CarrierRepository{
		connection: connection,
	}
}

func (cr *CarrierRepository) GetCarriers() ([]model.Carrier, error) {
	query := "select * from carrier"
	rows, err := cr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Carrier{}, err
	}

	var carrierList []model.Carrier
	var carrierObj model.Carrier

	for rows.Next() {
		err = rows.Scan(
			&carrierObj.ID,
			&carrierObj.Name,
			&carrierObj.Service,
			&carrierObj.Deadline,
			&carrierObj.Price,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Carrier{}, err
		}

		carrierList = append(carrierList, carrierObj)
	}

	rows.Close()

	return carrierList, nil
}

func (cr *CarrierRepository) CreateCarrier(carrier model.Carrier) (int, error) {
	var id int

	queryStr := "insert into carrier (name, service, deadline, price, dispatcher_id) values ($1, $2, $3, $4, $5) returning id"
	query, err := cr.connection.Prepare(queryStr)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(carrier.Name, carrier.Service, carrier.Deadline, carrier.Price, carrier.DispatcherID).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (cr *CarrierRepository) GetCarrierById(id int) (*model.Carrier, error) {
	query, err := cr.connection.Prepare("select * from carrier where id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var carrier model.Carrier

	err = query.QueryRow(id).Scan(
		&carrier.ID,
		&carrier.Name,
		&carrier.Service,
		&carrier.Deadline,
		&carrier.Price,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &carrier, err
}

func (cr *CarrierRepository) GetLastCarriers(lastCarriers string) (*[]model.Carrier, error) {
	var rows *sql.Rows
	var err error
	queryStr := "select * from carrier order by id desc"
	if lastCarriers != "" {
		queryStr += " limit $1"

		query, err := cr.connection.Prepare(queryStr)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		rows, err = query.Query(lastCarriers)
	} else {
		rows, err = cr.connection.Query(queryStr)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	carrierList, err := cr.getLastCarriersResultFormat(rows)
	rows.Close()
	if err != nil {
		return nil, err
	}

	return carrierList, nil
}

func (cr *CarrierRepository) GetLastCarriersByDispatcherID(lastCarriers string, dispatcherID string) (*[]model.Carrier, error) {
	var rows *sql.Rows
	var err error
	queryStr := "select * from carrier where dispatcher_id = $1 order by id desc"
	query, err := cr.connection.Prepare(queryStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rows, err = query.Query(dispatcherID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	carrierList, err := cr.getLastCarriersResultFormat(rows)
	rows.Close()
	if err != nil {
		return nil, err
	}

	return carrierList, nil
}

func (cr *CarrierRepository) getLastCarriersResultFormat(rows *sql.Rows) (*[]model.Carrier, error) {
	var carrierList []model.Carrier
	var carrierObj model.Carrier

	for rows.Next() {
		err := rows.Scan(
			&carrierObj.ID,
			&carrierObj.Name,
			&carrierObj.Service,
			&carrierObj.Deadline,
			&carrierObj.Price,
			&carrierObj.DispatcherID,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		carrierList = append(carrierList, carrierObj)
	}
	return &carrierList, nil
}
