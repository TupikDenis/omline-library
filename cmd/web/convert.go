package main

import "strconv"

func convertUInt(id string) uint {
	idu64, _ := strconv.ParseUint(id, 10, 32)
	idu := uint(idu64)
	return idu
}
