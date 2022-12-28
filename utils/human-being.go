package utils

type HumanBeing interface {
	Stop()
	Move()
	Talk()
	Speed() float32
}