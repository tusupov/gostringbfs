package list

type list struct {
	items map[string]string
}

func New() list {
	return list{
		items: make(map[string]string),
	}
}

func (s *list) Put(key, value string) {
	s.items[key] = value
}

func (s list) FindText(needText string) ([]string, bool) {

	n := len(needText)
	if n == 0 {
		return nil, false
	}

	used := make([]bool, n+1)
	path := make([]int, n+1)
	code := make([]string, n+1)

	queue := make(chan int, n)
	defer close(queue)

	used[0] = true
	queue <- 0

LOOP:
	for {

		select {

		case i := <-queue:

			for j := n; j >= i+1; j-- {
				if used[j] {
					continue
				}
				if c, ok := s.items[needText[i:j]]; ok {
					path[j] = i
					code[j] = c
					if j == n {
						break LOOP
					}
					used[j] = true
					queue <- j
				}
			}

		default:
			break LOOP

		}

	}

	if path[n] == 0 {
		return nil, false
	}

	// Get results
	result := make([]string, 0)
	pos := n
	for pos != 0 {
		result = append(result, code[pos])
		pos = path[pos]
	}

	// Reverse result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result, true

}
