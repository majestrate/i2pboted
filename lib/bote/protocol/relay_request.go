package protocol

type RelayRequest struct {
	ID       CID
	Delay    uint32
	Next     Destination
	Return   ReturnChain
	Data     []byte
	HashCash []byte
}
