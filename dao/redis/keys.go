package redis

const (
	Prefix             = "yanblue:"
	KeyPostTimeZSet    = "post:time"
	KeyPostScoreZSet   = "post:score"
	KeyPostVotedZSetPF = "post:voted:"

	KeyCommunitySetPF = "community:"
)

// get redis key
func getRedisKey(key string) string {
	return Prefix + key
}
