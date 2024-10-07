package words

type Repository interface {
	Save(word *Word) error
}
