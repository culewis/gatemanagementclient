/***********************************************************************************
 * a simple client that will be used to exercise/prototype communicating with an 
 * etcd server;
 **********************************************************************************/
package gatemanagementclient


import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"fmt"
)


// the location of the etcd server
const gatemanagerurl string = "http://Johns-Macbook-Air.local:4001/v1/keys"


/**
 * attempts to retrieve the value for the passed in "key" from the etcd server;
 * this method returns the EtcdMessage contents and error (if any)
 **/
func GetKey(key string) (*EtcdMessage, string) {
	response, err := http.Get(fmt.Sprintf("%s/%s", gatemanagerurl, key))
	return ProcessResponse(response, err);
}


/**
 * attempts to create a key/value combination on the etcd server; this method
 * returns the EtcdMessage contents and error (if any)
 **/
func PostKey(key string, value string) (*EtcdMessage, string) {
	// response, err := http.Post(fmt.Sprintf("%s/%s" gatemanagerurl, key),
	// 						   url.Values{"value": {value}})
	// return ProcessResponse(response, err);
	return nil, "this method crapped its pants! ha!"
}


/**
 * a "helper" method used to process a response from the etcd server; this method
 * returns the EtcdMessage contents and an error (if any)
 **/
func ProcessResponse(response *http.Response, err error) (*EtcdMessage, string) {
	log.Printf("response Status: %s\n", response.Status)

	if err != nil {
		response.Body.Close()
		return nil, err.Error()
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err.Error()
	} else if body == nil || len(body) == 0 {
		return nil, "response body was empty"
	}

	log.Printf("response body is = %s\n", body)

	var m EtcdMessage
	err  = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err.Error()
	}

	log.Printf("EtcdMessage is %+v\n", m)
	return &m, ""
}