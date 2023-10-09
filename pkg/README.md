# Packages Directory

The `pkg` directory contains reusable packages and libraries that can be used across our Golang-based monorepo project.
These packages are designed to provide common functionality and utilities that can be shared among different modules of
the project.

## Directory Structure

- **pkg1**: Description of package 1 and its purpose.
    - `pkg1.go`: Source code file for package 1.
    - `README.md`: Detailed explanation and usage instructions for package 1.

- **pkg2**: Description of package 2 and its purpose.
    - `pkg2.go`: Source code file for package 2.
    - `README.md`: Detailed explanation and usage instructions for package 2.

- **pkg3**: Description of package 3 and its purpose.
    - `pkg3.go`: Source code file for package 3.
    - `README.md`: Detailed explanation and usage instructions for package 3.

## Usage

To use a specific package, import it in your Go code as you would with any other package. Refer to the README.md file in
the respective package directory for detailed usage instructions and examples.

Example of importing `pkg1`:

```go
package main

import (
	"github.com/your-repo/pkg/pkg1"
)
```
