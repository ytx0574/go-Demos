package model


type Person struct {
	Name string
	Age int
	Height int
}

type SlicePerson []Person
type SortPersonType int
type SortPersonRule float64

var (
	SortPersonTypeValue SortPersonType
 	SortPersonRuleValue SortPersonRule
)


const KSortPersonType_Age SortPersonType = 1
const KSortPersonType_Height SortPersonType = 2
const KSortPersonRule_Ascending = 1
const KSortPersonRule_Descending = 2

func (self SlicePerson) Len() int {
	return len(self)
}
func (self SlicePerson) Less(i, j int) bool {
	if SortPersonTypeValue == KSortPersonType_Height {
		if SortPersonRuleValue == KSortPersonRule_Descending {
			return self[i].Height > self[j].Height //前面大于后面 降序
		}else {
			return self[i].Height < self[j].Height //前面小于后面 升序
		}
	}else {
		if SortPersonRuleValue == KSortPersonRule_Descending {
			return self[i].Age > self[j].Age
		}else {
			return self[i].Age < self[j].Age
		}
	}
}
func (self SlicePerson) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}