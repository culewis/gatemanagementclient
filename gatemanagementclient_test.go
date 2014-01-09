/***********************************************************************************
 * this is a simple set of tests used to exercise the functionality found in
 * gatemanagementclient.go
 **********************************************************************************/
package gatemanagementclient


import "testing"


/**
 * tests that retrieving a key that does not exist returns the expected 
 * errorCode: 100 and message: "Key not found"
 **/
func TestGetKeyInvalidKey(t *testing.T) {
	message, error := GetKey("whoopsie")
	if error != "" {
		t.Errorf("An unexpected error occurred: %s\n", error)
	} else if message.ErrorCode != 100 || message.Message != "Key not found" {
		t.Errorf("Message ErrorCode was not of the expected type: %+v\n", message);
	}
}

/**
 * tests that creating and retrieving a key creates the key/value combination
 * properly without encountering any errors
 **/
func TestGetKey(t *testing.T) {
	configAction, error := PostKey("foo", "bar")
	if error != "" {
		t.Errorf("An unexpected error occurred: %s\n", error)
	} else if configAction == nil {
		t.Errorf("The configAction response was unexpectedly null")
	}
}