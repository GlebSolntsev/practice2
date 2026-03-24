package main

import "fmt"

type IBuilder interface {
	BuildHeader()
	BuildBody(data string)
	BuildFooter()
	GetReport() string
}

type TextReportBuilder struct {
	report string
}

func (b *TextReportBuilder) BuildHeader() {
	b.report += "--- НАЧАЛО ОТЧЕТА ---\n"
}

func (b *TextReportBuilder) BuildBody(data string) {
	b.report += fmt.Sprintf("ДАННЫЕ: %s\n", data)
}

func (b *TextReportBuilder) BuildFooter() {
	b.report += "--- КОНЕЦ ОТЧЕТА ---\n"
}

func (b *TextReportBuilder) GetReport() string {
	return b.report
}

type Director struct {
	builder IBuilder
}

func (d *Director) SetBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) ConstructReport(content string) {
	d.builder.BuildHeader()
	d.builder.BuildBody(content)
	d.builder.BuildFooter()
}

func main() {
	builder := &TextReportBuilder{}
	director := Director{}

	director.SetBuilder(builder)
	director.ConstructReport("Студент прошел 5 курсов из 10")

	fmt.Println("Builder Pattern Demo:")
	fmt.Println(builder.GetReport())
}
