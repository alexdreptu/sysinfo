package uptime

import (
	"time"

	"github.com/cactus/gostrftime"

	"golang.org/x/sys/unix"
)

type fetchFunc func(info *unix.Sysinfo_t) error

type Uptime struct {
	Uptime time.Duration
	Sep    byte
	F      fetchFunc
}

func (u Uptime) String() string {
	var buf [32]byte
	w := len(buf)

	up := uint64(u.Uptime)
	neg := u.Uptime < 0
	if neg {
		up = -up
	}
	w--
	buf[w] = 's'

	w, up = formatFrac(buf[:w], up, 9)

	// up is now integer seconds
	w = formatInt(buf[:w], up%60)
	up /= 60

	// up is now integer minutes
	if up > 0 {
		var l byte = 'm'
		w = w - 2
		buf[w] = l
		buf[w+1] = u.Sep
		// buf[bytes.IndexByte(buf[:], l)+1] = u.Sep
		w = formatInt(buf[:w], up%60)
		up /= 60

		// up is now integer hours
		if up > 0 {
			var l byte = 'h'
			w = w - 2
			buf[w] = l
			buf[w+1] = u.Sep
			// buf[bytes.IndexByte(buf[:], l)+1] = u.Sep
			w = formatInt(buf[:w], up%24)
			up /= 24
		}

		// up is now integer days
		if up > 0 {
			var l byte = 'd'
			w = w - 2
			buf[w] = l
			buf[w+1] = u.Sep
			// buf[bytes.IndexByte(buf[:], l)+1] = u.Sep
			w = formatInt(buf[:w], up)
		}
	}

	if neg {
		w--
		buf[w] = '-'
	}

	return string(buf[w:])
}

// Since returns the date and time when the system booted in RFC1123
func (u Uptime) Since() string {
	then := time.Now().Add(-u.Uptime)
	return then.Format(time.RFC1123)
}

// SinceFormat takes a layout as input and returns the date and time
// when the system booted, according to the layout format
// FIXME: fix this and remove the gostrftime depdendency
func (u Uptime) SinceFormat(layout string) string {
	then := time.Now().Add(-u.Uptime)
	return gostrftime.Format(layout, then)
}

// Fetch updates the Uptime struct woth new values
func (u *Uptime) Fetch() error {
	if u.Sep == 0 {
		u.Sep = ':'
	}

	if u.F == nil {
		u.F = unix.Sysinfo
	}

	si := unix.Sysinfo_t{}
	if err := u.F(&si); err != nil {
		return err
	}

	u.Uptime = time.Duration(si.Uptime) * time.Second
	return nil
}

// formatFrac formats the fraction of v/10**prec (e.g., ".12345") into the
// tail of buf, omitting trailing zeros. It omits the decimal
// point too when the fraction is 0. It returns the index where the
// output bytes begin and the value v/10**prec.
func formatFrac(buf []byte, v uint64, prec int) (nw int, nv uint64) {
	// Omit trailing zeros up to and including decimal point.
	w := len(buf)
	print := false
	for i := 0; i < prec; i++ {
		digit := v % 10
		print = print || digit != 0
		if print {
			w--
			buf[w] = byte(digit) + '0'
		}
		v /= 10
	}
	if print {
		w--
		buf[w] = '.'
	}
	return w, v
}

// formatInt formats v into the tail of buf.
// It returns the index where the output begins.
func formatInt(buf []byte, v uint64) int {
	w := len(buf)
	if v == 0 {
		w--
		buf[w] = '0'
	} else {
		for v > 0 {
			w--
			buf[w] = byte(v%10) + '0'
			v /= 10
		}
	}
	return w
}

func New(sep ...byte) (*Uptime, error) {
	uptime := &Uptime{}

	if len(sep) != 0 {
		uptime.Sep = sep[0]
	}

	if err := uptime.Fetch(); err != nil {
		return &Uptime{}, err
	}

	return uptime, nil
}
