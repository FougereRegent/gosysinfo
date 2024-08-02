package gosysinfo

type hardware interface {
	ParseContent(content string) interface{}
}
