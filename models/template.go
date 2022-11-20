package models

type TemplateBlocks struct {
	EmailRegister map[string]Template `binding:"required"`
}

type Template struct {
	Exists bool `binding:"required"`
	Data   map[string]string
}

type TemplateVars struct {
	Variable string `binding:"required"`
	Value    string `binding:"required"`
}
