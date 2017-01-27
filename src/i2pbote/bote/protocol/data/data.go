package data

type Type byte

const EmailEncrypted = Type(0x45)
const EmailUnecrypted = Type(0x55)
const Index = Type(0x49)
const DeleteInfo = Type(0x54)
const PeerList = Type(0x50)
const DirectoryEntry = Type(0x43)

func (t Type) Byte() byte {
	return byte(t)
}

func New(t Type, version byte, buff []byte) (d []byte) {
	d = make([]byte, 1+1+len(buff))
	d[0] = t.Byte()
	d[1] = version
	copy(d[2:], buff)
	return
}
