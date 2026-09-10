package piscine

func UltimateDivMod(a *int, b *int) {
	mod := *a % *b
	*a = *a / *b
	*b = mod
}
