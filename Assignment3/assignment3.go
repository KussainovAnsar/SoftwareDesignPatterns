package main

import (
	"fmt"
	"sync"
)

type FileSystemManager struct {
	receiver FileSystemReceiver
}

var instance *FileSystemManager
var once sync.Once

func GetFileSystemManager() *FileSystemManager {
	once.Do(func() {
		instance = &FileSystemManager{
			receiver: &ConcreteFileSystemReceiver{},
		}
	})
	return instance
}

func (fsm *FileSystemManager) ExecuteCommand(command Command) {
	command.Execute(fsm.receiver)
}

type FileSystemReceiver interface {
	CreateFile(fileName string)
	DeleteFile(fileName string)
	CopyFile(source, destination string)
	MoveFile(source, destination string)
}

type ConcreteFileSystemReceiver struct{}

func (fs *ConcreteFileSystemReceiver) CreateFile(fileName string) {
	fmt.Printf("Creating file: %s\n", fileName)
}

func (fs *ConcreteFileSystemReceiver) DeleteFile(fileName string) {
	fmt.Printf("Deleting file: %s\n", fileName)
}

func (fs *ConcreteFileSystemReceiver) CopyFile(source, destination string) {
	fmt.Printf("Copying file from %s to %s\n", source, destination)
}

func (fs *ConcreteFileSystemReceiver) MoveFile(source, destination string) {
	fmt.Printf("Moving file from %s to %s\n", source, destination)
}

type Command interface {
	Execute(receiver FileSystemReceiver)
}

type ConcreteCommand struct {
	FileName    string
	Source      string
	Destination string
}

func (cc *ConcreteCommand) Execute(receiver FileSystemReceiver) {
	switch cc.FileName {
	case "create":
		receiver.CreateFile(cc.Source)
	case "delete":
		receiver.DeleteFile(cc.Source)
	case "copy":
		receiver.CopyFile(cc.Source, cc.Destination)
	case "move":
		receiver.MoveFile(cc.Source, cc.Destination)
	default:
		fmt.Println("Invalid command")
	}
}

func main() {
	fileSystemManager := GetFileSystemManager()
	createCommand := &ConcreteCommand{FileName: "create", Source: "file.txt"}
	deleteCommand := &ConcreteCommand{FileName: "delete", Source: "file.txt"}
	copyCommand := &ConcreteCommand{FileName: "copy", Source: "source.txt", Destination: "destination.txt"}
	moveCommand := &ConcreteCommand{FileName: "move", Source: "source.txt", Destination: "destination.txt"}
	fileSystemManager.ExecuteCommand(createCommand)
	fileSystemManager.ExecuteCommand(deleteCommand)
	fileSystemManager.ExecuteCommand(copyCommand)
	fileSystemManager.ExecuteCommand(moveCommand)
}
