// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	bar "example.com/bar"
)

// Injectors from wire.go:

func injectFooBar() FooBar {
	foo := provideFoo()
	barBar := bar.ProvideBar()
	fooBar := provideFooBar(foo, barBar)
	return fooBar
}
