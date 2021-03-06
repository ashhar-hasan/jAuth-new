package address

import (
	"encoding/json"

	utilhttp "github.com/jabong/florest-core/src/common/utils/http"
	gm "github.com/onsi/gomega"
)

func validateHealthCheckResponse(responseBody string) {
	var utilResponse utilhttp.Response
	err := json.Unmarshal([]byte(responseBody), &utilResponse)
	gm.Expect(err).To(gm.BeNil())

	utilResponseData := utilResponse.Data
	if v, ok := utilResponseData.(map[string]interface{}); ok {
		node, addressNodePresent := v["address"]
		gm.Expect(addressNodePresent).To(gm.Equal(true))
		if addressNodePresent {
			if body, ok := node.(map[string]interface{}); ok {
				status, statusPresent := body["status"]
				gm.Expect(statusPresent).To(gm.Equal(true))
				gm.Expect(status).To(gm.Equal("success"))
			}
		}
	}
}

func GetRequestBody(response string) (utilhttp.Response, error) {
	var utilResponse utilhttp.Response
	err := json.Unmarshal([]byte(response), &utilResponse)
	return utilResponse, err
}
