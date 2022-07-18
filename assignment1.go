package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Library []string

//Building a Function to add New members and store in a text file
func (*Library) addNewMember() {
	//flag to display Welcome Message
	Welcome := flag.String("Newmember:", "Welcome to Adding New Member Section", "specify your book name")
	flag.Parse()
	fmt.Println(*Welcome)
	//Taking names from the user
	Newmember := make(Library, 0)
	var i, j int
	//total Numbers of users to input
	fmt.Println("Enter Number of members to input:")
	fmt.Scanln(&j)
	i = 0
	scanner := bufio.NewScanner(os.Stdin)
	for i < j {
		fmt.Print("Enter New member name: ")

		scanner.Scan()

		text := scanner.Text()
		//Appending all taken names to the Slice
		if len(text) != 0 {

			fmt.Println(text)
			Newmember = append(Newmember, text)
		} else {
			break
		}
		i++
	}

	//creating a text file for storing names everytime the function runs
	file, err := os.OpenFile("RegisteredUsers.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, data := range Newmember {
		_, _ = datawriter.WriteString(data + "\n")
	}
	//error handling for flush
	er := datawriter.Flush()
	if er != nil {
		println("error")
	}
}

//Building a Function to add New book and store in a text file
func (*Library) addNewBook() {
	//Flag to display Welcome message
	Welcome := flag.String("AddBooks:", "Welcome to Adding New Book Section", "specify your book name")
	flag.Parse()
	fmt.Println(*Welcome)

	//Defining different variables and taking input from the user
	//check for checking if the book is digitalbook or physicalbook
	var check, name, author string
	var kind Booktype
	fmt.Println("enter Physical or Digital:")
	fmt.Scanln(&check)
	//Enter Book details
	fmt.Println("Enter Book Name:")
	fmt.Scanln(&name)
	fmt.Println("Enter Author Name:")
	fmt.Scanln(&author)
	fmt.Println("Enter Book Kind:")
	fmt.Scanln(&kind)
	//checking if there is input or not
	if len(check) != 0 && len(name) != 0 && len(author) != 0 {
		Newbook := DigitalBook{
			Name:   name,
			Author: author,
			Kind:   kind,
		}

		//Storing entries in a json file
		//checking if a digital book then storing in a Digitalbook.json file
		if check == "Digitalbook" || check == "digitalbook" {
			NewDigitalBook(Newbook)
		}
		//checking if a physucal book then storing in a Physicalbook.json file
		if check == "physicalbook" || check == "Physicalbook" {
			NewPhysicalBook(Newbook)
		}

	} else {
		fmt.Println("error")
	}
}
func (*Library) collected() {
	//Flag to display Welcome message
	Welcome := flag.String("returnBooks:", "Book return Section", "specify your book name")
	flag.Parse()
	fmt.Println(*Welcome)
	//Asking the name of the member returning book
	var name string
	fmt.Println("Name of the user:")
	fmt.Scanln(&name)
	//reading the file
	content, err := ioutil.ReadFile("issuedList.json")
	//checking errors
	if err != nil {
		log.Fatal(err)
	}
	//storing file data into another variable of similar type
	user2 := map[string]PhysicalBook{}

	//unmarshalling the data and printing the data
	err = json.Unmarshal(content, &user2)
	if err != nil {
		log.Fatal(err)
	}
	for key, _ := range user2 {
		if key == name {
			fmt.Println("Return Successful")

		} else {
			fmt.Println("Return failed")
		}

	}
}

// Book an interface type with Certain methods
type Book interface {
	Bookdetails()
}

// Bookdetails function to tell about the book details
// Bookdetails of digitalBooks
func (d DigitalBook) Bookdetails() {
	//reading the file
	content, err := ioutil.ReadFile("Digitalbook.json")
	//checking errors
	if err != nil {
		log.Fatal(err)
	}
	//storing file data into another variable of similar type
	user1 := map[string]DigitalBook{}
	//unmarshalling the data and printing the data
	err = json.Unmarshal(content, &user1)
	//error check
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user1)

}

// Bookdetails of Physicalbooks
func (p PhysicalBook) Bookdetails() {
	//reading the file
	content, err := ioutil.ReadFile("Physicalbook.json")
	//checking errors
	if err != nil {
		log.Fatal(err)
	}
	//storing file data into another variable of similar type
	user2 := PhysicalBook{}
	//unmarshalling the data and printing the data
	err = json.Unmarshal(content, &user2)
	//error check
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", user2)
}

// Borrow functions to borrow book for the library
func Borrow() bool {
	//Flag to display Welcome message
	Welcome := flag.String("Borrowbooks:", "Welcome to Borrow New Book Section", "specify your book name")
	flag.Parse()
	fmt.Println(*Welcome)
	//Name entered by the user to check if user registered or not
	var name string
	var bo bool
	fmt.Println("Please enter your Name")
	fmt.Scanln(&name)
	//Opening the register users File to check the name
	file, err := os.Open("RegisteredUsers.txt")
	if err != nil {
		fmt.Println("error")
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, line := range lines {
		//if the name given by the user matches with the name in the register username then user can borrow book
		if line == name {
			fmt.Printf("Welcome %s\n", name)
			break
		}
		return false
	}

	//Defining different variables and taking input from the required user
	var check, bookname, author string
	var kind Booktype
	fmt.Println("enter Physical or Digital:")
	fmt.Scanln(&check)
	//Enter Book details
	fmt.Println("Enter Book Name:")
	fmt.Scanln(&bookname)
	fmt.Println("Enter Author Name:")
	fmt.Scanln(&author)
	fmt.Println("Enter Book Kind:")
	fmt.Scanln(&kind)
	//checking if there is input or not
	if len(check) != 0 && len(name) != 0 && len(author) != 0 {
		book := PhysicalBook{
			Name:   bookname,
			Author: author,
			Kind:   kind,
		}
		//checking if book is physical or digital
		//if book is physical
		if check == "physicalbook" || check == "Physicalbook" {
			//reading the file
			content, err := ioutil.ReadFile("Physicalbook.json")
			//checking errors
			if err != nil {
				log.Fatal(err)
			}
			//storing file data into another variable of similar type
			user2 := PhysicalBook{}
			//unmarshalling the data and printing the data
			err = json.Unmarshal(content, &user2)
			if err != nil {
				log.Fatal(err)
			}
			//comparing the data of both the struct,if same then return true
			if book == user2 {
				fmt.Println("Book issued")
				bo = true
			}
		}

		//if book is digital
		if check == "Digitalbook" || check == "digitalbook" {
			//checking for the number of copies required by the user
			var numberofcopies int
			fmt.Println("number of copies")
			fmt.Scanln(&numberofcopies)
			//reading the file
			content, err := ioutil.ReadFile("Digitalbook.json")
			//checking errors
			if err != nil {
				log.Fatal(err)
			}
			//storing file data into another variable of similar type
			user2 := map[int]PhysicalBook{}

			//unmarshalling the data and printing the data
			err = json.Unmarshal(content, &user2)
			if err != nil {
				log.Fatal(err)

			}
			//comparing the two struct as a value in two maps to check if the required book is available or not
			i := 0
			//checking the capacity and availability
			for i < numberofcopies {
				for key, value := range user2 {
					if key < numberofcopies {
						fmt.Println("not Available")
						return false
					}
					if value == book {
						fmt.Println("Book issued")
						bo = true
					}
					i++
				}
			}
		}
		//Storing the member:Book issued
		if bo == true {
			issue := make(map[string]PhysicalBook)
			issue[name] = book
			//Marshaling the data
			bytes, err := json.Marshal(issue)
			//error check
			if err != nil {
				log.Fatalln(err)
			}
			//writing data into json file
			if err = ioutil.WriteFile("issuedList.json", bytes, 0644); err != nil {
				log.Fatalln(err)
			}
		}
	}
	return false

}

// Borrower to show the respective borrower and book issued
func Borrower() {
	//Flag to display Welcome message
	Welcome := flag.String("issued:", "The Following is the member:Book issued", "specify your book name")
	flag.Parse()
	fmt.Println(*Welcome)
	//reading the file
	content, err := ioutil.ReadFile("issuedList.json")
	//checking errors
	if err != nil {
		log.Fatal(err)
	}
	//storing file data into another variable of similar type
	user2 := map[string]PhysicalBook{}

	//unmarshalling the data and printing the data
	err = json.Unmarshal(content, &user2)
	if err != nil {
		log.Fatal(err)

	}
	//printing the list
	fmt.Println(user2)
}

// Booktype enum to define a book-type
type Booktype int

//assigning each Booktype a constant value
const (
	Hardback     Booktype = 0
	Paperback    Booktype = 1
	Encyclopedia Booktype = 2
	Magazine     Booktype = 3
	Comic        Booktype = 4
	Manga        Booktype = 5
	SelfHelp     Booktype = 6
)

//function to check the value and return the designated Booktype
func (booktype Booktype) String() string {
	switch booktype {
	case Hardback:
		return "Hardback"
	case Paperback:
		return "Paperback"
	case Encyclopedia:
		return "Encyclopedia"
	case Magazine:
		return "Magazine"
	case Comic:
		return "Comic"
	case Manga:
		return "Manga"
	case SelfHelp:
		return "SelfHelp"
	default:
		return "Unknown"

	}
}

// DigitalBook struct
type DigitalBook struct {
	Name   string
	Author string
	Kind   Booktype
}

// NewDigitalBook constructor
func NewDigitalBook(New DigitalBook) *DigitalBook {
	//entering Number of Copies
	var copies int
	fmt.Println("Enter Number of Copies:")
	fmt.Scanln(&copies)
	//map for digitalbooks to maintain record of number of copies with a give book
	dig := make(map[int]DigitalBook)
	dig[copies] = New
	//marshaling the data
	bytes, err := json.Marshal(dig)
	if err != nil {
		log.Fatalln(err)
	}
	//writing data into json file
	if err = ioutil.WriteFile("Digitalbook.json", bytes, 0644); err != nil {
		log.Fatalln(err)
	}
	return nil
}

// PhysicalBook struct
type PhysicalBook struct {
	Name   string
	Author string
	Kind   Booktype
}

// NewPhysicalBook constructor
func NewPhysicalBook(New DigitalBook) *PhysicalBook {
	//Marshaling the data
	bytes, err := json.Marshal(New)
	if err != nil {
		log.Fatalln(err)
	}
	//writing data into json file
	if err = ioutil.WriteFile("Physicalbook.json", bytes, 0644); err != nil {
		log.Fatalln(err)
	}
	return nil
}
//calling all the functions as per required order
func main() {
	var s string
	var lib Library
	//lib.addNewMember()
	lib.addNewBook()
	fmt.Println("Details of PhysicalBook or DigitalBook")
	fmt.Scanln(&s)
	var b Book
	if s == "physicalbook" || s == "Physicalbook" {
		b = PhysicalBook{}
		b.Bookdetails()
	} else {
		b = DigitalBook{}
		b.Bookdetails()
	}
	Borrow()
	Borrower()
	lib.collected()
}
