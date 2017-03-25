package handlers

import (
	"net/http"
	"strings"
)

type PuzzleHandler struct {
}

func NewPuzzleHandler() *PuzzleHandler {
	return &PuzzleHandler{}
}

func (h *PuzzleHandler) GetAnswer(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	ReplyOk(w, h.Calculate(query))
}

func (h *PuzzleHandler) Calculate(query string) string {
	result := []string{query}
	queryLength := len(query)
	queryChars := []rune(query)

	i := 0
	for queryLength > i {
		//		fmt.Printf("checking > %d of %d : %c of %s\r\n", i, queryLength, queryChars[i], string(queryChars))
		if queryChars[i] == 'a' || queryChars[i] == 'i' || queryChars[i] == 'u' || queryChars[i] == 'e' || queryChars[i] == 'o' {
			queryChars = append(queryChars[:i], queryChars[i+1:queryLength]...) // delete this character from array
			queryLength--

			result = append(result, string(queryChars))
			continue
		}
		i++
	}

	return strings.Join(result, "-")
}
