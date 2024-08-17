package mpesa

const (
	PayBillOnline   = "CustomerPayBillOnline"
	QueryStatus     = "TransactionStatusQuery"
	BusinessPayment = "BusinessPayment"
	Reversal        = "TransactionReversal"
	Accepted        = "Accepted"
	Rejected        = "Rejected"
	C2B00011        = "Invalid MSISDN"
	C2B00012        = "Invalid Account Number"
	C2B00013        = "Invalid Amount"
	C2B00014        = "Invalid KYC Details"
	C2B00015        = "Invalid Shortcode"
	C2B00016        = "Other Error"
)

type Credentials struct {
	Username, Password string
}

type TokenResult struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type StkpushResponse struct {
	MerchantRequestId   string `json:"MerchantRequestId"`
	CheckoutRequestId   string `json:"CheckoutRequestId"`
	ResponseCode        string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	CustomerMessage     string `json:"CustomerMessage"`
}

type StkPushRequestError struct {
	RequestId    string `json:"RequestId"`
	ErrorCode    string `json:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage"`
}

type StkInitiationPayload struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	Amount            string `json:"Amount"`
	PartyA            string `json:"PartyA"`
	PartyB            string `json:"PartyB"`
	PhoneNumber       string `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	AccountReference  string `json:"AccountReference"`
	TransactionDesc   string `json:"TransactionDesc"`
}

type Item struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type CallbackMetadata struct {
	Item []Item `json:"Item"`
}

type StkCallback struct {
	MerchantRequestID string           `json:"MerchantRequestId"`
	CheckoutRequestID string           `json:"CheckoutRequestId"`
	ResultCode        string           `json:"ResultCode"`
	ResultDesc        string           `json:"ResultDesc"`
	CallbackMetadata  CallbackMetadata `json:"CallbackMetadata"`
}

type ResultParameter struct {
	Key   string `json:"Key"`
	Value string `json:"Value,omitempty"`
}

type StatusResult struct {
	ConversationID           string `json:"ConversationID"`
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ReferenceData            struct {
		ReferenceItem struct {
			Key string `json:"Key"`
		} `json:"ReferenceItem"`
	} `json:"ReferenceData"`
	ResultCode       int    `json:"ResultCode"`
	ResultDesc       string `json:"ResultDesc"`
	ResultParameters struct {
		ResultParameter []ResultParameter `json:"ResultParameter"`
	} `json:"ResultParameters"`
	ResultType    int    `json:"ResultType"`
	TransactionID string `json:"TransactionID"`
}

type StatusResponse struct {
	Result StatusResult
}

type Body struct {
	StkCallback StkCallback `json:"StkCallback"`
}

type StkCallBack struct {
	Body Body `json:"Body"`
}

type StatusRequest struct {
	Initiator                string `json:"Initiator"`
	SecurityCredential       string `json:"SecurityCredential"`
	OriginatorConversationID string `json:"OriginatorConversationID"`
	CommandID                string `json:"CommandID"`
	TransactionID            string `json:"TransactionID"`
	PartyA                   string `json:"PartyA"`
	IdentifierType           string `json:"IdentifierType"`
	ResultURL                string `json:"ResultURL"`
	QueueTimeOutURL          string `json:"QueueTimeOutURL"`
	Remarks                  string `json:"Remarks"`
	Occasion                 string `json:"Occasion"`
}

type C2BRequest struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	InitiatorName            string `json:"InitiatorName"`
	SecurityCredential       string `json:"SecurityCredential"`
	CommandID                string `json:"CommandID"`
	Amount                   string `json:"Amount"`
	PartyA                   string `json:"PartyA"`
	PartyB                   string `json:"PartyB"`
	Remarks                  string `json:"Remarks"`
	QueueTimeOutURL          string `json:"QueueTimeOutURL"`
	ResultURL                string `json:"ResultURL"`
	Occasion                 string `json:"Occasion"`
}

type C2BResponse struct {
	ConversationID           string `json:"ConversationID"`
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}

type ResultPayload struct {
	ResultCode string `json:"ResultCode"`
	ResultDesc string `json:"ResultDesc"`
}

type ReversalRequest struct {
	Initiator              string `json:"Initiator"`
	SecurityCredential     string `json:"SecurityCredential"`
	CommandID              string `json:"CommandID"`
	TransactionID          string `json:"TransactionID"`
	Amount                 string `json:"Amount"`
	ReceiverParty          string `json:"ReceiverParty"`
	RecieverIdentifierType string `json:"RecieverIdentifierType"`
	ResultURL              string `json:"ResultURL"`
	QueueTimeOutURL        string `json:"QueueTimeOutURL"`
	Remarks                string `json:"Remarks"`
	Occasion               string `json:"Occasion"`
}

type ReversalResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}
