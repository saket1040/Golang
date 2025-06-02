package main

import "fmt"

// Base Bird interface
type Bird interface {
    Eat()
}

// FlyingBird extends Bird with flying capability
/*
Yes, this is an example of interface embedding in Go. The FlyingBird interface embeds the Bird interface, meaning that any type that implements FlyingBird must also implement all the methods of Bird in addition to the methods defined in FlyingBird.
this means that any type satisfying FlyingBird must implement both Eat() (from Bird) and Fly().
It’s especially useful when you want to create specialized interfaces that build on more general ones.

Interface embedding is allowed and means "must implement all methods of the embedded interface(s) plus any new ones.
*/
type FlyingBird interface {
    Bird
    Fly()
}

type Sparrow struct{}

func (s *Sparrow) Eat() {
    fmt.Println("Sparrow is eating")
}

func (s *Sparrow) Fly() {
    fmt.Println("Sparrow is flying")
}

type Penguin struct{}

func (p *Penguin) Eat() {
    fmt.Println("Penguin is eating")
}

// Function that expects a FlyingBird (not all birds)
func MakeBirdFly(fb FlyingBird) {
    fb.Fly()
}

func main() {
    sparrow := &Sparrow{}
    penguin := &Penguin{}

    sparrow.Eat()
    penguin.Eat()

    MakeBirdFly(sparrow)
    // MakeBirdFly(penguin) // ❌ compile error if uncommented, as expected
}