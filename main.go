package main

import "fmt"

type Person struct {
	name string
}

func (person *Person) getName() {
	fmt.Println("Person Name : ", person.name)
}

func main() {
	person := Person{"John Doe"}
	person.getName()

	job := Job{"123", "Software Development"}
	job.getName()
}

// the interface method(s) should be implemented, we can add the interface later on also, we don't
// have to mention the interface name or extends interface in the Person or Job structs

type CommonInterface interface {
	getName()
}

type Job struct {
	id   string
	name string
}

func (job *Job) getName() {
	fmt.Printf("Name of job is : %s, ID is: %s \n", job.name, job.id)
}
