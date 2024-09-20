package main

import (
	"fmt"
	"errors"
)

type List interface {
	Size() int
	Get(index int) (int,error)	
	Add(e int) 
	AddOnIndex(e int, index int) error
	Remove(index int) error
}

type ArrayList struct {
	v []int
	inserted int
}

func (l *ArrayList) Init(size int) error {
	if size > 0 {
		l.v = make([]int, size)
		return nil
	} else {
		return errors.New("Size <= 0")
	}
}

func (l *ArrayList) doubleV() {
	newSize := len(l.v)*2
	newV := make([]int, newSize)
	for i:=0; i < len(l.v); i++ {
		newV[i] = l.v[i]
	} 
	l.v = newV
}

func (l *ArrayList) Size() int{
	return l.inserted
}

func (l *ArrayList) Get(index int) (int,error){
	if index>=0 && index < l.inserted {
		return l.v[index], nil
	} else{
		return index, errors.New("Index fora dos limites da lista") 
	}
}

func (l *ArrayList) Add(e int) {
	l.v[l.inserted] = e
	l.inserted++
}

func main(){
	l := &ArrayList{}
	l.Init(10)
	l.Add(5)
	l.Add(8)
	l.Add(3)
	l.Add(5)
	l.Add(8)
	l.Add(3)
	l.Add(5)
	l.Add(8)
	l.Add(3)
	l.Add(5)
	l.Add(8)
	l.Add(3)
	val,_ := l.Get(0)
	fmt.Println(val)
	
}
