package servicetest

import (
	"github.com/jabong/florest-core/src/common/logger"
)

func initTestLogger() {
	logger.Initialise("../../../config/testdata/testloggerAsync.json")
}
