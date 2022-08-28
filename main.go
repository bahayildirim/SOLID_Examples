package main

import "fmt"

func main() {
}

//Single Responsibility Principle

//Incorrect
func func1() {
	y := 10
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}

	if y != 10 {
		fmt.Printf("y is not 10")
	} else {
		fmt.Printf("y is 10")
	}
}

//Correct
func func2() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
}

func func3() {
	y := 10
	if y != 10 {
		fmt.Printf("y is not 10")
	} else {
		fmt.Printf("y is 10")
	}
}

//Open/Closed Principle

//Incorrect
func old1(num int) {
	if num < 10 {
		fmt.Println("num is smaller than 10")
	}
}

func new1(number float32) {
	if num < 10 {
		fmt.Println("num is smaller than 10")
	}
}

//Correct
func old2(num int) {
	if num < 10 {
		fmt.Println("num is smaller than 10")
	}
}

func new2(num []int) {
	for i := 0; i < len(num); i++ {
		if num[i] < 10 {
			fmt.Println("num is smaller than 10")
		}
	}
}

//Liskov Substitution Principle relies on classes,so it cannot be shown on Go

//Interface Segregation Principle

//Incorrect
func CustomerInterface1() {
	/*
		checkProducts()
		buyProduct()
		login()
		addProduct()
		removeProduct()
		updateProduct()
	*/
}

//Correct
func CustomerInterface2() {
	/*
		checkProducts()
		buyProduct()
		login()
	*/
}

//Dependency Inversion Principle

//Incorrect
func BudgetReport1() {
	/*
		open(database)
		save(database)
	*/
}

func database1() {
	/*
		get()
		insert()
		update()
		delete()
	*/
}
func Module1() {
	/*
		database = database1()
		report = BudgetReport(database)
	*/
}

//Correct
func BudgetReport2() {
	/*
		open(database)
		save(database)
	*/
}

func get() {}

func insert() {}

func update() {}

func delete() {}

func sql() {
	/*
		get()
		insert()
		update()
		delete()
	*/
}

func mongo() {
	/*
		get()
		insert()
		update()
		delete()
	*/
}

func Module2() {
	/*
		database = database1()
		report = BudgetReport(database)
	*/
}
