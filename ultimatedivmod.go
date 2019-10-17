package piscine

func UltimateDivMod(a *int, b *int) {
	y := *a
	x := *b
	a = y / x
	b = y % x
}
