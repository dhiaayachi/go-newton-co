package newton

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dhiaayachi/go-newton-co/query"
)

const baseUrl = "https://api.newton.co/v1"

type Newton struct {
	clientId     string
	clientSecret string
}

func New(ClientId string, ClientSecret string) *Newton {
	return &Newton{ClientId, ClientSecret}
}

func (n *Newton) sign(req *http.Request) error {
	var (
		method      string
		contentType string
		apiPath     string
		hashedBody  string
		currentTime string
	)

	method = req.Method
	if req.Method != http.MethodGet {
		contentType = "application/json"
	}

	apiPath = req.URL.Path
	if req.Body != http.NoBody {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return fmt.Errorf("read all body: %w", err)
		}

		hash := sha256.Sum256(b)
		hashedBody = fmt.Sprintf("%x", hash[:])

		req.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	}

	currentTime = strconv.FormatInt(time.Now().Unix(), 10)

	toJoin := []string{
		method,
		contentType,
		apiPath,
		hashedBody,
		currentTime,
	}

	raw := strings.Join(toJoin, ":")

	mac := hmac.New(sha256.New, []byte(n.clientSecret))
	if _, err := mac.Write([]byte(raw)); err != nil {
		return fmt.Errorf("mac write: %w", err)
	}

	signedBase64 := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	req.Header.Add("NewtonDate", currentTime)
	req.Header.Add("NewtonAPIAuth", n.clientId+":"+signedBase64)

	return nil
}

func (n *Newton) Do(query query.Query) (interface{}, error) {
	body, err := query.GetBody()
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(
		query.GetMethod(),
		baseUrl + query.GetPath(), 
		bytes.NewBuffer(body))
	q := req.URL.Query()
	for _, a := range query.GetParameters() {
		q.Add(a.Key, a.Value)
	}
	req.URL.RawQuery = q.Encode()
	if query.GetMethod() != http.MethodGet {
		req.Header.Add("content-type", "application/json")
	}

	if !query.IsPublic() {
		err := n.sign(req)
		if err != nil {
			return nil, err
		}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	parsedResponse := query.GetResponse()
	err = n.parseResponse(res, parsedResponse)
	if err != nil {
		return nil, err
	}

	return parsedResponse, err
}

func (n *Newton) parseResponse(res *http.Response, toParseTo interface{}) error {
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Printf("error:%s", err.Error())
		}
	}()

	if toParseTo == nil {
		return nil
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed :: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed :: %d :: %s", res.StatusCode, body)
	}

	err = json.Unmarshal(body, toParseTo)
	if err != nil {
		return err
	}

	return nil
}
