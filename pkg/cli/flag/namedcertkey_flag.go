package flag

import (
	"errors"
	"flag"
	"strings"
)

// NamedCertKey is a flag value parsing "certfile,keyfile" and "certfile,keyfile:name,name,name".
type NamedCertKey struct {
	Names             []string
	CertFile, KeyFile string
}

var _ flag.Value = &NamedCertKey{}

func (nck *NamedCertKey) String() string {
	ckf := nck.CertFile + "," + nck.KeyFile
	if len(nck.Names) > 0 {
		ckf = ckf + ":" + strings.Join(nck.Names, ",")
	}

	return ckf
}

func (nck *NamedCertKey) Set(value string) error {
	cs := strings.SplitN(value, ":", 2)
	var kct string
	if len(cs) == 2 {
		var names string
		kct, names = strings.TrimSpace(cs[0]), strings.TrimSpace(cs[1])
		if names == "" {
			return errors.New("empty names list is not allowed")
		}
		nck.Names = nil
		for _, name := range strings.Split(names, ",") {
			nck.Names = append(nck.Names, strings.TrimSpace(name))
		}
	} else {
		nck.Names = nil
		kct = strings.TrimSpace(cs[0])
	}
	cs = strings.Split(kct, ",")
	if len(cs) != 2 {
		return errors.New("expected comma separated certificate and key file paths")
	}
	nck.CertFile = strings.TrimSpace(cs[0])
	nck.KeyFile = strings.TrimSpace(cs[1])

	return nil
}

func (*NamedCertKey) Type() string {
	return "namedCertKey"
}

// NamedCertKeyArray is a flag value parsing NamedCertKeys, each passed with its own
// flag instance (in contrast to comma separated slices).
type NamedCertKeyArray struct {
	value   *[]NamedCertKey
	changed bool
}

// NewNamedCertKeyArray creates a new NamedCertKeyArray with the internal value
// pointing to p.
func NewNamedCertKeyArray(p *[]NamedCertKey) *NamedCertKeyArray {
	return &NamedCertKeyArray{
		value: p,
	}
}

var _ flag.Value = &NamedCertKeyArray{}

func (a *NamedCertKeyArray) Set(value string) error {
	nck := NamedCertKey{}
	err := nck.Set(value)
	if err != nil {
		return err
	}
	if !a.changed {
		*a.value = []NamedCertKey{nck}
		a.changed = true
	} else {
		*a.value = append(*a.value, nck)
	}

	return nil
}

func (a *NamedCertKeyArray) String() string {
	ncks := make([]string, 0, len(*a.value))
	for i := range *a.value {
		ncks = append(ncks, (*a.value)[i].String())
	}

	return "[" + strings.Join(ncks, ";") + "]"
}

func (*NamedCertKeyArray) Type() string {
	return "namedCertKey"
}
