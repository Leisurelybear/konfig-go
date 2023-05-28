package messagequeue

var (
	store   MessageStore
	Produce IProducer
)

func Init() {
	store = NewInMemoryMessageStore()
	Produce = NewProducer(store)
}
