package worklist

type QueryOutput struct {
	Output string
}

type Worklist struct {
	jobs chan QueryOutput
}

func (w *Worklist) Add(work QueryOutput) {
	w.jobs <- work
}

func (w *Worklist) Next() QueryOutput {
	j := <-w.jobs
	return j
}

func New(bufSize int) Worklist {
	return Worklist{make(chan QueryOutput, bufSize)}
}

func NewJob(content string) QueryOutput {
	return QueryOutput{content}
}

func (w *Worklist) Finalize(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		w.Add(QueryOutput{""})
	}
}
