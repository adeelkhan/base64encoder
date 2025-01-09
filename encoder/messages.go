package encoder

// messages
type RequestEncode struct {
	PlainText string
}

type ResponseEncode struct {
	Base64EncodedString string
}

type RequestDecode struct {
	Base64EncodedString string
}

type ResponseDecode struct {
	PlainText string
}
type ResponseError struct {
	Message string
}
