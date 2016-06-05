package main

type relay interface {
	process()
	pass(message message)
}