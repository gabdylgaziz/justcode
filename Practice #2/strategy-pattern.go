package main

import (
	"fmt"
)

type Strategy interface{
	Process() error
	DoSomething() error
}

type FirstAlgorithm struct {}


func (f *FirstAlgorithm) Process() error {
	fmt.Println("1st algo is in process")
	//do something

	return nil
}

type SecondAlgorithm struct {}

func (f *SecondAlgorithm) Process() error {
	fmt.Println("2nd algo is in process")
	//do something

	return nil
}

type Context struct{
	activeStrategy Strategy
}

func(c *Context) SetStrategy(strategy Strategy){
	c.activeStrategy = strategy
}