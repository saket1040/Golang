package main

import (
    "fmt"
    "splitwise/service"
    "splitwise/storage"
    "splitwise/strategy"
)

func main() {
    store := storage.NewMemoryStorage()

    userService := service.NewUserService(store)
    groupService := service.NewGroupService(store)
    expenseService := service.NewExpenseService(store, &strategy.EqualSplit{})

    userService.CreateUser("u1", "Alice", "alice@example.com")
    userService.CreateUser("u2", "Bob", "bob@example.com")
    userService.CreateUser("u3", "Charlie", "charlie@example.com")

    groupService.CreateGroup("g1", "Trip", []string{"u1", "u2", "u3"})

    err := expenseService.AddExpense("u1", "g1", 150.0, "Hotel booking")
    if err != nil {
        fmt.Println("Error adding expense:", err)
    } else {
        fmt.Println("Expense added successfully")
    }
}