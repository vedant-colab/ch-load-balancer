package hashing

func Fnv1aHashFunc(input string) uint64 {
	var hash uint64 = 14695981039346656037
	const prime uint64 = 1099511628211

	for i := range input {
		hash ^= uint64(input[i])
		hash *= prime
	}
	return hash
}
