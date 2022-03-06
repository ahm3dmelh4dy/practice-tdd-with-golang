package append

// AppendByte appends given data to a given slice.
func AppendByte(slice []byte, data ...byte) []byte {
	n := len(slice)
	dataLength := len(data)
	m := n + dataLength

	slice = reallocate(m, slice)
	slice = slice[0:m]
	copy(slice[n:m], data)
	return slice
}

// reallocate the slice if  m > cap(slice), add 1 if the cap(slice) is zero.
func reallocate(m int, s []byte) []byte {
	if m > cap(s) {
		// extending the slice by creating a new slice, copy the data to it, then add a reference to its underlying array to the old slice.
		newSlice := make([]byte, (m+1)*2)
		copy(newSlice, s)
		s = newSlice
	}
	return s
}
