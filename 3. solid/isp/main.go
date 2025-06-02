package main

import "fmt"

// Now, a basic printer like an OldPrinter can only Print() — it can’t Scan() or Fax(). 
// But to implement Machine, it must stub out those methods, which is unnecessary and wrong.
// type Machine interface {
// 	Print(doc string)
// 	Scan(doc string)
// 	Fax(doc string)
// }

type Printer interface {
	Print(doc string)
}

type Scanner interface {
	Scan(doc string)
}

type Photocopier struct {}

func (p *Photocopier) Print(doc string) {
	fmt.Println(doc)
}

func (p *Photocopier) Scan(doc string) {
	fmt.Println("Scanning doc ", doc)
}

type BasicPrinter struct {}

func (b *BasicPrinter) Print(doc string) {
    fmt.Println("Basic print: ", doc)
}

func main() {
	fmt.Println("Hello, World!")
	doc := "Hey i m a great doc and hey gpt how are u"
	ph := &Photocopier{}

	ph.Print(doc)
	ph.Scan(doc)

	bp := &BasicPrinter{}
	bp.Print(doc)
}