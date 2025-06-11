package shape

//TODO: create a Circle struct, that has field Radius of type float64

type Circle struct {
	Radius float64
}

//Implement Shape interface methods

func(c *Circle)Area() float64{
	return c.Radius * 3.14 * c.Radius
}


func(c *Circle)Parameter() float64{
	return c.Radius * 3.14 * 2
}


func(c *Circle)Resize(resizer float64){
	c.Radius = c.Radius * resizer
}