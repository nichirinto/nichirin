package pingdto

type PingInputDto struct{}

type PingOutputDto struct {
	Time int64 `json:"time"`
}
