package dht

type PacketType byte

const RetrieveReq = PacketType(0x51)
const DeleteQuery = PacketType(0x4c)
const StoreReq = PacketType(0x53)
const DelEmailReq = PacketType(0x44)
const IndexDelReq = PacketType(0x58)
const FindClose = PacketType(0x43)
