# Domain-Driven Design (DDD) Entities

The `entity` directory contains domain-driven design (DDD) entities for our Golang-based monorepo project. Entities
represent the core domain objects that capture essential business concepts and have a distinct identity.

## Directory Structure

- **user**: User entity representing a system user.
    - `user.go`: User entity definition.
    - `README.md`: Detailed explanation of the User entity and its attributes.

- **product**: Product entity representing a product in our application.
    - `product.go`: Product entity definition.
    - `README.md`: Detailed explanation of the Product entity and its attributes.

- **order**: Order entity representing an order placed by a user.
    - `order.go`: Order entity definition.
    - `README.md`: Detailed explanation of the Order entity and its attributes.

- **...**: Other domain entities, if applicable.

## Usage

Entities within the `entity` directory are the building blocks of our domain-driven design. They encapsulate domain
logic and attributes specific to various concepts within our application's domain.

To use an entity, import its package and create instances in your application code.

Example for the User entity:

```go
package main

import (
	"github.com/your-repo/entity/user"
)

func main() {
	newUser := user.NewUser("John Doe", "john@example.com")

	// Use the User entity
	userID := newUser.ID()
	userName := newUser.Name()

	// Perform operations on the User entity
	newUser.UpdateEmail("newemail@example.com")

	// ...
}
```
