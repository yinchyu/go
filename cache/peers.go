package cache

type PeePicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	Get(group, key string) ([]byte, error)
}
