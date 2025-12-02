package triangle

type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT = "not a triangle"
	Equ = "equilateral"
	Iso = "isosceles"
	Sca = "scalene"
)

func KindFromSides(a, b, c float64) Kind {
	var k Kind

    if a != b && a != c && b != a && b != c {
        k = Sca
    }else if a == b && b == c {
        k = Equ
    }else if a == b || b == c || a == c{
        k = Iso
    }

    if a + b < c || b + c < a || c + a < b {
        k = NaT
    }

    if a <= 0 || b <= 0 || c <= 0 {
        k = NaT
    }
    
	return k
}