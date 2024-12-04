package uuid

import ("github.com/google/uuid")

func ShortUUID() string {
	return uuid.New().String()[:8]
}

func LongUUID() string {
	return uuid.New().String()
}