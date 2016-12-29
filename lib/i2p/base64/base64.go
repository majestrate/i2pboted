package base64

import base "encoding/base64"

// i2p base64 encoding
var Encoding *base.Encoding = base.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-~")
