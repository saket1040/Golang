package strategy

type PaymentStrategy interface {
	Pay(inserted int, price int) error
}