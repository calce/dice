// random 64 bit ID generator
package dice

import (
	"net"
	"time"
	"math/rand"
	"sync/atomic"
	"encoding/hex"
	cr "crypto/rand"
	"encoding/binary"
)

type Dice struct {
	atom int64
	mac [8]byte
}

var defaultDice *Dice = nil

func (this *Dice) spin() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	atomic.AddInt64(&this.atom, r.Int63())
	cr.Read(this.mac[:])
	this.mac[0] |= 0x01
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, iface := range interfaces {
			if len(iface.HardwareAddr) >= 6 {
				copy(this.mac[:], iface.HardwareAddr)				
				break
			}
		}
	}
}

func (this *Dice) random() string {	
	seed := time.Now().UnixNano()
	atomic.AddInt64(&this.atom, seed)
	r := rand.New(rand.NewSource(seed*this.atom))
	i := r.Uint32()	
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b[:], i)
	return hex.EncodeToString(b)
}

func (this *Dice) fromMAC() string {
	mac := this.mac
	cr.Read(mac[6:])
	b := []byte{mac[0]*mac[4], mac[1]*mac[5], mac[2]*mac[6], mac[3]*mac[7]}
	return hex.EncodeToString(b)
}

func getDice() *Dice {
	if defaultDice == nil { 
		defaultDice = &Dice{
			atom: 0,
		}
		defaultDice.spin() 
	}
	return defaultDice
}

// generates a random ID
func Spin() string {
	return getDice().random()
}

// generates a less random ID using MAC address
func LessSpin() string {
	return getDice().fromMAC()	
}