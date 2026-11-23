package libclog

type Progress struct {
	printer *Printer
}

func NewProgress() *Progress {
	return &Progress{
		printer: &Printer{},
	}
}
