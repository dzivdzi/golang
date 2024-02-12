package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

// User represents a user with an id and first name.
type User struct {
	ID    int
	First string
}

// MockDatastore is a temporary service  that stores retrievable data.
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
		return fmt.Errorf("User %v already exists!", u.ID)
	}
	md.Users[u.ID] = u
	return nil
}

// Datastore defines an interface for storing retrievable data.
// Any TYPE that implements this interface (has these two methods) is also of TYPE `Datastore`
// This means any value of TYPE `MockDatastore` is also of TYPE `Datastore`
// This means we could have a value of TYPE `*sql.db` and it can also be of TYPE `Datastore`
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

	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	s := rString("Yes")
	fmt.Println(s)
	rrd := returnero()
	fmt.Println(rrd())

	fmt.Println(printSquare(square, 3)) // -> You pass the function just as a name directly in the function which takes the specified func type in order to work

	timeFunc(doWork)

}

func rString(text string) string {
	return text + " I am!"
}

func returnero() func() int {
	return func() int {
		return 42
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func square(n int) int {
	return n * n
}

func printSquare(f func(int) int, a int) string {
	x := f(a)
	return fmt.Sprintf("The number %v squared is %v", a, x)
}

func powinator(a float64) func() float64 {
	c := 0.0 // this is the power incrementor - 1,2,3,4,5,6 etc ...
	return func() float64 {
		c++
		return math.Pow(a, c)
	}
}

func doWork() {
	for i := 0; i < 2_000; i++ {
		fmt.Println(i)
	}
}

func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
