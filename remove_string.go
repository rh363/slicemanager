package slicemanager

/*
require string slice and index to remove
*/

func remove_string(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...) // it take content in slice until the index passed and append every content after the index to it, ... is used for pass element in second slice one by one
}
