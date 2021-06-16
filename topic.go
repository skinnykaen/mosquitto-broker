package mqtt

type Topic struct {
	Id uint `json:"id"`
	Id_User uint `json:"id_user"`
	TopicData TopicData `json:"topic_data"`
}

type TopicData struct {
	Name string `json:"topicname" binding:"required"`
	Password string `json:"passwordtopic"`
}