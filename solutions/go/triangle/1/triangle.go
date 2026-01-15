package triangle

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

func KindFromSides(a, b, c float64) Kind {
	isTriangle := a+b >= c && b+c >= a && a+c >= b
	allSidesZero := a == 0 && b == 0 && c == 0

	if !isTriangle || allSidesZero {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}
