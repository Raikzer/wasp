// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import (
	"github.com/iotaledger/hive.go/logger"
	"github.com/mr-tron/base58"
)

const (
	OBJTYPE_ARRAY int32 = 0x20

	OBJTYPE_ADDRESS int32 = 1
	OBJTYPE_AGENT   int32 = 2
	OBJTYPE_BYTES   int32 = 3
	OBJTYPE_COLOR   int32 = 4
	OBJTYPE_HASH    int32 = 5
	OBJTYPE_INT     int32 = 6
	OBJTYPE_MAP     int32 = 7
	OBJTYPE_STRING  int32 = 8
)

const KeyFromString int32 = 0x4000

var HostTracing = false
var ExtendedHostTracing = false

type HostObject interface {
	Exists(keyId int32) bool
	GetBytes(keyId int32) []byte
	GetInt(keyId int32) int64
	GetObjectId(keyId int32, typeId int32) int32
	GetString(keyId int32) string
	GetTypeId(keyId int32) int32
	SetBytes(keyId int32, value []byte)
	SetInt(keyId int32, value int64)
	SetString(keyId int32, value string)
}

// KvStoreHost implements WaspLib.client.ScHost interface
// it allows access to
type KvStoreHost struct {
	keyIdToKey    [][]byte
	keyIdToKeyMap [][]byte
	keyToKeyId    map[string]int32
	log           *logger.Logger
	objIdToObj    []HostObject
	useBase58Keys bool
}

func (host *KvStoreHost) Init(null HostObject, root HostObject, log *logger.Logger) {
	host.log = log.Named("wasmtrace")
	host.log = log
	host.objIdToObj = nil
	host.keyIdToKey = [][]byte{[]byte("<null>")}
	host.keyToKeyId = make(map[string]int32)
	host.keyIdToKeyMap = make([][]byte, len(keyMap)+1)
	for k, v := range keyMap {
		host.keyIdToKeyMap[-v] = []byte(k)
	}
	host.TrackObject(null)
	host.TrackObject(root)
}

func (host *KvStoreHost) Exists(objId int32, keyId int32) bool {
	return host.FindObject(objId).Exists(keyId)
}

func (host *KvStoreHost) FindObject(objId int32) HostObject {
	if objId < 0 || objId >= int32(len(host.objIdToObj)) {
		panic("FindObject: invalid objId")
		objId = 0
	}
	return host.objIdToObj[objId]
}

func (host *KvStoreHost) FindSubObject(obj HostObject, keyId int32, typeId int32) HostObject {
	if obj == nil {
		// use root object
		obj = host.FindObject(1)
	}
	return host.FindObject(obj.GetObjectId(keyId, typeId))
}

func (host *KvStoreHost) GetBytes(objId int32, keyId int32) []byte {
	obj := host.FindObject(objId)
	if !obj.Exists(keyId) {
		host.Trace("MustGetBytes o%d k%d missing key", objId, keyId)
		return nil
	}
	value := obj.GetBytes(keyId)
	host.Trace("MustGetBytes o%d k%d = '%s'", objId, keyId, base58.Encode(value))
	return value
}

func (host *KvStoreHost) GetInt(objId int32, keyId int32) int64 {
	host.TraceAll("GetInt(o%d,k%d)", objId, keyId)
	value := host.FindObject(objId).GetInt(keyId)
	host.Trace("GetInt o%d k%d = %d", objId, keyId, value)
	return value
}

func (host *KvStoreHost) getKeyFromId(keyId int32) []byte {
	// find predefined key
	if keyId < 0 {
		return host.keyIdToKeyMap[-keyId]
	}

	// find user-defined key
	return host.keyIdToKey[keyId & ^KeyFromString]
}

func (host *KvStoreHost) GetKeyFromId(keyId int32) []byte {
	host.TraceAll("GetKeyFromId(k%d)", keyId)
	key := host.getKeyFromId(keyId)
	if (keyId & KeyFromString) == 0 {
		// originally a byte slice key
		host.Trace("GetKeyFromId k%d='%s'", keyId, base58.Encode(key))
		return key
	}
	// originally a string key
	host.Trace("GetKeyFromId k%d='%s'", keyId, string(key))
	return key
}

func (host *KvStoreHost) getKeyId(key []byte, fromString bool) int32 {
	// cannot use []byte as key in maps
	// so we will convert to (non-utf8) string
	// most will have started out as string anyway
	keyString := string(key)

	// first check predefined key map
	keyId, ok := keyMap[keyString]
	if ok {
		return keyId
	}

	// check additional user-defined keys
	keyId, ok = host.keyToKeyId[keyString]
	if ok {
		return keyId
	}

	// unknown key, add it to user-defined key map
	keyId = int32(len(host.keyIdToKey))
	if fromString {
		keyId |= KeyFromString
	}
	host.keyToKeyId[keyString] = keyId
	host.keyIdToKey = append(host.keyIdToKey, key)
	return keyId
}

func (host *KvStoreHost) GetKeyIdFromBytes(bytes []byte) int32 {
	encoded := base58.Encode(bytes)
	if host.useBase58Keys {
		// transform byte slice key into base58 string
		// now all keys are byte slices from strings
		bytes = []byte(encoded)
	}

	keyId := host.getKeyId(bytes, false)
	host.Trace("GetKeyIdFromBytes '%s'=k%d", encoded, keyId)
	return keyId
}

func (host *KvStoreHost) GetKeyIdFromString(key string) int32 {
	keyId := host.getKeyId([]byte(key), true)
	host.Trace("GetKeyIdFromString '%s'=k%d", key, keyId)
	return keyId
}

func (host *KvStoreHost) GetKeyStringFromId(keyId int32) string {
	return string(host.GetKeyFromId(keyId))
}

func (host *KvStoreHost) GetObjectId(objId int32, keyId int32, typeId int32) int32 {
	host.TraceAll("GetObjectId(o%d,k%d,t%d)", objId, keyId, typeId)
	subId := host.FindObject(objId).GetObjectId(keyId, typeId)
	host.Trace("GetObjectId o%d k%d t%d = o%d", objId, keyId, typeId, subId)
	return subId
}

func (host *KvStoreHost) GetString(objId int32, keyId int32) string {
	value := host.getString(objId, keyId)
	if value == nil {
		return ""
	}
	return *value
}

func (host *KvStoreHost) getString(objId int32, keyId int32) *string {
	obj := host.FindObject(objId)
	if !obj.Exists(keyId) {
		host.Trace("MustGetString o%d k%d missing key", objId, keyId)
		return nil
	}
	value := obj.GetString(keyId)
	host.Trace("MustGetString o%d k%d = '%s'", objId, keyId, value)
	return &value
}

func (host *KvStoreHost) PopFrame(frame []HostObject) {
	host.objIdToObj = frame
}

func (host *KvStoreHost) PushFrame() []HostObject {
	// reset frame to contain only null and root object
	// create a fresh slice to allow garbage collection
	// it's up to the caller to save and/or restore the old frame
	pushed := host.objIdToObj
	host.objIdToObj = make([]HostObject, 2)
	copy(host.objIdToObj, pushed[:2])
	return pushed
}

func (host *KvStoreHost) SetBytes(objId int32, keyId int32, bytes []byte) {
	host.FindObject(objId).SetBytes(keyId, bytes)
	host.Trace("SetBytes o%d k%d v='%s'", objId, keyId, base58.Encode(bytes))

}

func (host *KvStoreHost) SetInt(objId int32, keyId int32, value int64) {
	host.TraceAll("SetInt(o%d,k%d)", objId, keyId)
	host.FindObject(objId).SetInt(keyId, value)
	host.Trace("SetInt o%d k%d v=%d", objId, keyId, value)
}

func (host *KvStoreHost) SetString(objId int32, keyId int32, value string) {
	host.FindObject(objId).SetString(keyId, value)
	host.Trace("SetString o%d k%d v='%s'", objId, keyId, value)
}

func (host *KvStoreHost) Trace(format string, a ...interface{}) {
	if HostTracing {
		host.log.Debugf(format, a...)
	}
}

func (host *KvStoreHost) TraceAll(format string, a ...interface{}) {
	if ExtendedHostTracing {
		host.Trace(format, a...)
	}
}

func (host *KvStoreHost) TrackObject(obj HostObject) int32 {
	objId := int32(len(host.objIdToObj))
	host.objIdToObj = append(host.objIdToObj, obj)
	return objId
}
