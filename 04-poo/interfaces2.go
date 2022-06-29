package main

type Encrypt interface {
	Encrypt([]byte) string
}
type Decrypt interface {
	Decrypt([]byte) string
}

func Send(encrypter Encrypt) {
	// method implementation
}

func Recived(encrypter Decrypt) {
	// method implementation
}
