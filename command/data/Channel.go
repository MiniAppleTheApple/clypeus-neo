package data

type Channel struct {
	id string
}

func (channel *Channel) GetID() string {
	return channel.id
}
