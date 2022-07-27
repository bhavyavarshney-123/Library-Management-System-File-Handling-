# Library-Management-System
• Only a registered member of the library can borrow books from the library.

• A book may be physical or digital in nature which lends itself to different borrowing constraints which is that physical books can only be borrowed by one member at a time while digital books contain a certain number of copies each of which can be borrowed by a member.

• Only the member who has currently borrowed a book can return it to the library.

In an effort towards developing and modelling this library management system,  the Go Programming Language and the following language elements

1.  enums to define a type BookType with variants such as  Hardback, Paperback, Encyclopedia, Magazine, Comic, etc. Each of these types can be associated with either a physical or digital book (or both)

2. Book must be an Interface type with methods Bookdetails whether it is a digital or physical book. Methods to set a borrower to it (returns a Boolean). When Borrow() is called on the Book to set the borrower
  a. PhysicalBook objects will only allow one to borrow and return false if already borrowed
  b. DigitalBook objects will allow multiple borrowers until their capacity is full and return false if capacity is full.
	
3. PhysicalBook and DigitalBook is a structs that satisfy the Book interface and implement their own constructors NewPhysicalBook and NewDigitalBook

4. Library is a type that has methods to add new books to the inventory and register new members to the userbase. A member must have the ability to borrow a book from a Library and return it.
