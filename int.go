package chain

// Int takes a bunch of int-values and creates an Int Collection out of it.
func Int(values ...int) IntC {
	result := make(IntC, len(values))

	for index, value := range values {
		result[index] = value
	}

	return result
}
