package model

import "fmt"

type Student struct {
	Name string `json:"name"`
	Age int `aaa:""`
	Score float64
	privatefiled int
	privatefiled2 int
}

func (self Student)String() string {
	return fmt.Sprintf("Name = %v, Age = %d, Score = %f", self.Name, self.Age, self.Score)
}
func (self Student)string_private() {

}

func (self Student)GetName() string {
	return self.Name
}

func (self *Student)AddScore(score float64) (newScore float64, flag bool) {
	self.Score = score + self.Score
	fmt.Printf("addScore:%f, oldScore:%f\n", score, self.Score)
	flag = true
	return self.Score, flag
}
