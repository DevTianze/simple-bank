package api

import (
	db "github.com/DevTianze/simple-bank/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/golang/mock/mockgen/model"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	// add routes to router
	server.router = router
	return server
}


// q: what's a struct pointer? what does it do?
// a: it's a pointer to a struct, it's used to pass a struct by reference

// q: can I use a interface pointer?
// a: no, interface is a type, not a struct, you can't use a pointer to a type

// q: why do we need to use a pointer to a struct?
// a: because we want to modify the struct, not just read it

// q: what's the difference between a pointer to a struct and a pointer to a type?
// a: a pointer to a struct can be used to modify the struct, but a pointer to a type can't be used to modify the type

// q: what's the difference between a pointer to a struct and a struct?
// a: a pointer to a struct can be used to modify the struct, but a struct can't be used to modify the struct

// q: what's the difference between a pointer to a type and a type?
// a: a pointer to a type can't be used to modify the type, but a type can be used to modify the type

// q: what's use of pointer to a function
// a: it's used to pass a function by reference

// a: in what case would I need to use a pointer to a struct?
// q: when you want to modify the struct

// q: I didn't modify it above in the NewServer function?
// a: you didn't modify it in the NewServer function, but you modified it in the createAccount function

// q: how ?
// a: you modified it by calling the store.CreateAccount function
//   the store.CreateAccount function takes a pointer to a CreateAccountParams struct as an argument
//   the store.CreateAccount function modifies the CreateAccountParams struct


// a: how does that count as a modify?
// q: because the CreateAccountParams struct is modified by the store.CreateAccount function

// q: what has it to do with the pointer of store?
// a: the store is a pointer to a struct, so when you call the store.CreateAccount function, you are modifying the struct

// q: do you mean that store struct does not have a CreateAccount function before? 
// a: yes, the store struct does not have a CreateAccount function before

// q: therefore, as I called it to use a function it doesn't have before, that's a modification?
// a: yes, that's a modification


// q: why don't I just change the store struct without using a pointer?
// a: because you want to modify the store struct, not just read it

// q: can I change the struct?
// a: no, you can't change the struct because it's a struct, not a pointer to a struct

// q: so pointer is like create a new struct?
// a: no, pointer is not like create a new struct, pointer is a pointer to a struct

// q: but How am I able to change the struct by using pointer? does it create a new object in the memory?
// a: no, it doesn't create a new object in the memory, it just points to the struct

// q: what exactly is a pointer?
// a: a pointer is a variable that stores the address of a value

// q: does the pointer variable stores the changes I made to it as well?
// a: no, the pointer variable doesn't store the changes you made to it, it just stores the address of the value

// q: how's my change to the pointer variable reflected in the struct?
// a: your change to the pointer variable is reflected in the struct because the pointer variable points to the struct

// q: why don't I just copy a new struct and change it and use it here?
// a: because you want to modify the struct, not just read it