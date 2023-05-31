package service

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	"fmt"
)

// Service is package for any logic needed in this program

type ServiceInterface interface {
	AddCart(productName string, quantity int) error
	RemoveCart(productName string) error
	ShowCart() ([]entity.CartItem, error)
	ResetCart() error
	GetAllProduct() ([]entity.Product, error)
	Pay(money int) (entity.PaymentInformation, error)
}

type Service struct {
	database database.DatabaseInterface
}

func NewService(database database.DatabaseInterface) *Service {
	return &Service{
		database: database,
	}
}

func (s *Service) AddCart(productName string, quantity int) error {
	product, err := s.database.GetProductByName(productName)
	if err != nil {
		return err
	} // TODO: replace this

	cartItems, err := s.database.GetCartItems()
	if err != nil {
		return err
	}

	if quantity <= 0 {
		return fmt.Errorf("invalid quantity")
	}

	cartItems = append(cartItems,
		entity.CartItem{
			ProductName: product.Name,
			Quantity:    quantity,
			Price:       product.Price,
		},
	)
	err = s.database.SaveCartItems(cartItems)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) RemoveCart(productName string) error {
	_, err := s.database.GetProductByName(productName)
	if err != nil {
		return err
	}

	cartItems, err := s.database.GetCartItems()
	if err != nil {
		return err
	}

	newCartItems := []entity.CartItem{}
	isInCart := false
	for i, cartItem := range cartItems {
		if cartItem.ProductName != productName {
			newCartItems = append(newCartItems, cartItems[i])
			isInCart = true
		}
	}

	if !isInCart {
		return fmt.Errorf("product not found")
	}

	err = s.database.SaveCartItems(newCartItems)
	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (s *Service) ShowCart() ([]entity.CartItem, error) {
	carts, err := s.database.GetCartItems()
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (s *Service) ResetCart() error {
	emptyCartItems := []entity.CartItem{}
	err := s.database.SaveCartItems(emptyCartItems)
	return err // TODO: replace this
}

func (s *Service) GetAllProduct() ([]entity.Product, error) {
	return s.database.GetProductData(), nil // TODO: replace this
}

func (s *Service) Pay(money int) (entity.PaymentInformation, error) {
	cartItems, err := s.ShowCart()
	if err != nil {
		return entity.PaymentInformation{}, err
	}

	totalPrice := 0
	for _, cartItem := range cartItems {
		totalPrice += cartItem.Price * cartItem.Quantity
	}

	change := money - totalPrice
	if change < 0 {
		return entity.PaymentInformation{}, fmt.Errorf("money is not enough")
	}
	s.ResetCart()

	return entity.PaymentInformation{
		TotalPrice:  totalPrice,
		Change:      change,
		ProductList: cartItems,
		MoneyPaid:   money,
	}, nil 
}
