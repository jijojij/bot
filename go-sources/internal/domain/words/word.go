package words

type Word struct {
	id          int
	origin      string
	translation string
}

func NewWord(id int, origin string, translation string) *Word {
	return &Word{
		id:          id,
		origin:      origin,
		translation: translation,
	}
}

func (w *Word) ID() int {
	return w.id
}

func (w *Word) Origin() string {
	return w.origin
}

func (w *Word) Translation() string {
	return w.translation
}

func (w *Word) String() string {
	return w.origin + " - " + w.translation
}
