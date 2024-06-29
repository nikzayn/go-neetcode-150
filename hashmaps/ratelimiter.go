package hashmaps

type Logger struct {
	requests map[string]int
	limit    int
}

func (m *Logger) requestLoggerInit(timeLimit int) {
	m.requests = make(map[string]int)
	m.limit = timeLimit
}

func (m *Logger) messageRequestDescision(request string, timestamp int) bool {
	for index, _ := range m.requests {
		if index == request && timestamp-m.requests[request] >= m.limit {
			m.requests[request] = timestamp
			return true
		} else if index == request {
			return false
		}
	}
	m.requests[request] = timestamp
	return true
}
