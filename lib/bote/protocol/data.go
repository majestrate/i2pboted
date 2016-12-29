package protocol

type DataType byte

const EmailEncrypted = DataType(0x45)
const EmailUnecrypted = DataType(0x55)
const Index = DataType(0x49)
const DeleteInfo = DataType(0x54)
const PeerList = DataType(0x50)
const DirectoryEntry = DataType(0x43)
