package APIHost

type OperationRequest struct {
	PassengerID string   `json:"passengerId"`
	ServiceName string   `json:"serviceName"`
	MethodName  string   `json:"methodName"`
	Parameters  []string `json:"parameters"`
}

type OperationResponse struct {
	Result interface{} `json:"result"`
	Error  string      `json:"error"`
}
