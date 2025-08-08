package main

type Task interface {
	Execute() error
	Rollback() error
}
