package session

import "time"

const defaultSessionKeyName = "dj"
const defaultDomain = ""
const defaultExpires = 2 * time.Hour

//todo 开发阶段设置为一个星期,后续改回来
const defaultGCLifetime = 24 * 7 * time.Hour

//const defaultGCLifetime = 60 * time.Second
const defaultSecure = false
const defaultSessionIDInURLQuery = false
const defaultSessionIDInHTTPHeader = false
const defaultCookieLen uint32 = 32

const expirationAttributeKey = "_sid_"
