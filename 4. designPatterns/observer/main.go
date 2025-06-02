package main

import "fmt"

type Stock struct {
	Name  string
	Price float64
	Observers []Observer
}

func (s *Stock) Register(observer Observer) {
    for _, obs := range s.Observers {
        if obs == observer {
            return // already registered
        }
    }
    s.Observers = append(s.Observers, observer)
}

func (s *Stock) Unregister(observer Observer) {
	for i, o := range s.Observers {
		if o == observer {
			s.Observers = append(s.Observers[:i], s.Observers[i+1:]...)
			break
		}
	}
}

func (s *Stock) SetPrice(newPrice float64) {
	s.Price = newPrice
	s.NotifyAll()
}

func (s *Stock) NotifyAll() {
	for _, val := range s.Observers {
		val.Update(s.Name, s.Price)
	}
}

type Observer interface{
	Update(stockName string, price float64)
}

type EmailClient struct{}

func (e *EmailClient) Update(stockName string, price float64) {
	fmt.Println("updated price by Email", stockName, price)
}

type TradingBot struct{}

func (t *TradingBot) Update(stockName string, price float64) {
	fmt.Println("updated price by Bot", stockName, price)
}

func main() {
	stock := &Stock{
		Name: "TCS",
	}
	email := &EmailClient{}
	bot := &TradingBot{}

	stock.Register(email)
	stock.Register(bot)

	stock.SetPrice(3500.50)
}
