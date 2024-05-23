package escpos

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/google/gousb"
)

type Printer struct {
	buf   bytes.Buffer
	dev   *gousb.Device
	ctx   *gousb.Context
	WIDTH int
}

func (p *Printer) Print(s string) *Printer {
	p.buf.WriteString(s)
	return p
}

func (p *Printer) Println(s string) *Printer {
	return p.Print(s + EOL)
}

func (p *Printer) Sprintln(format string, params ...any) *Printer {
	s := fmt.Sprintf(format, params...)
	return p.Println(s)
}

func (p *Printer) Rule(c string) *Printer {
	return p.Print(strings.Repeat(c, p.WIDTH))
}

func (p *Printer) Cut() *Printer {
	return p.Print(PAPER_FULL_CUT)
}

func (p *Printer) Feed(lines int) *Printer {
	return p.Print(strings.Repeat(EOL, lines))
}

func (p *Printer) Flush() error {
	intf, done, err := p.dev.DefaultInterface()
	if err != nil {
		return err
	}

	defer done()

	ep, err := intf.OutEndpoint(1)
	if err != nil {
		return err
	}

	for {
		output := p.buf.Next(64)

		if len(output) == 0 {
			output = []byte(CTL_LF)
		}

		_, err := ep.Write(output)

		if err != nil {
			return err
		}
	}
}

func NewUSB() (*Printer, error) {
	p := Printer{
		buf:   *bytes.NewBufferString(HW_INIT),
		ctx:   gousb.NewContext(),
		WIDTH: 42,
	}

	p.ctx.Debug(10)
	dev, err := p.ctx.OpenDeviceWithVIDPID(0x04b8, 0x0e02)

	if err != nil {
		return nil, err
	}

	p.dev = dev
	p.dev.SetAutoDetach(true)

	return &p, nil
}

func (p *Printer) Close() {
	p.dev.Close()
}
