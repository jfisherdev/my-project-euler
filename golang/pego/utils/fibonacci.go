package pego

type FibonacciSequence struct {
	Sequence []int
}

func (f *FibonacciSequence) Next() int {
	curTerm := len(f.Sequence)
	var next int
	if curTerm == 0 {
		next = 0
	} else if curTerm == 1 {
		next = 1
	} else {
		next = f.Sequence[curTerm-1] + f.Sequence[curTerm-2]
	}
	f.Sequence = append(f.Sequence, next)
	return next
}
