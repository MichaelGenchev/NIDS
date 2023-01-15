package sbd



type SignatureStorage interface {
	FindAll() ([]*Signature, error)
}