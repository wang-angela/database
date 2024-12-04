package main

import (
	"fmt"

	"github.com/wang-angela/database"
)

func main() {
	// Example usage from fig. 2
	inmemoryDB := database.NewInMemoryDB()

	// should return null, because A doesn’t exist in the DB yet
	fmt.Printf("Get(A): %v\n", inmemoryDB.Get("A"))

	// should throw an error because a transaction is not in progress
	if err := inmemoryDB.Put("A", 5); err != nil {
		fmt.Printf("Put(A, 5): Error - %s\n", err)
	}

	// starts a new transaction
	if err := inmemoryDB.BeginTransaction(); err != nil {
		fmt.Printf("BeginTransaction: Error - %s\n", err)
	} else {
		fmt.Println("Transaction started.")
	}

	// set’s value of A to 5, but its not committed yet
	if err := inmemoryDB.Put("A", 5); err != nil {
		fmt.Printf("Put(A, 5): Error - %s\n", err)
	}

	// should return null, because updates to A are not committed yet
	val := inmemoryDB.Get("A")
	if val == nil {
		fmt.Println("Get(A): <nil>")
	} else {
		fmt.Printf("Get(A): %d\n", *val)
	}


	// update A’s value to 6 within the transaction
	if err := inmemoryDB.Put("A", 6); err != nil {
		fmt.Printf("Put(A, 6): Error - %s\n", err)
	}

	// commits the open transaction
	if err := inmemoryDB.Commit(); err != nil {
		fmt.Printf("Commit: Error - %s\n", err)
	} else {
		fmt.Println("Transaction committed.")
	}

	// should return 6, that was the last value of A to be committed
	val = inmemoryDB.Get("A")
	if val == nil {
		fmt.Println("Get(A): <nil>")
	} else {
		fmt.Printf("Get(A): %d\n", *val)
	}


	// throws an error, because there is no open transaction
	if err := inmemoryDB.Commit(); err != nil {
		fmt.Printf("Commit: Error - %s\n", err)
	}

	// throws an error because there is no ongoing transaction
	if err := inmemoryDB.Rollback(); err != nil {
		fmt.Printf("Rollback: Error - %s\n", err)
	}

	// should return null because B does not exist in the database
	fmt.Printf("Get(B): %v\n", inmemoryDB.Get("B"))

	// starts a new transaction
	if err := inmemoryDB.BeginTransaction(); err != nil {
		fmt.Printf("BeginTransaction: Error - %s\n", err)
	} else {
		fmt.Println("Transaction started.")
	}

	// Set key B’s value to 10 within the transaction
	if err := inmemoryDB.Put("B", 10); err != nil {
		fmt.Printf("Put(B, 10): Error - %s\n", err)
	}

	// Rollback the transaction - revert any changes made to B
	if err := inmemoryDB.Rollback(); err != nil {
		fmt.Printf("Rollback: Error - %s\n", err)
	} else {
		fmt.Println("Transaction rolled back.")
	}

	// Should return null because changes to B were rolled back
	fmt.Printf("Get(B): %v\n", inmemoryDB.Get("B"))
}