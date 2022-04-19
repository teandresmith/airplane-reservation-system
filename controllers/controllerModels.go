package controllers

type Response struct {
	Message string      `bson:"message" json:"message"`
	Result  interface{} `bson:"result" json:"result"`
}

type Error struct {
	Message string `bson:"message" json:"message"`
	Error   string `bson:"error" json:"error"`
}