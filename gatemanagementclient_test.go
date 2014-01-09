/***********************************************************************************
 * this is a simple set of tests used to exercise the functionality found in
 * gatemanagementclient.go
 **********************************************************************************/
package gatemanagementclient


import "testing"


var client = GateManagementClient{"http://Johns-Macbook-Air.local:4001/v1/keys"}


/**
 * tests that retrieving a key that does not exist returns the expected 
 * errorCode: 100 and message: "Key not found"
 **/
func TestGetKeyInvalidKey(t *testing.T) {
	message, error := client.GetKey("whoopsie")
	if error != nil {
		t.Errorf("An unexpected error occurred: %s\n", error)
	} else if message == nil {
		t.Errorf("message was unexpectedly nil")	
	} else {
		etcdErrorMessage, ok := message.(EtcdErrorMessage)
		if !ok {
			t.Errorf("An unexpected type was returned from the call to GetKey")
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
	message, error := client.PostKey("foo", "bar")
	if error != nil {
		t.Errorf("An unexpected error occurred: %s\n", error)
	} else if message == nil {
		t.Errorf("The configAction response for PostKey was unexpectedly nil")
	} else {
		etcdActionMessage, ok := message.(EtcdActionMessage)
		if !ok {
			t.Errorf("An unexpected type was returned from the call to PostKey")
		} else if etcdActionMessage.Action != "set" || etcdActionMessage.Key != "/foo" ||
				  etcdActionMessage.Value != "bar" {
		  	t.Errorf("PostKey EtcdActionMessage did not return the expected values.  " +
		  				"EtcdActionMessage was %+v\n", etcdActionMessage)
		} else {
			message, error := client.DeleteKey("foo")
			if error != nil {
				t.Errorf("An unexpected error occurred: %s\n", error)
			} else if message == nil {
				t.Errorf("The configAction response for DeleteKey was unexpectedly nil")
			} else {
				etcdActionMessage, ok := message.(EtcdActionMessage)
				if !ok {
					t.Errorf("An unexpected type was returned from the call to DeleteKey")
				} else if etcdActionMessage.Action != "delete" || etcdActionMessage.Key != "/foo" ||
						  etcdActionMessage.PrevValue != "bar" {
				  	t.Errorf("DeleteKey EtcdActionMessage did not return the expected values.  " +
				  				"EtcdActionMessage was %+v\n", etcdActionMessage)
			  	}
			}
		}
	}
}