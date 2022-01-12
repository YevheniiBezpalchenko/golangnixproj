package repositories

type UserRepositoryInterface interface {
	GetByEmail(email string) User
	Insert(user *User) error
}

type SupplierRepositoryInterface interface {
	GetAll() (suppliers []*Supplier, err error)
}

type OrderRepositoryInterface interface {
	Get(id int) (order *Order, err error)
	GetAll() (order []*Order, err error)
	Create() (order []*Order, err error)
}
type ProductRepositoryInterface interface {
	Get(id int) (product []*Products, err error)
	Create() (order []*Order, err error)
}
type OrderProductRepositoryInterface interface {
	GetAll() (orders []*OrderProducts, err error)
	Create(user *Products, order *Order) (err error)
}
