package main

 import (
         "fmt"
         "strconv"
         "strings"
         "errors"
         "os"
 )

 func EncodeIp(ipString string) (uint32, error) {
  var ipInteger uint32
  parts := strings.Split(ipString, ".")
  if len(parts) != 4 {
    return 0, errors.New("IP string is not of right format")
  }

  shiftSize := uint(24)
  for _, part := range parts {
    num, convErr := partToInt(part)
    if convErr != nil { return 0, convErr }
    num32 := uint32(num)
    num32 = num32 << shiftSize
    shiftSize -= 8
    ipInteger += num32
  }
  return ipInteger, nil
}

func partToInt(part string) (int, error) {
  part = strings.TrimSpace(part)
  num, convErr := strconv.Atoi(part)
  if convErr != nil {
    return 0, errors.New(fmt.Sprintf("couldn't convert %s to number", part))
  } else if num > 255 {
    return 0, errors.New("each part of an ip address cannot be larger than 255")
  }
  return num, nil
}

 func main() {
  ipString := os.Getenv("IP_STRING")
  if ipString == "" { return }

  if ipInteger, err := EncodeIp(ipString); err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("ipInteger: ", ipInteger)
  }
 }
