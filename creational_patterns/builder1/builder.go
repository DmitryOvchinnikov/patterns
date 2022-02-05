package main

import "fmt"

// BlueprintBuilder has fields exported.
type BlueprintBuilder struct {
	PropertyOne   string
	PropertyTwo   string
	PropertyThree string
}

func newBlueprintBuilder() *BlueprintBuilder {
	return &BlueprintBuilder{}
}

func (b *BlueprintBuilder) SetPropertyOne(value string) {
	b.PropertyOne = value
}

func (b *BlueprintBuilder) SetPropertyTwo(value string) {
	b.PropertyTwo = value
}

func (b *BlueprintBuilder) SetPropertyThree(value string) {
	b.PropertyThree = value
}

// Build returns a fully finished Blueprint object.
func (b *BlueprintBuilder) Build() (*Blueprint, error) {
	if b.PropertyOne != "" && b.PropertyTwo == "" {
		return nil, fmt.Errorf(" Build(): no PropertyOne, no PropertyTwo")
	}

	return &Blueprint{
		b.PropertyOne,
		b.PropertyTwo,
		b.PropertyThree,
	}, nil
}
