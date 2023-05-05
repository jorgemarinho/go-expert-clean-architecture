package database

import (
	"database/sql"

	"github.com/jorgemarinho/go-expert-clean-architecture/internal/entity"
)

type OrderRepository struct {
	*sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.DB.Prepare("insert into orders (id, price, tax, final_price) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.DB.QueryRow("select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) FindAll() ([]entity.Order, error) {
	rows, err := r.DB.Query("select id, price, tax, final_price from orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []entity.Order
	for rows.Next() {
		var id string
		var price, tax, finalPrice float64
		if err := rows.Scan(&id, &price, &tax, &finalPrice); err != nil {
			return nil, err
		}
		orders = append(orders, entity.Order{ID: id, Price: price, Tax: tax, FinalPrice: finalPrice})
	}
	return orders, nil
}
