package main

import (
	"fmt"
	"strings"
)

func securityCheck(in string) error {
	in = strings.Replace(in, "\r\n", "\n", -1)

	c := strings.Split(in, anchorStr)
	if len(c) != 2 {
		return fmt.Errorf("don't modify anchor")
	}

	upper, lower := c[0], c[1]

	// dont allow multi statement in one line in upper half
	if strings.Contains(upper, ";") {
		return fmt.Errorf("dont allow multi statement in one line in upper half")
	}

	// upper half must starts with 'from diagrams'
	uls := strings.Split(upper, "\n")
	for _, ul := range uls {
		ul = strings.TrimSpace(ul)
		if ul == "" {
			continue
		}
		if strings.HasPrefix(ul, "#") {
			continue
		}
		if !strings.HasPrefix(ul, "from diagrams") {
			return fmt.Errorf("only allow import from diagrams: %s", ul)
		}
	}

	// don't allow import in lower half
	if strings.Contains(lower, "import") {
		return fmt.Errorf("import should be placed in above of the anchor")
	}

	return nil
}
