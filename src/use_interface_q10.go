package main

import "fmt"

type input_port interface {
	run() error
	do_run()
}

type output_port interface {
	run(i input_Person)
}

type input_Person struct {
	Name string
}

type output_Person struct {
	Name string
}

type MyError struct {
	code string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error: %s\n", e.code)
}

func (i input_Person) run() error {
	if i.Name == "" {
		return &MyError{"invalid value"}
	}
	var o output_port = output_Person{Name: i.Name}
	o.run(i)
	return nil
}

func (i input_Person) do_run() {
	fmt.Printf("%s\n", i.Name)
}

func (o output_Person) run(i input_Person) {
	i.do_run()
}

func main() {
	var i input_port = input_Person{""}
	if err := i.run(); err != nil {
		fmt.Println(err)
	}
}
