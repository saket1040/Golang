package main

import "fmt"

// we have invoker, receiver and command
// invoker - it is a remote control
// receiver - it is a light
// command - it is a command to turn on or off the light
// a comand will always have an undo functionality

// Command interface
type Command interface {
	Execute()
}

// Receiver - it a dump object
type Light struct{}

func (l *Light) On() {
	fmt.Println("Light is ON")
}

func (l *Light) Off() {
	fmt.Println("Light is OFF")
}

// Concrete Command: Turns the light ON
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

// Concrete Command: Turns the light OFF
type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

// Invoker
type RemoteControl struct {
	commandMap map[string]Command
}

func NewRemoteControl() *RemoteControl {
	return &RemoteControl{commandMap: make(map[string]Command)}
}

func (r *RemoteControl) RegisterCommand(action string, cmd Command) {
	r.commandMap[action] = cmd
}

func (r *RemoteControl) HandleInput(action string) {
	if cmd, ok := r.commandMap[action]; ok {
		cmd.Execute()
	} else {
		fmt.Println("Unknown command:", action)
	}
}

// Main
func main() {
	light := &Light{}

	remote := NewRemoteControl()
	remote.RegisterCommand("turn_on", &LightOnCommand{light})
	remote.RegisterCommand("turn_off", &LightOffCommand{light})

	// Simulating string-based input (e.g., from user or API)
	inputs := []string{"turn_on", "turn_off", "toggle"}

	for _, input := range inputs {
		remote.HandleInput(input)
	}
}