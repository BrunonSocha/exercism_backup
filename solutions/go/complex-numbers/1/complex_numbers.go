package complexnumbers
import "math"
// Define the Number type here.
type Number struct {
	real float64
	imaginary float64
}

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.imaginary
}

func (n1 Number) Add(n2 Number) Number {
	return Number{n1.real+n2.real, n1.imaginary+n2.imaginary}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.real-n2.real, n1.imaginary-n2.imaginary}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{(n1.real * n2.real - n1.imaginary * n2.imaginary), (n1.imaginary * n2.real + n1.real * n2.imaginary)}
}

func (n Number) Times(factor float64) Number {
	return Number{(factor)*(n.real), (factor)*(n.imaginary)}
}

func (n1 Number) Divide(n2 Number) Number {
	return Number{(n1.real * n2.real + n1.imaginary * n2.imaginary) / (math.Pow(n2.real, 2) + math.Pow(n2.imaginary, 2)), (n1.imaginary * n2.real - n1.real * n2.imaginary) / (math.Pow(n2.real, 2) + math.Pow(n2.imaginary, 2))}
}

func (n Number) Conjugate() Number {
	return Number{n.real, -1 * n.imaginary}
}

func (n Number) Abs() float64 {
	return math.Sqrt((math.Pow(n.real, 2) + math.Pow(n.imaginary, 2)))
}

func (n Number) Exp() Number {
	return Number{math.Pow(math.E, n.real) * math.Cos(n.imaginary), math.Pow(math.E, n.real) * math.Sin(n.imaginary)}
}
