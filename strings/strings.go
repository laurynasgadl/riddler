package strings

import "math/rand"

type Strings struct {
	r *rand.Rand
}

func NewStrings(r *rand.Rand) *Strings {
	return &Strings{
		r: r,
	}
}

// Charset type for random string generation char sets
type Charset string

// various charsets for random string generation
const (
	LowercaseAlphaCharset Charset = "abcdefghijklmnopqrstuvwxyz"
	UppercaseAlphaCharset Charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NumCharset            Charset = "0123456789"
	SymbolsCharset        Charset = "!@#$%^&*+?"

	AlphaNumCharset     = LowercaseAlphaCharset + UppercaseAlphaCharset + NumCharset
	AlphaNumDashCharset = LowercaseAlphaCharset + UppercaseAlphaCharset + NumCharset + "-_"
)

var (
	DefaultLength  uint    = 20
	DefaultCharset Charset = AlphaNumCharset

	defaultOptions = &options{
		Length:  DefaultLength,
		Charset: DefaultCharset,
	}
)

type options struct {
	Length  uint
	Charset Charset
}

type Option func(o *options)

func WithLength(l uint) Option {
	return func(o *options) {
		o.Length = l
	}
}

func WithCharset(cs Charset) Option {
	return func(o *options) {
		o.Charset = cs
	}
}

func applyOptions(opts []Option) *options {
	optCpy := &options{}
	*optCpy = *defaultOptions
	for _, o := range opts {
		o(optCpy)
	}
	return optCpy
}

// Generate returns a randomly generated string using the provided options
func (s *Strings) Generate(opts ...Option) string {
	o := applyOptions(opts)

	b := make([]byte, o.Length)
	for i := range b {
		b[i] = o.Charset[s.r.Intn(len(o.Charset))]
	}
	return string(b)
}
