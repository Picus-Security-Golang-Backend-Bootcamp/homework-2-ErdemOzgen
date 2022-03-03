package model

func DeleteSlice(id int, b []Book) []Book { // check for int values and return error if needed
	b[IDtoIndex(id)].Delete()
	b[IDtoIndex(id)] = b[len(b)-1]
	b = b[:len(b)-1]
	return b
}
