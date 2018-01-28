package main

import "fmt"

func main2() {
	p1 := user{
		name:  "samson",
		age:   22,
		email: "sam@test.com",
	}

	p2 := user{"abebe", 44, "abe@test.com"}
	fmt.Println(p1, "\n", p2)

	ad := admin{
		person: user{
			name:  "Balca",
			age:   55,
			email: "balcha@test.com",
		},
		level: "super",
	}
	fmt.Println(ad)

	var dur Duration
	dur = (Duration)(int64(1000)) // No implicit type conversions
	fmt.Println(dur)

	// Methods
	bill := user{"Bill", 23, "bill@email.com"}
	lisa := &user{"Lisa", 31, "lisa@email.com"}

	bill.notify()
	lisa.notify()

	bill.changeEmail("bill@comcast.com")
	bill.notify()

	lisa.changeEmail("lisa@comcast.com")
	lisa.notify()

	// Test
	str := ""
	fmt.Println([]byte(str))

}

type user struct {
	name  string
	age   int
	email string
}

type admin struct {
	person user
	level  string
}

type Duration int64

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}
