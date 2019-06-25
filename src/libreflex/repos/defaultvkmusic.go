package repos

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type DefaultVkMusicRepo struct {
	httpClient *http.Client
	vkApiToken string
}

func NewDefaultVkMusicRepo(vkApiToken string) *DefaultVkMusicRepo {
	repo := &DefaultVkMusicRepo{vkApiToken: vkApiToken}
	repo.httpClient = &http.Client{Timeout: 15 * time.Second}
	return repo
}

func (r *DefaultVkMusicRepo) GetAudio(audioIds []string) (map[string]interface{}, error) {
	uri, err := url.Parse("https://api.vk.com/method/audio.getById")
	if err != nil {
		return nil, err
	}

	query := uri.Query()
	query.Set("access_token", r.vkApiToken)
	query.Set("v", "5.82")
	query.Set("audios", strings.Join(audioIds, ","))
	uri.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "VKAndroidApp/5.25-3000 (Android 6.0.0; SDK 23; armeabi-v7a; LGE Nexus 5; ru)")

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	jsn := make(map[string]interface{})
	err = json.Unmarshal(body, &jsn)
	if err != nil {
		return nil, err
	}

	return jsn, nil
}
