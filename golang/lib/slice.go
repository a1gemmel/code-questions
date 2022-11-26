package lib

func RemoveAtIndex(index int, arr []string) []string {
	newArr := make([]string, 0, len(arr)-1)

	for i, el := range arr {
		if i != index {
			newArr = append(newArr, el)
		}
	}
	return newArr
}
