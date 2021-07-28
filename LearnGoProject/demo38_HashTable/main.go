package main

import (
	"fmt"
	"reflect"
)

type LinkNode struct {
	Id int
	Age int
	Name string
	Weight float32
	next *LinkNode
}

type Link struct {
	Head *LinkNode
}

func (this *Link)insert(node *LinkNode) {
	if this.Head == nil {
		this.Head = node
		return
	}

	cur := this.Head
	for {
		//找到小于当前node的id的对象 就退出
		if node.Id > cur.Id {
			break
		}else if cur.next == nil {//当然指针下一个为空, 标记着是最后一个对象
			break
		}

		cur = cur.next
	}

	//提前储存next对象, 并插入新的node
	t := cur.next
	cur.next = node
	node.next = t
}
func (this *Link)find(r map[string]interface{}) []*LinkNode {

	var l []*LinkNode
	t := this.Head

	for {
		if t == nil {
			break
		}
		rType := reflect.TypeOf(t).Elem()
		rVal := reflect.ValueOf(t).Elem()

		var findVal *LinkNode
		for key, val := range r {
			_, ok := rType.FieldByName(key)
			if !ok {
				break
			}else {
				fieldVale := rVal.FieldByName(key)
				var oVal interface{}
				switch val.(type) {
				case int:
					oVal = int(fieldVale.Int())
				case float64:
					oVal = fieldVale.Float()
				case string:
					oVal = fieldVale.String()
				}

				if oVal == val {
					findVal = t
				}else {
					//一旦找到的值不相等, 直接跳出循环
					findVal = nil
					break
				}
			}
		}

		if findVal != nil {
			l = append(l, findVal)
		}

		t = t.next
	}

	return l
}

func (this *Link)update(node *LinkNode) bool {
	t := this.Head
	var cur *LinkNode
	var pre *LinkNode
	for {
		if t == nil {
			break
		}else if t.Id == node.Id && t.Name == node.Name {  //通过id+名字来匹配
			cur = t
			break
		}
		pre = t
		t = t.next
	}

	if cur.next != nil {
		node.next = cur.next
	}

	if cur != nil {
		if cur == this.Head {
			this.Head = node
		}else {
			pre.next = node
		}
	}

	return cur != nil
}
func (this *Link)delete(node *LinkNode) bool {
	t := this.Head
	var val *LinkNode
	var pre *LinkNode
	for {
		if t == nil {
			break
		}else if t.Id == node.Id && t.Name == node.Name {
			val = t
			break
		}
		pre = t
		t = t.next
	}

	if val != nil {
		pre.next = val.next
		return true
	}else {
		return false
	}
}

func (this *Link)show() {
	t := this.Head
	for {
		if t == nil {
			break
		}
		fmt.Printf("ID:%v, Name:%v, age:%v, weight:%v -->", t.Id, t.Name, t.Age, t.Weight)
		t = t.next
	}
}

type HashTable struct {
	info []Link
}

func (this *HashTable)insert(node *LinkNode) {
	this.info[this.hash(node.Id)].insert(node)
}

func (this *HashTable)find(filterMap map[string]interface{}) []*LinkNode {
	id, ok := filterMap["Id"]
	if !ok {
		panic("筛选map必须带入node的id")
	}

	return this.info[this.hash(int (id.(int)))].find(filterMap)
}

func (this *HashTable)update(node *LinkNode) bool {
	return this.info[this.hash(node.Id)].update(node)
}

func (this *HashTable)delete(node *LinkNode) bool {
	return this.info[this.hash(node.Id)].delete(node)
}

func (this *HashTable)showInfo()  {
	for i, v := range this.info {
		fmt.Printf("队列:%v >>>", i)
		v.show()
		fmt.Println()
	}
}


func (this *HashTable)hash(id int) int {
	//通过特殊方式把数据散到各自的列表 这里简单的处理id
	return id % len(this.info)
}




func CreateHashTable(linkCount int) HashTable {
	return HashTable{info: make([]Link, linkCount)}
}

func main() {
	hashTable := CreateHashTable(10)

	hashTable.showInfo()

	hashTable.insert(&LinkNode{
		Id: 1,
		Name: "张三",
		Age: 1,
		Weight: 111,
	})
	hashTable.insert(&LinkNode{
		Id: 21,
		Name: "李四",
		Age: 11,
		Weight: 211,
	})
	hashTable.insert(&LinkNode{
		Id: 3,
		Name: "王五",
		Age: 33,
		Weight: 333,
	})
	hashTable.insert(&LinkNode{
		Id: 11,
		Name: "赵六",
		Age: 11,
		Weight: 111,
	})
	hashTable.insert(&LinkNode{
		Id: 16,
		Name: "钱多",
		Age: 21,
		Weight: 123,
	})
	hashTable.insert(&LinkNode{
		Id: 16,
		Name: "钱多多",
		Age: 21,
		Weight: 123,
	})
	hashTable.showInfo()

	//todo:不可直接用结构体转map. 转map的同时, 字段的类型会发生变化. 比如int -> float64. 导致在使用reflect校验时无法识别(值相同, 类型不同)
	//findNode := LinkNode{
	//	"Id": 16,
	//	"Name": "钱多",
	//	"Age": 21,
	//}
	//
	//
	//bytes, error := json.Marshal(findNode)
	//fmt.Println(error)
	//var findMap map[string]interface{}
	//json.Unmarshal(bytes, &findMap)

	findList := hashTable.find(map[string]interface{}{
		"Id": 16,
		"Name": "钱多",
		"Age": 21,
	})
	fmt.Printf("查找的数据%v\n", findList)
	for _, v := range findList {
		fmt.Printf("ID:%v, Name:%v, age:%v, weight:%v\n", v.Id, v.Name, v.Age, v.Weight)
	}

	hashTable.update(&LinkNode{
		Id: 16,
		Name: "钱多多",
		Age: 21,
		Weight: 66,
	})
	hashTable.showInfo()

	hashTable.delete(&LinkNode{
		Id:11,
		Name: "赵六",
	})
	hashTable.showInfo()
}