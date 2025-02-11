# Go Silent Map Access

This repo demonstrates a common but easily overlooked error in Go: accessing a non-existent key in a map.

Accessing a non-existent map key will not result in a runtime panic.  Instead, the zero value of the map's value type is returned. This can be problematic because the absence of an error makes it difficult to detect this type of bug.

**bug.go** shows the problematic code.
**bugSolution.go** provides solutions to this problem.