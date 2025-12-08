package utils

import (
	"errors"
	"log"
	"os"

	"github.com/sqids/sqids-go"
)

var Sqids *sqids.Sqids

func InitSqids() {
	alphabet := os.Getenv("SQIDS_ALPHABET")
	if alphabet == "" {
		alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	}

	s, err := sqids.New(sqids.Options{
		Alphabet:  alphabet,
		MinLength: 8,
	})
	if err != nil {
		log.Fatal("failed to init sqids:", err)
	}

	Sqids = s
}

func EncodeID(id uint) string {
	encoded, err := Sqids.Encode([]uint64{uint64(id)})
	if err != nil {
		return ""
	}
	return encoded
}

func DecodeID(encoded string) (uint, error) {
	ids := Sqids.Decode(encoded)
	if len(ids) == 0 {
		return 0, errors.New("invalid id")
	}
	return uint(ids[0]), nil
}
