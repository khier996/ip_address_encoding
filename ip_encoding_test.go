package main

import "testing"

func TestConvertIp(t *testing.T) {
  encodedIp, encodeErr := EncodeIp("172.168.5.1")
  if encodedIp != 2896692481 {
    t.Errorf("Converted ip was incorrect; got: %d; want: %d", encodedIp, 2896692481)
  } else if encodeErr != nil {
    t.Errorf("Should not report an error for string '172.168.5.1'")
  }

  encodedIp, encodeErr = EncodeIp("172 . 168 .5.1")
  if encodedIp != 2896692481 {
    t.Errorf("Converted ip was incorrect; got: %d; want: %d", encodedIp, 2896692481)
  } else if encodeErr != nil {
    t.Errorf("Should not report an error for string '172 . 168 .5.1'")
  }

  encodedIp, encodeErr = EncodeIp("17 2.168.5.1")
  if encodeErr == nil {
    t.Errorf("Should report an error for string '17 2.168.5.1'")
  }

  encodedIp, encodeErr = EncodeIp("172.256.5.1")
  if encodeErr == nil {
    t.Errorf("Should report an error for string '172.256.5.1'")
  }
}

