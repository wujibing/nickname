package nickname

import (
	"math/rand"
)

func Get() string {
	rand.Seed(rand.Int63())
	return adjDic[rand.Intn(len(adjDic) - 1)] + nounDic[rand.Intn(len(nounDic) - 1)]
}