package rpdf

type ResourceManager interface {
}

type stdResourceManeger struct {
}

func newStdRecourceManager() ResourceManager {
	return &stdResourceManeger{}
}
