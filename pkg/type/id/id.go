package id

// ID is a type for the uint underlying type with phantom type parameter that
// can be used to store a unique identifier for an item of type T.
// It is typically used to store a unique identifier for an item in a data structure or database.
// It is important to note that the ID should be unique and not reused for different items.
//
// Example:
//
//	type Item struct {
//		ID   ID[Item]
//		Name string
//	}
type ID[T any] uint
