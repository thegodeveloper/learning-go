package structs

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

// Individual 1. Basic Struct Definition
type Individual struct {
	Name string
	Age  int
	City string
}

// Product 2. Struct with Different Field Types
type Product struct {
	ID        int
	Name      string
	Price     float64
	InStock   bool
	Tags      []string
	CreatedAt time.Time
}

// Address 3. Struct with Embedded Fields (Composition)
type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

type Personnel struct {
	Individual // Embedded struct (anonymous field)
	Address    // Embedded struct
	ID         int
	Department string
	Salary     float64
	IsManager  bool
}

// Rectangle 4. Struct with Methods (Receiver Functions)
type Rectangle struct {
	width  float64
	height float64
}

// Area Value receiver - doesn't modify the struct
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Perimeter Value receiver
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Note: I am aware of not to mix value and pointer receivers

// Scale Pointer receiver - can modify the struct
func (r *Rectangle) Scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

// SetDimensions Pointer receiver
func (r *Rectangle) SetDimensions(width, height float64) {
	r.width = width
	r.height = height
}

// BankAccount 5. Struct with Private and Public Fields
type BankAccount struct {
	AccountNumber string  // Public (exported)
	balance       float64 // Private (unexported)
	ownerName     string  // Private
}

// Shape 6. Interface Implementation
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// User 7. Struct with Tags (for JSON, database mapping, etc.)
type User struct {
	ID       int    `json:"id" db:"user_id" validate:"required"`
	Username string `json:"username" db:"username" validate:"required,min=3"`
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"-" db:"password_hash"`               // "-" excludes from JSON
	FullName string `json:"full_name,omitempty" db:"full_name"` // omitempty excludes if empty
}

// NewBankAccount Constructor function (common Go pattern)
func NewBankAccount(accountNumber, ownerName string, initialBalance float64) *BankAccount {
	return &BankAccount{
		AccountNumber: accountNumber,
		balance:       initialBalance,
		ownerName:     ownerName,
	}
}

// Config 8. Anonymous Structs
type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"database"`
	Server struct {
		Port int    `json:"port"`
		Host string `json:"host"`
	} `json:"server"`
}

// Point 9. Struct Comparison and Copying
type Point struct {
	X, Y int
}

// Balance Getter method for private field
func (b *BankAccount) Balance() float64 {
	return b.balance
}

// Deposit Setter method for private field with validation
func (b *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("cannot deposit zero or negative amount")
	}
	b.balance += amount
	return nil
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("cannot withdraw zero or negative amount")
	}
	if amount > b.balance {
		return fmt.Errorf("insufficient balance to withdraw")
	}
	b.balance -= amount
	return nil
}

func Basics(show bool) {
	if show {
		fmt.Println("--- Go Structures Comprehensive Example ---")

		// 1. Basic Struct Usage
		fmt.Println("1. Basic Struct Declaration and Initialization")

		// Different ways to create structs
		var p1 Individual                         // Zero value initialization
		p2 := Individual{}                        // Empty initialization
		p3 := Individual{"Alice", 30, "New York"} // Positional initialization
		p4 := Individual{
			Name: "Bob",
			Age:  30,
			City: "New York",
		}

		fmt.Printf("p1 (zero value): %+v\n", p1)
		fmt.Printf("p2 (empty): %+v\n", p2)
		fmt.Printf("p3 (positional): %+v\n", p3)
		fmt.Printf("p4 (named): %+v\n", p4)
		fmt.Println()

		// 2. Accessing and Modifying Fields
		fmt.Println("2. Accessing and Modifying Fields:")
		p1.Name = "Charlie"
		p1.Age = 35
		p1.City = "Chicago"

		fmt.Printf("Modified p1: %+v\n", p1)
		fmt.Printf("p1.Name: %s, p1.Age: %d\n", p1.Name, p1.Age)
		fmt.Println()

		// 3. Struct Pointers
		fmt.Println("3. Struct Pointers:")

		p5 := &Individual{"David", 40, "Denver"} // Pointer to struct
		fmt.Printf("p5 pointer: %p\n", p5)
		fmt.Printf("p5 value: %+v\n", *p5)

		// Go automatically dereferences struct pointers
		fmt.Printf("p5.Name (auto-dereference): %s\n", p5.Name)

		// Explicit dereferencing - not recommended but available
		fmt.Printf("(*p5).Name (explicit): %s\n", (*p5).Name)

		// Modifying through pointer
		p5.Age = 41
		fmt.Printf("Modified through pointer: %+v\n", *p5)

		// 4. Embedded Structs (Composition)
		fmt.Println("4. Embedded Structs (Composition):")

		per := Personnel{
			Individual: Individual{
				Name: "Emily",
				Age:  27,
				City: "Seattle",
			},
			Address: Address{
				Street:  "123 Main St",
				City:    "Seattle",
				State:   "WA",
				ZipCode: "98101",
			},
			ID:         1001,
			Department: "Engineering",
			Salary:     75000,
			IsManager:  false,
		}
		fmt.Printf("Personnel: %+v\n", per)

		// Accessing embedded fields directly
		fmt.Printf("Personnel Name: %s\n", per.Name)
		fmt.Printf("Personnel Street: %s\n", per.Street)
		fmt.Printf("Personnel Department: %s\n", per.Department)

		// Accessing embedded fields explicitly
		fmt.Printf("Individual Info: %+v\n", per.Individual)
		fmt.Printf("Address Info: %+v\n", per.Address)

		// 5. Methods on Structs
		fmt.Println("5. Methods on Structs:")

		rect := Rectangle{width: 10, height: 5}
		fmt.Printf("Rectangle: %+v\n", rect)
		fmt.Printf("Area: %.2f\n", rect.Area())
		fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())

		// Method with pointer receiver
		fmt.Printf("Before scaling: %+v\n", rect)
		rect.Scale(2)
		fmt.Printf("After scaling by 2: %+v\n", rect)

		rect.SetDimensions(15, 8)
		fmt.Printf("After setting new dimensions: %+v\n", rect)
		fmt.Println()

		// 6. Private Fields and Constructor Pattern
		fmt.Println("6. Private Fields and Constructor Pattern:")

		account := NewBankAccount("ACC-001", "John Doe", 1000.0)
		fmt.Printf("New Account: %+v\n", account)
		fmt.Printf("Initial Balance: %.2f\n", account.Balance())

		// Deposit money
		err := account.Deposit(500)
		if err != nil {
			fmt.Printf("Deposit Error: %v\n", err)
		} else {
			fmt.Printf("After Deposit: $%.2f\n", account.Balance())
		}

		// Withdraw money
		err = account.Withdraw(300)
		if err != nil {
			fmt.Printf("Withdrawal Error: %v\n", err)
		} else {
			fmt.Printf("After Withdrawal: $%.2f\n", account.Balance())
		}

		// Try invalid operation
		err = account.Withdraw(2000)
		if err != nil {
			fmt.Printf("Invalid Withdrawal: %v\n", err)
		}
		fmt.Println()

		// 7. Interface Implementation
		fmt.Println("7. Interface Implementation:")
		shapes := []Shape{
			Rectangle{width: 10, height: 5},
			Circle{Radius: 3},
		}

		for i, shape := range shapes {
			fmt.Printf("Shape %d:\n", i+1)
			fmt.Printf("  Type: %T\n", shape)
			fmt.Printf("  Area: %.2f\n", shape.Area())
			fmt.Printf("  Perimeter: %.2f\n", shape.Perimeter())
		}
		fmt.Println()

		// 8. Struct Tags and JSON
		fmt.Println("8. Struct Tag and JSON:")

		user := User{
			ID:       1,
			Username: "johndoe",
			Email:    "john@example.com",
			Password: "secret123",
			FullName: "John Doe",
		}

		// Convert to JSON
		jsonData, err := json.Marshal(user)
		if err != nil {
			fmt.Printf("JSON marshal error: %v\n", err)
		} else {
			fmt.Println("JSON representation %s\n", string(jsonData))
		}

		// Reflect on struct tags
		userType := reflect.TypeOf(user)
		for i := 0; i < userType.NumField(); i++ {
			field := userType.Field(i)
			jsonTag := field.Tag.Get("json")
			dbTag := field.Tag.Get("db")
			fmt.Printf("Field %s: json='%s', db='%s'\n", field.Name, jsonTag, dbTag)
		}
		fmt.Println()

		// 9. Anonymous Structs
		fmt.Println("9. Anonymous Structs:")

		// Inline anonymous struct
		settings := struct {
			Theme    string
			Language string
			Debug    bool
		}{
			Theme:    "dark",
			Language: "en",
			Debug:    true,
		}
		fmt.Printf("Settings: %+v\n", settings)

		// Anonymous struct in slice
		people := []struct {
			Name string
			Role string
		}{
			{"Alice", "Developer"},
			{"Bob", "Designer"},
			{"Charlie", "Manager"},
		}

		fmt.Println("People:")
		for _, person := range people {
			fmt.Printf("  %s - %s\n", person.Name, person.Role)
		}
		fmt.Println()

		// 10. Nested Anonymous Structs
		fmt.Println("10. Nested Anonymous Structs:")

		config := Config{}
		config.Database.Host = "localhost"
		config.Database.Port = 5432
		config.Database.Username = "admin"
		config.Database.Password = "secret"
		config.Server.Host = "0.0.0.0"
		config.Server.Port = 8080
		fmt.Printf("Config: %+v\n", config)
		fmt.Println()

		// 11. Struct Comparison and Copying
		fmt.Println("11. Struct Comparison and Copying:")

		point1 := Point{X: 1, Y: 2}
		point2 := Point{X: 1, Y: 2}
		point3 := Point{X: 3, Y: 4}
		fmt.Printf("point1: %+v\n", point1)
		fmt.Printf("point2: %+v\n", point2)
		fmt.Printf("point3: %+v\n", point3)
		fmt.Printf("point1 == point2: %t\n", point1 == point2)
		fmt.Printf("point1 == point3: %t\n", point1 == point3)

		// Copying structs (value semantics)
		point4 := point1 // Creates a copy
		point4.X = 10
		fmt.Printf("Original point1: %+v\n", point1)
		fmt.Printf("Modified copy point4: %+v\n", point4)
		fmt.Println()

		// 12. Struct with Slice and Map Fields
		fmt.Println("12. Struct with Complex Fields:")

		product := Product{
			ID:        101,
			Name:      "Laptop",
			Price:     999.99,
			InStock:   true,
			Tags:      []string{"electronics", "computer", "portable"},
			CreatedAt: time.Now(),
		}

		fmt.Printf("Product: %+v\n", product)
		fmt.Printf("Tags: %s\n", strings.Join(product.Tags, ", "))
		fmt.Printf("Created: %s\n", product.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Println()

		// 13. Empty Struct Usage
		fmt.Println("13. Empty Struct Usage:")
		// Empty struct as a signal
		type Signal struct{}

		// Using empty struct as set values (memory efficient)
		visited := make(map[string]struct{})
		items := []string{"apple", "banana", "apple", "cherry", "banana"}

		for _, item := range items {
			visited[item] = struct{}{}
		}

		fmt.Printf("Original items: %v\n", items)
		fmt.Printf("Unique items: ")
		for item := range visited {
			fmt.Printf("%s", item)
		}
		fmt.Printf("\nEmpty struct size: %d bytes\n", unsafe.Sizeof(struct{}{}))
		fmt.Println()

		// 14. Method Sets and Value vs Pointer Receivers
		fmt.Println("14. Method Sets and Value vs Pointer Receivers:")
		rect1 := Rectangle{width: 5, height: 3}
		rect2 := &Rectangle{width: 7, height: 4}
		// Both value and pointer can call value receiver methods
		fmt.Printf("rect1.Area(): %.2f\n", rect1.Area())
		fmt.Printf("rect2.Area(): %.2f\n", rect2.Area())

		// Both value and pointer can call pointer receiver methods
		// (Go automatically takes address of value or dereferences pointer)
		rect1.Scale(1.5)
		rect2.Scale(2.0)
		fmt.Printf("After scaling - rect1: %+v\n", rect1)
		fmt.Printf("After scaling - rect1: %+v\n", *rect2)
	}
}
