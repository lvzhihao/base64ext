package base64ext

type Base64Ext interface {
	Encode([]byte) string
	Decod(string) ([]byte, error)
}
