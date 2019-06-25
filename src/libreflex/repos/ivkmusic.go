package repos

type VkMusicRepo interface {
	//Returns JSON response of VK API method audio.getById
	GetAudio(audioIds []string) (map[string]interface{}, error)
}
