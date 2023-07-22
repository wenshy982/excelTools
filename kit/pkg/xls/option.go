package xls

import (
	"path/filepath"
)

const (
	Sep              = string(filepath.Separator)                  // 路径分隔符
	defaultInput     = "." + Sep + "input" + Sep + "input.xlsx"    // 默认输入文件
	defaultOutput    = "." + Sep + "output" + Sep + "output.xlsx"  // 默认输出文件
	defaultTemplate  = "." + Sep + "input" + Sep + "template.xlsx" // 默认模板文件
	defaultErrors    = "." + Sep + "input" + Sep + "errors.xlsx"   // 默认错误文件
	defaultOutPutCSV = "." + Sep + "output" + Sep + "output.csv"   // 默认输出 CSV 文件
)

var DefaultOpt = []Opt{
	WithInput(defaultInput),
	WithOutput(defaultOutput),
	WithTemplate(defaultTemplate),
	WithErrors(defaultErrors),
	WithOutputCSV(defaultOutPutCSV),
}

type Opt func(xls *XLS)

// WithInput 设置输入文件
func WithInput(input string) Opt {
	return func(o *XLS) {
		o.Input = input
	}
}

// WithOutput 设置输出文件
func WithOutput(output string) Opt {
	return func(o *XLS) {
		o.Output = output
		o.OutputDir = filepath.Dir(output)
	}
}

// WithTemplate 设置模板文件
func WithTemplate(template string) Opt {
	return func(o *XLS) {
		o.Template = template
	}
}

// WithErrors 设置错误文件
func WithErrors(errors string) Opt {
	return func(o *XLS) {
		o.Errors = errors
	}
}

// WithOutputCSV 设置输出 CSV 文件
func WithOutputCSV(outputCSV string) Opt {
	return func(o *XLS) {
		o.OutPutCSV = outputCSV
	}
}
