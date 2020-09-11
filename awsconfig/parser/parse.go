package parser

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

type Section struct {
	Header string
	Values map[string]string
}

func Parse(r io.Reader) ([]*Section, error) {
	s := bufio.NewScanner(r)
	var p parser
	for {
		if !s.Scan() {
			if err := s.Err(); err != nil {
				return nil, fmt.Errorf("could not read: %w", err)
			}
			return p.sections, nil
		}
		if err := p.parse(s.Text()); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}
	}
}

type parser struct {
	sections []*Section
	current  *Section
}

var (
	section  = regexp.MustCompile(`^\[\s*(.+?)\s*]$`)
	keyValue = regexp.MustCompile(`^\s*(.+?)\s*=\s*(.*?)\s*$`)
	skipLine = regexp.MustCompile(`^\s*$|^\s*#`)
)

func (p *parser) parse(l string) error {
	if m := section.FindStringSubmatch(l); m != nil {
		p.beginSection(m[1])
		return nil
	}
	if m := keyValue.FindStringSubmatch(l); m != nil {
		if p.current == nil {
			return fmt.Errorf("key-value pair before begin of section")
		}
		p.current.Values[m[1]] = m[2]
		return nil
	}
	if skipLine.MatchString(l) {
		return nil
	}
	return fmt.Errorf("invalid line `%s`", l)
}

func (p *parser) beginSection(header string) {
	profile := &Section{
		Header: header,
		Values: make(map[string]string),
	}
	p.sections = append(p.sections, profile)
	p.current = profile
}
