package main

import (
	"fmt"
	"sync"
)

// Интерфейс продукта
type ReportBuilder interface {
	SetHeader(title string)
	GetResult() string
}

// Конкретные продукты
type PDFBuilder struct{ data string }

func (b *PDFBuilder) SetHeader(t string) { b.data = "[PDF] " + t }
func (b *PDFBuilder) GetResult() string  { return b.data }

type HTMLBuilder struct{ data string }

func (b *HTMLBuilder) SetHeader(t string) { b.data = "<h1>" + t + "</h1>" }
func (b *HTMLBuilder) GetResult() string  { return b.data }

// Singleton Фабрика
type ConcreteReportFactory struct{}

var instance *ConcreteReportFactory
var once sync.Once

func GetFactory() *ConcreteReportFactory {
	once.Do(func() {
		instance = &ConcreteReportFactory{}
	})
	return instance
}

func (f *ConcreteReportFactory) CreateBuilder(format string) ReportBuilder {
	switch format {
	case "pdf":
		return &PDFBuilder{}
	case "html":
		return &HTMLBuilder{}
	default:
		return nil
	}
}

func main_1() {
	// Инициализация модуля: go mod init pattern_factory
	factory := GetFactory()

	pdf := factory.CreateBuilder("pdf")
	pdf.SetHeader("Отчет по обучению")

	html := factory.CreateBuilder("html")
	html.SetHeader("Отчет по обучению")

	fmt.Println("Factory Method Demo:")
	fmt.Println(pdf.GetResult())
	fmt.Println(html.GetResult())
}
