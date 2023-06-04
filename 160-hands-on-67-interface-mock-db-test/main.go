package main

import (
	"fmt"
	"log"
)

// User represents a user with an id and first name.
type User struct {
	ID    int
	First string
}

// MockDatastore is a temporary service that stores retrievable data.
type MockDatastore struct {
	Users map[int]User
}

func (md MockDatastore) GetUser(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("User %v not found", id)
	}
	return user, nil
}

func (md MockDatastore) SaveUser(u User) error {
	_, ok := md.Users[u.ID]
	if ok {
		return fmt.Errorf("User %v already exists", u.ID)
	}
	md.Users[u.ID] = u
	return nil
}

// Datastore defines an interface for storing retrievable data.
// Any TYPE that implements this interface (has these two methods) is also of TYPE `Datastore`.
// This means any value of TYPE `MockDatastore` is also of TYPE `Datastore`
// This means we could have a value of TYPE `*sql.DB` and it can also be of TYPE `Datastore`
// This means we can write functions to take TYPE `Datastore` and use either of these:
// -- `MockDatastore`
// -- `*sql.DB`
// https://pkg.go.dev/database/sql#Open
type Datastore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

// Service represents a service that stores retrievable data.
// We will attach methods to `Service` so that we can use either of these:
// -- `MockDatastore`
// -- `*sql.DB`
type Service struct {
	ds Datastore
}

func (s Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}

func (s Service) SaveUser(u User) error {
	return s.ds.SaveUser(u)
}

func main() {
	db := MockDatastore{
		Users: make(map[int]User),
	}

	srvc := Service{
		ds: db,
	}

	u1 := User{
		ID:    1,
		First: "James",
	}

	err := srvc.SaveUser(u1)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	u1Returned, err := srvc.GetUser(u1.ID)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	fmt.Println(u1)
	fmt.Println(u1Returned)
}

/*
Pointers are used in method receivers in Go for two main reasons:

1. Efficiency: When a large struct is passed by value to a method (i.e., without using a pointer), the entire struct must be copied, which can be inefficient in terms of memory and performance. When a pointer to the struct is passed instead, only the memory address of the struct is passed, which is typically much more efficient.

2. Modification: If you want your method to modify the state of the receiver, the receiver must be a pointer. This is because Go is a pass-by-value language, meaning that when a value is passed to a function, a copy of that value is made in the function's scope. If the function modifies the copy, the original value is not affected. But if a pointer to the value is passed, the function can modify the original value.

In the case of your code, both `Service` and `MockDatastore` methods are defined with pointer receivers because the methods may need to modify the state of the receiver. For example, `MockDatastore.SaveUser` needs to add a user to the `Users` map of the `MockDatastore`. If `md` were not a pointer, `SaveUser` would receive a copy of the `MockDatastore`, and changes would only be made to that copy, not the original `MockDatastore`.

```go
func (md *MockDatastore) SaveUser(u User) error {
    _, ok := md.Users[u.ID]
    if ok {
        return fmt.Errorf("User %v already exists", u.ID)
    }
    md.Users[u.ID] = u
    return nil
}
```

Also, note that even when methods only read the receiver's data, you might still want to use a pointer receiver for efficiency reasons, especially if the struct is large.

*/

/*
A map is just a hash table.
The data is arranged into an array of buckets.
...
When the hashtable grows, we allocate a new array
of buckets twice as big. Buckets are incrementally
copied from the old bucket array to the new bucket array.
*/
// https://cs.opensource.google/go/go/+/refs/tags/go1.20.4:src/runtime/map.go
/*
// A header for a Go map.
type hmap struct {
	// Note: the format of the hmap is also encoded in cmd/compile/internal/reflectdata/reflect.go.
	// Make sure this stays in sync with the compiler's definition.
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra *mapextra // optional fields
}
*/
