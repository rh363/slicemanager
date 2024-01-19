package slicemanager

/*
require string slice,index where insert string and the string to insert
*/
func insert_string(slice []string, index int, s string) []string {
	var ns []string
	for i, line := range slice {
		ns = append(ns, line)
		if i == index {
			ns = append(ns, s)
		}
	}
	return ns
}
