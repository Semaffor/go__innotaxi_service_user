package general

import (
	"fmt"
	"strings"
)

// generateDollarSequence helps to create string in the following way: '$1, $2, ..., $n'
// for safe query injection.
func generateDollarSequence(paramsCount int) string {
	args := make([]string, 0)
	for i := 1; i <= paramsCount; i++ {
		args = append(args, fmt.Sprintf("$%d", i))
	}

	return strings.Join(args, ", ")
}
