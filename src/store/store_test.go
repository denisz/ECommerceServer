package store

import (
	"testing"
	"time"
	"github.com/asdine/storm"
	"log"
	"github.com/asdine/storm/q"
	"fmt"
)

func TestMainPackage(t *testing.T) {

}

func TestSearchUsersByIds(t *testing.T) {
	db, err := storm.Open("my.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := db.From("store")
	var filter []User
	query := store.Select(q.In("ID", []int{1, 2}))
	err = query.Find(&filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", filter)
}

func TestFindUserByManyField(t *testing.T) {
	db, err := storm.Open("my.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := db.From("store")
	var find User
	query := store.Select(
		q.Eq("Email", "john@provider.com"),
		q.Eq("Name", "John2"),
	)
	query.Find(&find)
	fmt.Printf("%v", find)
}

func TestGeneratePassword(t *testing.T) {
	fmt.Println("6 alphabets with 2 digits : ", HumanReadablePassword(6, 2)) // best option
	fmt.Println("3 alphabets with 8 digits : ", HumanReadablePassword(3, 8))
	fmt.Println("9 alphabets with 9 digits : ", HumanReadablePassword(9, 9))
}

func TestCreateUser(t *testing.T) {
	db, err := storm.Open("my.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := db.From("store")
	user := User{
		Group: "staff",
		Email: "john@provider.com",
		Name: "John",
		CreatedAt: time.Now(),
	}

	err = store.Save(&user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", user)
}

func TestFindAllProducts(t *testing.T) {
	db, err := storm.Open("my.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	store := db.From("store")

	var products []Product
	err = store.All(&products)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", products)
}

func TestChangeProduct(t *testing.T) {
	db, err := storm.Open("my.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	store := db.From("store")

	product := Product{
		ID: 5,
		Name:"Test Product2",
		CreatedAt: time.Now(),
	}
	store.Save(&product)
}

func TestDrop(t *testing.T) {
	db, err := storm.Open("my.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Drop("store")
}
