/***********************************************************************************
 * this is a simple set of tests used to exercise the functionality found in
 * gatemanagementclient.go
 **********************************************************************************/
package gatemanagementclient

import (
 	"fmt"
 	"runtime"
 	"testing"
)

const GATEMANAGERURL = "http://Johns-Macbook-Air.local:4001/v1/keys"

/**
 * tests that retrieving a key that does not exist returns the expected 
 * errorCode: 100 and message: "Key not found"
 **/
func TestGetKeyInvalidKey(t *testing.T) {
	var client = &GateManagementClient{GATEMANAGERURL}
	message, err := client.GetKey("whoopsie")

	if err != nil {
		t.Errorf("An unexpected error occurred: %s\n", err)
	} else if message == nil {
		t.Error("message was unexpectedly nil")	
	} else {
		etcdErrorMessage, ok := message.(EtcdErrorMessage)
		if !ok {
			t.Error("An unexpected type was returned from the call to GetKey")
		} else if etcdErrorMessage.ErrorCode != 100 || 
				  etcdErrorMessage.Message != "Key not found" {
			t.Errorf("EtcdErrorMessage was not of the expected type. EtcdErrorMessage was %+v\n", 
					 etcdErrorMessage)
		}
	}
}


/**
 * tests that creating a key creates the key/value combination properly without
 * encountering any errors
 **/
func TestPutKey(t *testing.T) {
	var client = &GateManagementClient{GATEMANAGERURL}
	message, err := client.PostKey("foo", "bar")

	if err != nil {
		t.Errorf("An unexpected error occurred: %s\n", err)
	} else if message == nil {
		t.Error("The configAction response for PostKey was unexpectedly nil")
	} else {
		etcdActionMessage, ok := message.(EtcdActionMessage)
		if !ok {
			t.Error("An unexpected type was returned from the call to PostKey")
		} else if etcdActionMessage.Action != "set" || etcdActionMessage.Key != "/foo" ||
				  etcdActionMessage.Value != "bar" {
		  	t.Errorf("PostKey EtcdActionMessage did not return the expected values.  " +
		  				"EtcdActionMessage was %+v\n", etcdActionMessage)
		} else {
			message, err := client.DeleteKey("foo")
			if err != nil {
				t.Errorf("An unexpected error occurred: %s\n", err)
			} else if message == nil {
				t.Error("The configAction response for DeleteKey was unexpectedly nil")
			} else {
				etcdActionMessage, ok := message.(EtcdActionMessage)
				if !ok {
					t.Error("An unexpected type was returned from the call to DeleteKey")
				} else if etcdActionMessage.Action != "delete" || etcdActionMessage.Key != "/foo" ||
						  etcdActionMessage.PrevValue != "bar" {
				  	t.Errorf("DeleteKey EtcdActionMessage did not return the expected values.  " +
				  				"EtcdActionMessage was %+v\n", etcdActionMessage)
			  	}
			}
		}
	}
}


/**
 * tests performance of creating keys and retrieving values from multiple clients performs
 * well and that the client is somewhat hardened against memory bloat or other untasteful
 * errors as a result of poor management of memory
 **/
 func TestClientIsHardened(t *testing.T) {
	const THRESHOLDOFPAIN = 512   // looked like etcd started to crack at about 1024
 	var client = &GateManagementClient{GATEMANAGERURL}
 	sem := make(chan int, THRESHOLDOFPAIN)
 	runtime.GOMAXPROCS(32)  // cap concurrent threads to 32

 	for i := 0; i < THRESHOLDOFPAIN; i++ {
 		go func(client *GateManagementClient, i int) {
	 		t.Logf("PostKey foo%v=bar%v", i, i)
	 		message, err := client.PostKey(fmt.Sprintf("foo%v", i), fmt.Sprintf("bar%v", i))

	 		if err != nil {
	 			t.Errorf("An unexpected error occurred: %s\n", err)
	 		} else if message == nil {
	 			t.Error("The message response for PostKey was unexpectedly nil")
	 		}

	 		sem <- 1
	 	} (client, i)
 	}
 	for i := 0; i < THRESHOLDOFPAIN; i++ { <-sem }

 	for i := 0; i < THRESHOLDOFPAIN; i++ {
 		go func(client *GateManagementClient, i int) {
	 		t.Logf("GetKey foo%v=bar%v", i, i)
	 		message, err := client.GetKey(fmt.Sprintf("foo%v", i))

	 		if err != nil {
	 			t.Errorf("An unexpected error occurred: %s\n", err)
	 		} else if message == nil {
	 			t.Error("The message response for GetKey was unexpectedly nil")
	 		}

	 		sem <- 1
	 	}(client, i)
 	}
 	for i := 0; i < THRESHOLDOFPAIN; i++ { <-sem }

	for i := 0; i < THRESHOLDOFPAIN; i++ {
 		go func(client *GateManagementClient, i int) {
	 		t.Logf("DeleteKey foo%v=bar%v", i, i)
	 		message, err := client.DeleteKey(fmt.Sprintf("foo%v", i))

	 		if err != nil {
	 			t.Errorf("An unexpected error occurred: %s\n", err)
	 		} else if message == nil {
	 			t.Error("The message response for DeleteKey was unepxectedly nil")
	 		}

	 		sem <- 1
	 	}(client, i)
 	}
 	for i := 0; i < THRESHOLDOFPAIN; i++ { <-sem }
 }