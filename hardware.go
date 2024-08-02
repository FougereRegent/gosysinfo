package main

type hardware interface {
	ParseContent(content string) interface{}
}
