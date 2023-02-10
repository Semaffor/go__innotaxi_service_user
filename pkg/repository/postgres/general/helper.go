package general

import (
	"fmt"
	"strings"
)

func generateDollarSequence(paramsCount int) string {
	args := make([]string, 0)
	for i := 1; i <= paramsCount; i++ {
		args = append(args, fmt.Sprintf("$%d", i))
	}

	return strings.Join(args, ", ")
}
