package repositories

import "github.com/YevheniiBezpalchenko/golangnixproj/internal/models"

type UserRepositoryInterface interface {
	GetByEmail(email string) models.User
	GetById(Id int) models.User
	Create(user *models.User) error
	Delete(Id int) error
	Update(id int, user *models.User) models.User
	GetAll() []models.User
}

type SupplierRepositoryInterface interface {
	GetAll() (suppliers []*models.Supplier, err error)
}

type OrderRepositoryInterface interface {
	Get(id int) (order *models.Order, err error)
	GetAll() (order []*models.Order, err error)
	Create() (order []*models.Order, err error)
}
type ProductRepositoryInterface interface {
	Get(id int) (product []*models.Products, err error)
	Create() (order []*models.Order, err error)
}
type OrderProductRepositoryInterface interface {
	GetAll() (orders []*models.OrderProducts, err error)
	Create(user *models.Products, order *models.Order) (err error)
}
