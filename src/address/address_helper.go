package address

import (
	"common/appconstant"
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"strconv"

	logger "github.com/jabong/florest-core/src/common/logger"
	"github.com/jabong/florest-core/src/common/profiler"
	utilHttp "github.com/jabong/florest-core/src/common/utils/http"
	"github.com/jabong/florest-core/src/components/cache"
	workflow "github.com/jabong/florest-core/src/core/common/orchestrator"
)

//DebugInfo represents a message
type DebugInfo struct {
	Key   string
	Value string
}

//Debug represents response
type Debug struct {
	MessageStack []DebugInfo
}

//addDebugContents add debug contents in the workflow data
func addDebugContents(io workflow.WorkFlowData, debug *Debug) {
	for k, v := range debug.MessageStack {
		io.ExecContext.SetDebugMsg(v.Key, v.Value)
		logger.Info(fmt.Sprintf("params.DebugInfo %v- %v", k, v))
	}
}

//urlEncode url encode the given string with data
func urlEncode(reqURL string, data []string) string {
	var URL *url.URL
	URL, _ = url.Parse(reqURL)
	parameters := url.Values{}
	for _, v := range data {
		parameters.Add("q", v)
	}
	URL.RawQuery = parameters.Encode()
	reqURL = URL.String()
	return reqURL
}

//Decrypt to decrypt an encrypted string
func Decrypt(encryptedData []string, debugInfo *Debug) []string {
	p := profiler.NewProfiler()
	p.StartProfile("AddressHelper#Decrypt")

	defer func() {
		p.EndProfileWithMetric([]string{"AddressHelper#Decrypt"})
	}()

	var (
		err               error
		res               []byte
		partialData, data []string
	)
	batchSize := int(appconstant.BatchSize)
	length := len(encryptedData)

	if length > batchSize {
		modulo := length % batchSize
		loop := math.Ceil(float64(length) / appconstant.BatchSize)
		loops := int(loop)
		for i := 0; i < loops; i++ {
			if modulo != 0 && i == (loops-1) {
				partialData = encryptedData[i*batchSize : (i*batchSize)+modulo]
			} else {
				partialData = encryptedData[i*batchSize : (i*batchSize)+batchSize]
			}
			res, err = encryptServiceObj.DecryptData(partialData, debugInfo)
			if err != nil {
				logger.Error("Decrypt: PartialResponse:: Data Decryption Error ", err.Error())
				for k := 0; k < len(partialData); k++ {
					data = append(data, "")
				}
			} else {
				d, _ := getDataFromServiceResponse(res)
				data = append(data, d...)
			}

		}
	} else {
		res, err = encryptServiceObj.DecryptData(encryptedData, debugInfo)
		if err != nil {
			logger.Error("Decrypt: Data Decryption Error ", err.Error())
		}
		data, err = getDataFromServiceResponse(res)
		if err != nil {
			data = append(data, "")
			logger.Error(fmt.Sprintf("Decrypt: getDataFromServiceResponse() Error:: %+v", err))
		}
	}

	return data
}

//getDataFromServiceResponse to parse the encryption/decryption service response
func getDataFromServiceResponse(body []byte) (data []string, err error) {
	p := profiler.NewProfiler()
	p.StartProfile("AddressHelper#getDataFromServiceResponse")

	defer func() {
		p.EndProfileWithMetric([]string{"AddressHelper#getDataFromServiceResponse"})
	}()

	r := utilHttp.Response{}
	err = json.Unmarshal(body, &r)

	if err != nil {
		logger.Error(fmt.Sprintf("decodeJson error is = %v", err))
		return data, err
	}

	tempData, _ := r.Data.([]interface{})
	var res []string
	for _, val := range tempData {
		d, ok := val.(string)
		if d == "" {
			res = append(res, "0")
		} else {
			res = append(res, d)
		}
		if !ok {
			logger.Error(fmt.Sprintf("Type Assertion fail. Error: %v, Data: %v", err, data))
		}
	}
	return res, nil
}

func decryptEncryptedFields(ef []EncryptedFields, params *RequestParams, debug *Debug) []DecryptedFields {
	rc := params.RequestContext
	var (
		encryptedPhoneString    []string
		encryptedAltPhoneString []string
		dp, dap                 int64
		err                     error
	)
	for _, v := range ef {
		encryptedPhoneString = append(encryptedPhoneString, v.EncryptedPhone)
		encryptedAltPhoneString = append(encryptedAltPhoneString, v.EncryptedAlternatePhone)
	}
	decryptedPhone := Decrypt(encryptedPhoneString, debug)
	decryptedAltPhone := Decrypt(encryptedAltPhoneString, debug)
	res := make([]DecryptedFields, 0)
	for k, v := range ef {
		if decryptedPhone[k] != "" {
			dp, err = strconv.ParseInt(decryptedPhone[k], 10, 64)
			if err != nil {
				logger.Error(fmt.Sprintf("Can not parse 'phone': %s for AddressId: %d into int64. ERROR:%v", v.EncryptedPhone, v.Id, err), rc)
			}
		} else {
			dp = 0
		}

		if decryptedAltPhone[k] != "" {
			dap, err = strconv.ParseInt(decryptedAltPhone[k], 10, 64)
			if err != nil {
				logger.Error(fmt.Sprintf("Can not parse 'alternate phone': %s for AddressId: %d into int64. ERROR:%v", v.EncryptedAlternatePhone, v.Id, err), rc)
			}
		} else {
			dap = 0
		}
		res = append(res, DecryptedFields{Id: v.Id, DecryptedPhone: dp, DecryptedAlternatePhone: dap})
	}
	return res
}

func mergeDecryptedFieldsWithAddressResult(ef []DecryptedFields, address *[]AddressResponse) {
	val := (*address)
	for k := range val {
		temp := &val[k]
		if temp.Id == ef[k].Id {
			temp.Phone = ef[k].DecryptedPhone
			temp.AlternatePhone = ef[k].DecryptedAlternatePhone
		}
	}
}

//GetAddressListCacheKey return the cache key to get/set user addresses
func GetAddressListCacheKey(userID string) string {
	return fmt.Sprintf(appconstant.AddressCacheKey, userID)
}

//getAddressListFromCache get user's address list from cache
func getAddressListFromCache(userId string, params QueryParams, debugInfo *Debug) ([]AddressResponse, error) {
	p := profiler.NewProfiler()
	p.StartProfile("AddressHelper#getAddressListFromCache")

	defer func() {
		p.EndProfileWithMetric([]string{"AddressHelper#getAddressListFromCache"})
	}()

	var address []AddressResponse

	cacheObj, errG := cache.Get(cache.Redis)
	if errG != nil {
		msg := fmt.Sprintf("Redis Config Error - %v", errG)
		logger.Error(msg)
	}
	addressListCacheKey := GetAddressListCacheKey(userId)

	debugInfo.MessageStack = append(debugInfo.MessageStack, DebugInfo{Key: "getAddressListFromCache:addressListCacheKey", Value: addressListCacheKey})
	result, err := cacheObj.Get(addressListCacheKey, false, false)
	if err != nil {
		logger.Error(fmt.Sprintf("Error while getting address list from cache %s", err))
		debugInfo.MessageStack = append(debugInfo.MessageStack, DebugInfo{Key: "getAddressListFromCache:Error", Value: "Error while getting address list from cache::" + err.Error()})
		return address, err
	}

	var addressList []AddressResponse
	data := result.Value.([]byte)
	byt := []byte(data)
	if err := json.Unmarshal(byt, &addressList); err != nil {
		debugInfo.MessageStack = append(debugInfo.MessageStack, DebugInfo{Key: "getAddressListFromCache.UnmarshalErr", Value: err.Error()})
		return address, err
	}
	debugInfo.MessageStack = append(debugInfo.MessageStack, DebugInfo{Key: "getAddressListFromCache:Result", Value: fmt.Sprintf("%+v", addressList)})

	return addressList, nil
}

//saveDataInCache save data in cache
func saveDataInCache(id string, ty string, value interface{}) error {
	p := profiler.NewProfiler()
	p.StartProfile("AddressHelper#getAddressListFromCache")

	defer func() {
		p.EndProfileWithMetric([]string{"AddressHelper#getAddressListFromCache"})
	}()

	cacheObj, errG := cache.Get(cache.Redis)
	if errG != nil {
		msg := fmt.Sprintf("Redis Config Error - %v", errG)
		logger.Error(msg)
	}

	var cacheKey string
	if ty == "address" {
		cacheKey = GetAddressListCacheKey(id)
	} else if ty == "locality" {
		cacheKey = ""
	}
	str, _ := json.Marshal(value)

	item := cache.Item{}
	item.Key = cacheKey
	item.Value = string(str)

	err := cacheObj.Set(item, false, false)
	if err != nil {
		return err
	}

	return nil
}

//invalidateCache invalidate cache key
func invalidateCache(key string) error {
	cacheObj, errG := cache.Get(cache.Redis)
	if errG != nil {
		msg := fmt.Sprintf("Redis Config Error - %v", errG)
		logger.Error(msg)
	}

	err := cacheObj.Delete(key)
	if err != nil {
		return err
	}
	return nil
}
