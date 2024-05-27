package main

import (
  "fmt"
  "example.com/bank/fileops"
)

const accountBalanceLocalFile = "balance.txt"


func main() {
  var accountBalance, err = fileops.ReadFloatFromLocalFile(accountBalanceLocalFile, 1000)
  if err != nil {
    fmt.Println("ERROR")
    fmt.Println(err)
    fmt.Println("----------")
    // panic(err)
  }

  fmt.Println("Welcome to Go Bank!")

  for {
    presentOptions()

    var choice int
    fmt.Print("Your Choice: ")
    fmt.Scan(&choice)

    switch choice {
    case 1:
      fmt.Println("Your balence is", accountBalance)
    case 2:
      fmt.Print("Deposit Amount: ")
      var depositAmount float64
      fmt.Scan(&depositAmount)
      if depositAmount <= 0 {
        fmt.Println("Invaild amount... Please enter a value higher than 0.")
        return
      }
      accountBalance += depositAmount
      fmt.Println("Updated Balence:", accountBalance)
      fileops.WriteFloatToLocalFile(accountBalance, accountBalanceLocalFile)
    case 3:
      fmt.Print("Withdraw Amount: ")
      var withdrawAmount float64
      fmt.Scan(&withdrawAmount)
      if withdrawAmount <= accountBalance && withdrawAmount > 0 {
        accountBalance -= withdrawAmount
        fileops.WriteFloatToLocalFile(accountBalance, accountBalanceLocalFile)
        continue
      } else {
        fmt.Print("Withdraw amount is higher than current balence. Please make sure you didn't make a typo.")
        continue
      }
    default:
      fmt.Println("Goodbye")
      return
    }
  }  
}


