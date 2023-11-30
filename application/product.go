package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductServiceInterface interface {
	Get(id string) (*ProductInterface, error)
	Create(name string, price float32) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReaderInterface interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriterInterface interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReaderInterface
	ProductWriterInterface
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float32
}

const (
	DISABLED = "Disabled"
	ENABLED  = "Enabled"
)

type Product struct {
	Id     string  `json:"id" valid:"uuidv4"`
	Name   string  `json:"name" valid:"required"`
	Price  float32 `json:"price" valid:"float, optional"`
	Status string  `json:"status" valid:"required"`
}

func NewProduct() *Product {
	product := Product{
		Id:     uuid.NewV4().String(),
		Status: DISABLED,
	}
	return &product
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}
	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("the status must be enabled or disabled")
	}
	if p.Price < 0 {
		return false, errors.New("the status must be greater or equal to zero")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 1 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("the price must be greater than zero to enable the product")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New("the price must be zero in order for the price to be changed")
}

func (p *Product) GetId() string {
	return p.Id
}
func (p *Product) GetName() string {

	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float32 {
	return p.Price
}
