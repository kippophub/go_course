package escape_analysis

type LargeObject struct {
	ID   int
	Name string
	Data []int
}

func ProcessByValue(obj LargeObject) int {
	sum := 0
	for _, r := range obj.Data {
		sum += r
	}
	return sum
}

func ProcessByPointer(obj *LargeObject) int {
	sum := 0
	obj.ID += 1
	for _, r := range obj.Data {
		sum += r
	}
	return sum
}

func CreateObjectOnStack() LargeObject {
	data := make([]int, 10)
	obj := LargeObject{ID: 1, Name: "StackObject", Data: data}
	return obj
}

func CreateObjectOnHeap() *LargeObject {
	data := make([]int, 10000)
	obj := LargeObject{ID: 2, Name: "HeapObject", Data: data}
	return &obj
}
