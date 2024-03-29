package margo

import "time"

type Metadata interface {
	Authorer
	Locationer
	Timer
	Licenser
	Commenter
}

type Authorer interface {
	Author() string
	Authors() []string
	SoftwareEditors() []string
}

type Locationer interface {
	Location() (x, y int64)
}

type Timer interface {
	CreatedAt() time.Time
	UpdatedAt() time.Time
}

type Licenser interface {
	Licenses() []string
}

type Commenter interface {
	Comment() string
	Comments() []string
}
