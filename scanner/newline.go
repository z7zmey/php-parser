package scanner

type NewLines struct {
	data []int
}

func (nl *NewLines) Append(p int) {
	if len(nl.data) == 0 || nl.data[len(nl.data)-1] < p {
		nl.data = append(nl.data, p)
	}
}

func (nl *NewLines) GetLine(p int) int {
	line := len(nl.data) + 1

	for i := len(nl.data) - 1; i >= 0; i-- {
		if p < nl.data[i] {
			line = i + 1
		} else {
			break
		}
	}

	return line
}
