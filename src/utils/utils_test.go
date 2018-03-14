package utils_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"utils"
)

func TestInArrayRegexp(t *testing.T) {

	var RegexpArr = []string{".*?consul.*", "infra1"}

	assert.Equal(t, true, utils.InArrayRegexp("consul", RegexpArr))
	assert.Equal(t, true, utils.InArrayRegexp("consul-infra1-02", RegexpArr))
	assert.Equal(t, true, utils.InArrayRegexp("consul-05", RegexpArr))
	assert.Equal(t, true, utils.InArrayRegexp("huembuem-consul-1", RegexpArr))
	assert.Equal(t, true, utils.InArrayRegexp("dfconsulgb", RegexpArr))
	assert.Equal(t, true, utils.InArrayRegexp("infra1", RegexpArr))
	assert.Equal(t, false, utils.InArrayRegexp("siam-test", RegexpArr))

}


func TestParseDcService(t *testing.T) {

	dc1, service1 := utils.ParseDcService("[billingsn]billing")
	dc2, service2 := utils.ParseDcService("[irontrade-sg]irontrade_sing-pgsql")

	assert.Equal(t, "billingsn", dc1)
	assert.Equal(t, "irontrade-sg", dc2)

	assert.Equal(t, "billing", service1)
	assert.Equal(t, "irontrade_sing-pgsql", service2)
}