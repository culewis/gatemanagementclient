/***********************************************************************************
 * a file used to group together all models used to unmarshal the HTTP body of
 * responses from the etc server into models
 **********************************************************************************/
package gatemanagementclient

/***********************************************************************************
 * a simple type used to unmarshal JSON responses back from the etcd server giving
 * context to how the etcd server responded to the specific request
 **********************************************************************************/
type EtcdMessage struct {
	ErrorCode int32
	Message string
	Cause string
	Index int32
}

// TODO: EtcdConfigAction to encapsulate responses from valid actions that took place
// like creating keys, retrieving keys, deleting keys, updating keys, etc(d) :)