/***********************************************************************************
 * a file used to group together all models used to unmarshal the HTTP body of
 * responses from the etc server into models
 **********************************************************************************/
package gatemanagementclient


/***********************************************************************************
 * a simple type used to unmarshal JSON responses back from the etcd server giving
 * context to how the etcd server encountered a problem processing a request
 **********************************************************************************/
type EtcdErrorMessage struct {
	ErrorCode int32
	Message string
	Cause string
	Index int32
}


/***********************************************************************************
 * a simple type used to unmarshal JSON responses back from the etcd server giving
 * context to how the etcd server handled an action processing request
 **********************************************************************************/
type EtcdActionMessage struct {
	Action string
	Key string
	Value string
	NewKey bool
	PrevValue string
	Index int32
}