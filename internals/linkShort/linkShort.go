package linkShort

type LinkFull struct {
	Link string `json:"link"`
}

type LinkShort struct {
	Link string
}

type ResponseLinkShort struct {
	LinkShort LinkShort `json:"link_short"`
}
