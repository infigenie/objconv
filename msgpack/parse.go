package msgpack

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/segmentio/objconv"
)

type Parser struct {
	r io.Reader // reader to load bytes from
	i int       // offset of the first unread byte in b
	j int       // offset + 1 of the last unread byte in b
	b [240]byte // read buffer
}

func NewParser(r io.Reader) *Parser {
	return &Parser{r: r}
}

func (p *Parser) Reset(r io.Reader) {
	p.r = r
	p.i = 0
	p.j = 0
}

func (p *Parser) Buffered() io.Reader {
	return bytes.NewReader(p.b[p.i:p.j])
}

func (p *Parser) ParseType() (t objconv.Type, err error) {
	var b []byte

	if b, err = p.peek(1); err != nil {
		return
	}

	switch b[0] {
	case Nil:
		t = objconv.Nil
		return
	}

	err = fmt.Errorf("objconv/msgpack: unknown tag '%#X'", b[0])
	return
}

func (p *Parser) ParseNil() (err error) {
	p.i++
	return
}

func (p *Parser) ParseBool() (v bool, err error) {
	return
}

func (p *Parser) ParseInt() (v int64, err error) {
	return
}

func (p *Parser) ParseUint() (v uint64, err error) {
	return
}

func (p *Parser) ParseFloat() (v float64, err error) {
	return
}

func (p *Parser) ParseString() (v []byte, err error) {
	return
}

func (p *Parser) ParseBytes() (v []byte, err error) {
	return
}

func (p *Parser) ParseTime() (v time.Time, err error) {
	return
}

func (p *Parser) ParseDuration() (v time.Duration, err error) {
	return
}

func (p *Parser) ParseError() (v error, err error) {
	return
}

func (p *Parser) ParseArrayBegin() (n int, err error) {
	return
}

func (p *Parser) ParseArrayEnd(n int) (err error) {
	return
}

func (p *Parser) ParseArrayNext(n int) (err error) {
	return
}

func (p *Parser) ParseMapBegin() (n int, err error) {
	return
}

func (p *Parser) ParseMapEnd(n int) (err error) {
	return
}

func (p *Parser) ParseMapValue(n int) (err error) {
	return
}

func (p *Parser) ParseMapNext(n int) (err error) {
	return
}

func (p *Parser) peek(n int) (b []byte, err error) {
	for (p.i + n) > p.j {
		if err = p.fill(); err != nil {
			return
		}
	}
	b = p.b[p.i : p.i+n]
	return
}

func (p *Parser) fill() (err error) {
	n := p.j - p.i
	copy(p.b[:], p.b[p.i:p.j])
	p.i = 0
	p.j = n

	if n, err = p.r.Read(p.b[p.j:]); n > 0 {
		err = nil
		p.j += n
	} else if err != nil {
		return
	} else {
		err = io.ErrNoProgress
		return
	}

	return
}