package gopher

//Gopher encapsultates all field and behaviours of a gopher!
type Gopher struct{
	Name string
	favFood string
}

//New creates an instance of a gopher struct.
func New (name string) Gopher{
	return Gopher{name, "crisps"}
}

//FavFood returns Gopher struct favFood field
func (g Gopher) FavFood() string{
	return g.favFood
}
