/***********************************************************************************
 * a simple client that will be used to exercise/prototype communicating with an 
 * etcd server
 **********************************************************************************/
package gatemanagementclient


import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"log"
)


// the location of the etcd server
const gatemanagerurl string = "http://Johns-Macbook-Air.local:4001/v1/keys"


/**
 * attempts to retrieve the value for the passed in "key" from the etcd server;
 * this method returns the message contents XOR a fatal error that occurred
 **/
func GetKey(key string) (interface{}, error) {
	response, httpErr := http.Get(fmt.Sprintf("%s/%s", gatemanagerurl, key))
	return ProcessResponse(response, httpErr)
}


/**
 * attempts to create a key/value combination on the etcd server; this method
 * returns the message contents XOR a fatal error that occurred
 **/
func PostKey(key string, value string) (interface{}, error) {
	response, err := http.PostForm(fmt.Sprintf("%s/%s", gatemanagerurl, key),
								   url.Values{"value": {value}})
	return ProcessResponse(response, err)
}


/**
 * attempts to delete a key on the etcd server; this method returns the message
 * contents XOR a fatal error that occurred; since there is no Delete function
 * on the http type (annoying) a client must be constructed and a new request 
 * must be created to explicitly make a DELETE http request
 **/
func DeleteKey(key string) (interface{}, error) {
	var client http.Client

	request, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", gatemanagerurl, key), nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	return ProcessResponse(response, err)
}


/**
 * a "helper" method used to process a response from the etcd server; this method
 * returns the message contents XOR a fatal error that occurred
 **/
func ProcessResponse(response *http.Response, err error) (interface{}, error) {
	log.Printf("response Status: %s\n", response.Status)

	if err != nil {
		response.Body.Close()
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	} else if body == nil || len(body) == 0 {
		return nil, errors.New("response body was empty")
	}

	log.Printf("response body is = %s\n", body)

	switch response.StatusCode {
	case 404:
		var m EtcdErrorMessage
		err = json.Unmarshal(body, &m)
		if err != nil {
			return nil, err
		}

		log.Printf("EtcdErrorMessage is %+v\n", m)
		return m, nil

	case 200:
		var m EtcdActionMessage
		err = json.Unmarshal(body, &m)
		if err != nil {
			return nil, err
		}

		log.Printf("EtcdActionMessage is %+v\n", m)
		return m, nil

	default:
		return nil, errors.New(fmt.Sprintf("unhandled response encountered: %s", 
										   response.Status))
	}
}