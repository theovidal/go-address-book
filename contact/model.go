package contact

type Contact struct {
	Name string `csv:"name"`
	Email string `csv:"email"`
	Address string `csv:"address"`
	Phone string `csv:"phone"`
}
