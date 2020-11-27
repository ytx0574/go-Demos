package model

type Student struct {
	Name string
	Score float64
}

type student struct {
	name string
	score float64
	stu student2 //go中的私有是按包区分的. 同一个包下面的私有可以访问
}

//使用工厂方法构建私有struct
func StudentInstance(name string, score float64) *student {
	var stu student = student{
		name: name,
		score: score,
		stu:student2{
			name: "student2 instance",
		},
	}
	return &stu
}
//外部访问私有变量, 创建一个get方法即可  类似oc的get set. 只是这里要自己手动创建
func (s *student) GetName() string {
	return s.name
}