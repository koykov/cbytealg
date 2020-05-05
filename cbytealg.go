package cbytealg

func init() {
	RegisterAnyToBytesFn(BytesToBytes)
	RegisterAnyToBytesFn(StrToBytes)
	RegisterAnyToBytesFn(BoolToBytes)
	RegisterAnyToBytesFn(IntToBytes)
	RegisterAnyToBytesFn(UintToBytes)
	RegisterAnyToBytesFn(FloatToBytes)
}
