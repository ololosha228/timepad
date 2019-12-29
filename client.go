package timepad

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/mitchellh/mapstructure"
)

func GetEvents() ([]*Event, error) {
	link := url.URL{
		Scheme: "https",
		Host:   "api.timepad.ru",
		Path:   "/v1/events",
	}
	link.RawQuery = url.Values{
		"fields":              []string{"location"},
		"limit":               []string{"10"},
		"sort":                []string{"+starts_at"},
		"cities":              []string{"Москва"},
		"moderation_statuses": []string{"featured"},
	}.Encode()

	println(link.String())
	res, err := http.Get(link.String())
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	total := struct {
		Total  int      `json:"total`
		Values []*Event `json:"values"`
	}{}

	err = json.Unmarshal(body, &total)
	if err != nil {
		return nil, err
	}

	return total.Values, nil
}

type Event struct {
	ID       int       `json:"id"`
	StartsAt time.Time `json:"starts_at"`
	Name     string    `json:"name"`
	URL      string    `json:"url"`
	Category string    `json:"categories"`
}

func (e *Event) UnmarshalJSON(b []byte) error {
	if e == nil {
		e = new(Event)
	}

	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02T15:04:05-0700", m["starts_at"].(string))
	if err != nil {
		return err
	}
	m["starts_at"] = t

	m["categories"] = m["categories"].([]interface{})[0].(map[string]interface{})["name"].(string)

	d, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  e,
	})

	old := e
	err = d.Decode(&m)
	if err != nil {
		e = old
		return err
	}
	return nil
}
