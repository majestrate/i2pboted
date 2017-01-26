package comm

type ResponseStatus byte

const OK = ResponseStatus(0)
const Err = ResponseStatus(1)
const NoData = ResponseStatus(2)
const InvalidPacket = ResponseStatus(3)
const InvalidHashcash = ResponseStatus(4)
const NotEnoughHashCash = ResponseStatus(5)

// response packet
type ResponsePacket struct {
}
