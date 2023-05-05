package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	FindAll() ([]Order, error)
}
