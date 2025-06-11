package shape

//TODO: create a Rectangle struct, that has a field Height of type float64
//  and a field Width of type float64

//Implement Shape interface methods



type Rectangle struct{
	Height float64
	Width float64
}


func(c *Rectangle)Area() float64{
	return c.Height * c.Width
}


func(c *Rectangle)Parameter() float64{
	return (c.Height + c.Width) *2
}


func(c *Rectangle)Resize(resizer float64){
	c.Height = c.Height *resizer
	c.Width = c.Width *resizer
}