package margo

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/pkg/errors"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/zRedShift/mimemagic"
)

type image struct {
	authors  []string
	comments []string
}

func (i *image) Author() string {
	return strings.Join(i.authors, ", ")
}
func (i *image) Authors() []string {
	return i.authors
}
func (i *image) SoftwareEditors() []string {
	log.Println("SoftwareEditors not implemented")
	return []string{}
}

func (i *image) Comment() string {
	return strings.Join(i.comments, "\n")
}
func (i *image) Comments() []string {
	return i.comments
}

func ImageMetadata(r io.Reader) (Image, error) {
	var buf bytes.Buffer
	tee := io.TeeReader(r, &buf)

	mediaType, err := mimemagic.MatchReader(tee, "")
	if err != nil {
		return nil, err
	}

	if mediaType.MediaType() != "image/jpeg" {
		return nil, errors.New("not an image file")
	}

	x, err := exif.Decode(&buf)
	if err != nil {
		return nil, err
	}

	img := image{}

	authorTag, err := x.Get(exif.Artist)
	if err == nil {
		a, err := authorTag.StringVal()
		if err == nil {
			img.authors = append(img.authors, a)
		}
	}

	commentTag, err := x.Get(exif.UserComment)
	if err == nil {
		// UserComment type is undefined, we get it as JSON thanks to the String() helper
		c := commentTag.String()
		if err == nil {
			img.comments = append(img.comments, c)
			fmt.Println(img.comments)
		}
	}

	return &img, nil

}
