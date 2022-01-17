package sdk

import (
	`encoding/json`
	`strconv`

	`github.com/ttlv/kdniaoGo`
	`github.com/ttlv/kdniaoGo/enum`
	`github.com/ttlv/kdniaoGo/request`
	`github.com/ttlv/kdniaoGo/response`
	`github.com/ttlv/kdniaoGo/util`
	`github.com/ttlv/kdniaoGo/util/http`
)

func NewExpressQuery(config kdniaoGo.KdniaoConfig, logger kdniaoGo.KdniaoLoggerInterface) ApiExpressQuery {
	return ApiExpressQuery{config, logger}
}

type ApiExpressQuery struct {
	config kdniaoGo.KdniaoConfig
	logger kdniaoGo.KdniaoLoggerInterface
}

func (obj ApiExpressQuery) GetRequest(logisticCode, shipperCode string) request.ExpressQueryRequest {
	req := request.NewExpressQueryRequest()
	req.SetConfig(obj.config)
	req.SetLogisticCode(logisticCode)
	return req
}

func (obj ApiExpressQuery) GetResponse(req request.ExpressQueryRequest) (response.ExpressQueryResponse, error) {
	url := enum.GATEWAY + enum.URI_BUSINESS

	req.UpdateRequestData()
	var resp response.ExpressQueryResponse
	httpResp, err := http.HttpPostForm(url, req.ToValues(), obj.logger)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiExpressQuery,http fail")
	} else if !httpResp.IsOk() {
		return resp, util.ErrorNew("ApiExpressQuery,code:" + strconv.Itoa(httpResp.GetCode()))
	}
	err = json.Unmarshal(httpResp.GetBytes(), &resp)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiExpressQuery,json decode fail")
	}
	if !resp.IsSuccess() {
		return resp, util.ErrorNew("ApiExpressQuery," + resp.GetError())
	}
	return resp, nil
}