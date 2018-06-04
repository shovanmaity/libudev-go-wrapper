package fdutil

import "syscall"

//FD_SET add a given file descriptor from  a  set
func FD_SET(p *syscall.FdSet, i int) {
	p.Bits[i/64] |= 1 << (uint(i) % 64)
}

//FD_ISSET tests to see if a file descriptor is part of the set
func FD_ISSET(p *syscall.FdSet, i int) bool {
	return (p.Bits[i/64] & (1 << (uint(i) % 64))) != 0
}

//FD_ZERO clears a set
func FD_ZERO(p *syscall.FdSet) {
	for i := range p.Bits {
		p.Bits[i] = 0
	}
}
