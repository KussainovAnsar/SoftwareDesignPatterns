package main

import "fmt"

type TextFormatter interface {
	Format(text string) string
}
type ConcreteTextFormatter struct{}

func (ctf *ConcreteTextFormatter) Format(text string) string {
	return text
}

type Decorator struct {
	Component TextFormatter
}
type BoldDecorator struct {
	Decorator
}

func (bd *BoldDecorator) Format(text string) string {
	return "<b>" + bd.Component.Format(text) + "</b>"
}

type ItalicsDecorator struct {
	Decorator
}

func (id *ItalicsDecorator) Format(text string) string {
	return "<i>" + id.Component.Format(text) + "</i>"
}

type UnderlineDecorator struct {
	Decorator
}

func (ud *UnderlineDecorator) Format(text string) string {
	return "<u>" + ud.Component.Format(text) + "</u>"
}

type StrikethroughDecorator struct {
	Decorator
}

func (sd *StrikethroughDecorator) Format(text string) string {
	return "<s>" + sd.Component.Format(text) + "</s>"
}

func main() {
	component := &ConcreteTextFormatter{}

	boldText := &BoldDecorator{Decorator{Component: component}}
	italicText := &ItalicsDecorator{Decorator{Component: component}}
	underlineText := &UnderlineDecorator{Decorator{Component: component}}
	italicBoldText := &ItalicsDecorator{Decorator{Component: boldText}}
	underlineItalicBoldText := &UnderlineDecorator{Decorator{Component: italicBoldText}}
	strikethroughText := &StrikethroughDecorator{Decorator{Component: component}}

	text := "Hello, here is the usage of decorator pattern!"
	formattedText := underlineItalicBoldText.Format(text)
	fmt.Println("Formatted text:", formattedText)
	boldFormattedText := boldText.Format(text)
	italicFormattedText := italicText.Format(text)
	underlineFormattedText := underlineText.Format(text)
	strikethroughFormattedText := strikethroughText.Format(text)

	fmt.Println("Bold:", boldFormattedText)
	fmt.Println("Italics:", italicFormattedText)
	fmt.Println("Underline:", underlineFormattedText)
	fmt.Println("Strikethrough:", strikethroughFormattedText)
}
