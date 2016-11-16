package sysinfo

import (
	"time"

	"github.com/cactus/gostrftime"

	"golang.org/x/sys/unix"
)

type Uptime struct {
	Uptime time.Duration
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

	w, up = fmtFrac(buf[:w], up, 9)

	// up is now integer seconds
	w = fmtInt(buf[:w], up%60)
	up /= 60

	// up is now integer minutes
	if up > 0 {
		w--
		buf[w] = 'm'
		w = fmtInt(buf[:w], up%60)
		up /= 60

		// up is now integer hours
		if up > 0 {
			w--
			buf[w] = 'h'
			w = fmtInt(buf[:w], up%24)
			up /= 24
		}

		// up is now integer days
		if up > 0 {
			w--
			buf[w] = 'd'
			w = fmtInt(buf[:w], up)
		}
	}

	if neg {
		w--
		buf[w] = '-'
	}

	return string(buf[w:])
}

func (u Uptime) UpSince() string {
	then := time.Now().Add(-u.Uptime)
	return then.Format(time.RFC1123)
}

func (u Uptime) UpSinceFormat(layout string) string {
	then := time.Now().Add(-u.Uptime)
	return gostrftime.Format(layout, then)
}

func (u *Uptime) Get() error {
	si := unix.Sysinfo_t{}
	if err := unix.Sysinfo(&si); err != nil {
		return err
	}

	u.Uptime = time.Duration(si.Uptime) * time.Second
	return nil
}

// fmtFrac formats the fraction of v/10**prec (e.g., ".12345") into the
// tail of buf, omitting trailing zeros.  it omits the decimal
// point too when the fraction is 0.  It returns the index where the
// output bytes begin and the value v/10**prec.
func fmtFrac(buf []byte, v uint64, prec int) (nw int, nv uint64) {
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

// fmtInt formats v into the tail of buf.
// It returns the index where the output begins.
func fmtInt(buf []byte, v uint64) int {
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
