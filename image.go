package margo

type Image interface {
	Authorer
	Commenter
	//	Metadata
	//	ImageDimensioner
}

type ImageDimensioner interface {
	Height() int64
	Width() int64
}
