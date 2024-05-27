package fileops

import (
  "errors"
  "fmt"
  "os"
  "strconv"
)


func WriteFloatToLocalFile(value float64, fileName string) {
  valueText := fmt.Sprint(value)
  os.WriteFile(fileName, []byte(valueText), 0644)
}

func ReadFloatFromLocalFile(fileName string, defaultValue float64) (float64, error) {
  data, err := os.ReadFile(fileName)
  if err != nil {
    return defaultValue, errors.New("Failed to find file")
  }
  valueText := string(data)
  value, err :=strconv.ParseFloat(valueText, 64) 
  if err != nil {
   return defaultValue, errors.New("Failed to parse store for vaild numbers") 
  }
  return value, nil
}

