package main

import (
  "fmt"
  "os"
  "strconv"
  "errors"
)

const accountBalanceLocalFile = "balance.txt"

func writeBalanceToLocalFile(balance float64) {
  balanceText := fmt.Sprint(balance)
  os.WriteFile(accountBalanceLocalFile, []byte(balanceText), 0644)
}

func readBalanceFromLocalFile() (float64, error) {
  data, err := os.ReadFile(accountBalanceLocalFile)
  if err != nil {
    return 1000, errors.New("Failed to find file")
  }
  balanceText := string(data)
  balance, err :=strconv.ParseFloat(balanceText, 64) 
  if err != nil {
   return 1000, errors.New("Failed to parse store for vaild numbers") 
  }
  return balance, nil
}

func main() {
  var accountBalance, err = readBalanceFromLocalFile()
  if err != nil {
    fmt.Println("ERROR")
    fmt.Println(err)
    fmt.Println("----------")
    // panic(err)
  }

  fmt.Println("Welcome to Go Bank!")
  var havingFun bool = true

  for havingFun {
    fmt.Println("What would you like to do?")
    fmt.Println("1. Check Balence")
    fmt.Println("2. Deposit Money")
    fmt.Println("3. Withdraw Money")
    fmt.Println("4. Exit")

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
      writeBalanceToLocalFile(accountBalance)
    case 3:
      fmt.Print("Withdraw Amount: ")
      var withdrawAmount float64
      fmt.Scan(&withdrawAmount)
      if withdrawAmount <= accountBalance && withdrawAmount > 0 {
        accountBalance -= withdrawAmount
        writeBalanceToLocalFile(accountBalance)
        continue
      } else {
        fmt.Print("Withdraw amount is higher than current balence. Please make sure you didn't make a typo.")
        continue
      }
    default:
      fmt.Println("Goodbye")
      //break
      return
    }

    // if choice == 1 {
    //   fmt.Println("Your balence is", accountBalence)
    //   continue
    // } else if choice == 2 {
    //   fmt.Print("Deposit Amount: ")
    //   var depositAmount float64
    //   fmt.Scan(&depositAmount)
    //   if depositAmount <= 0 {
    //     fmt.Println("Invaild amount... Please enter a value higher than 0.")
    //     havingFun = false
    //   }
    //   accountBalence += depositAmount
    //   fmt.Println("Updated Balence:", accountBalence)
    //   continue
    // } else if choice == 3 {
    //   fmt.Print("Withdraw Amount: ")
    //   var withdrawAmount float64
    //   fmt.Scan(&withdrawAmount)
    //   if withdrawAmount <= accountBalence && withdrawAmount > 0 {
    //     accountBalence -= withdrawAmount
    //     continue
    //   } else {
    //     fmt.Print("Withdraw amount is higher than current balence. Please make sure you didn't make a typo.")
    //     havingFun = false
    //   }
    // } else {
    //   fmt.Println("Goodbye")
    //   havingFun = false
    // }
    // 
  }  
}

