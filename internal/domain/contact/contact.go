package contact

type Contact struct{
	Mail string `validate:"required,email"`
}