package stateifc

type State interface {
	InsertMoney(amount int) error
	Dispense() error
}