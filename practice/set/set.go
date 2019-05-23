package set

type Set struct {
	Items []int
}

func NewSet(values []int) *Set {
	numOfValues := len(values)
	s := Set{
		Items: make([]int, numOfValues),
	}
	for i := 0; i < len(values); i++ {
		s.Items[i] = values[i]
	}
	return &s
}

func (s *Set) Size() int {
	return len(s.Items)
}

func (s *Set) Includes(value int) bool {
	includesValue := false
	for i := 0; i < len(s.Items); i++ {
		if s.Items[i] == value {
			includesValue = true
		}
	}
	return includesValue
}

// Adds a value by creating a new array of (original length + 1) length
// and adding in new value
func (s *Set) Add(value int) bool {
	valueAlreadyInSet := false
	for i := 0; i < len(s.Items); i++ {
		if s.Items[i] == value {
			valueAlreadyInSet = true
		}
	}

	if !valueAlreadyInSet {
		s.Items = expandItems(s.Items, value)
	}
	return true
}

// Deletes a value by replacing the value to delete with the last in the array,
// then creating a new array of (original length - 1) length and populating it
// with all the values in the current array except the last one
func (s *Set) Delete(value int) bool {
	for i := 0; i < len(s.Items); i++ {
		if s.Items[i] == value {
			s.Items[i] = s.Items[len(s.Items)-1]
			s.Items = shrinkItems(s.Items)
		}
	}
	return true
}

func expandItems(currentItems []int, itemToAdd int) []int {
	newItemsLength := len(currentItems) + 1
	updatedItems := make([]int, newItemsLength)
	for i := 0; i < len(currentItems); i++ {
		updatedItems[i] = currentItems[i]
	}
	updatedItems[newItemsLength-1] = itemToAdd
	return updatedItems
}

func shrinkItems(currentItems []int) []int {
	newItemsLength := len(currentItems) - 1
	updatedItems := make([]int, newItemsLength)
	for i := 0; i < newItemsLength; i++ {
		updatedItems[i] = currentItems[i]
	}
	return updatedItems
}
