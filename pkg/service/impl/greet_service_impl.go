package impl

// GreetService example service.
type GreetService interface {
	Say(input string) (string, error)
	Delete(s string) error
}
