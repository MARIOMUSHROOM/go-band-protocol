package ports

type ActionService interface {
	CheckRevenge(text string) string
	MaxChickensProtected(n, k int, positions []int) int
}
