package protocol

type DhtType byte

const RetrieveReq = DhtType(0x51)
const DeleteQuery = DhtType(0x4c)
const StoreReq = DhtType(0x53)
const DelEmailReq = DhtType(0x44)
const IndexDelReq = DhtType(0x58)
const FindClose = DhtType(0x43)
