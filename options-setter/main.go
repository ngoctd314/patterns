package main

// Person ...
type Person struct {
	firstName string
	lastName  string
	age       int
	city      string
	country   string
	gender    string
	edu       string
	hobbies   []string
}

// Option ...
type Option func(*Person)

// NewPerson ...
func NewPerson(options ...Option) *Person {
	person := new(Person)
	for _, option := range options {
		option(person)
	}
	return person
}

// WithFirstName ...
func WithFirstName(firstName string) Option {
	return func(p *Person) {
		p.firstName = firstName
	}
}

// WithLastName ...
func WithLastName(lastName string) Option {
	return func(p *Person) {
		p.lastName = lastName
	}
}

// WithAge ..
func WithAge(age int) Option {
	return func(p *Person) {
		p.age = age
	}
}

// WithCity ...
func WithCity(city string) Option {
	return func(p *Person) {
		p.city = city
	}
}
