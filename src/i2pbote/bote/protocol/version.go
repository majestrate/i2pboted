package protocol

type ProtocolVersion byte

func (v ProtocolVersion) Byte() byte {
	return byte(v)
}

const Version = ProtocolVersion(4)
