package hardware

type Hardware interface {
	ParseContent(content string) interface{}
}
