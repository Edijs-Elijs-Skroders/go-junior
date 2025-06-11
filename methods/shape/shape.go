package shape

//TODO: add method Area() to the interface, which returns float64 
// when implementing the interface, this should return the area of the shape
//TODO: add method Parameter() to the interface, which returns float64 
//when implementing interface, this should return the parameter of the shape
//TODO: add method Resize() to the interface, which takes type float64 paramter named "resizer"
//when implementing interface, this should resize the shapes values, e.g. shape.radius = shape.radius * resizer
type Shape interface{
	Area() float64
	Parameter() float64
	Resize(resizer float64)
}


//TODO: implement a function ResizeAll, that has a parameter of type Shape slice,
//and a parameter resizer of type float64
//This function should call Resize() on all of the shapes in the slice

func ResizeAll(slc []Shape, resizer float64){
	for _, v := range slc {
		v.Resize(resizer)
	}
}