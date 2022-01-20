package v1

import "context"

type Operator uint8

const (
	_INCLUDES = iota
	_EXCLUDES
	_EQUALS
	_EQUALS_NOT
	_CONTAINS
	_CONTAINS_NOT
	_GREATER
	_LOWER
	_EMPTY
	_NOT_EMPTY
	_EXISTS
	_EXISTS_NOT
)

var (
	OP_INCLUDES     Operator = _INCLUDES
	OP_EXCLUDES     Operator = _EXCLUDES
	OP_EQUALS       Operator = _EQUALS
	OP_EQUALS_NOT   Operator = _EQUALS_NOT
	OP_CONTAINS     Operator = _CONTAINS
	OP_CONTAINS_NOT Operator = _CONTAINS_NOT
	OP_GREATER      Operator = _GREATER
	OP_LOWER        Operator = _LOWER
	OP_EMPTY        Operator = _EMPTY
	OP_NOT_EMPTY    Operator = _NOT_EMPTY
	OP_EXISTS       Operator = _EXISTS
	OP_EXISTS_NOT   Operator = _EXISTS_NOT
)

func (op *Operator) Invert() Operator {
	if op == &OP_INCLUDES {
		return OP_EXCLUDES
	}
	if op == &OP_EXCLUDES {
		return OP_INCLUDES
	}
	if op == &OP_EQUALS {
		return OP_EQUALS_NOT
	}
	if op == &OP_EQUALS_NOT {
		return OP_EQUALS
	}
	if op == &OP_CONTAINS {
		return OP_CONTAINS_NOT
	}
	if op == &OP_CONTAINS_NOT {
		return OP_CONTAINS
	}
	if op == &OP_GREATER {
		return OP_LOWER
	}
	if op == &OP_LOWER {
		return OP_GREATER
	}
	if op == &OP_EMPTY {
		return OP_NOT_EMPTY
	}
	if op == &OP_NOT_EMPTY {
		return OP_EMPTY
	}
	if op == &OP_EXISTS {
		return OP_EXISTS_NOT
	}
	if op == &OP_EXISTS_NOT {
		return OP_EXISTS
	}
	return OP_INCLUDES
}

// Event defines events the sockets can use
type Event struct {
	UserID    string                 `dynamo:"userId" json:"userId"`
	ProjectID string                 `dynamo:"projectId" json:"projectId"`
	Name      string                 `dynamo:"name" json:"name"`
	Params    map[string]interface{} `dynamo:"params" json:"params,omitempty"`
	CreatedAt string                 `dynamo:"createdAt" json:"createdAt"`
	UpdatedAt string                 `dynamo:"updatedAt" json:"updatedAt"`
}

type DispatchEventParams struct {
	Data SocketData `json:"data"`
}

// Data sent to the socket
type SocketData struct {
	Event              string                 `json:"event"`
	Data               map[string]interface{} `json:"data"`
	MetaSelector       []MetaSelector         `json:"metaSelector"`
	SourceConnectionID string                 `json:"connectionId"`
}

type MetaSelector struct {
	Path     string      `json:"path"`
	Operator Operator    `json:"operator"`
	Value    interface{} `json:"value"`
}

// Data sent from the socket
type SocketAnswerData struct {
	Action       string     `json:"action"`
	Data         SocketData `json:"data"`
	ConnectionID string     `json:"connectionId"`
}

type EventClient struct {
	client *Client
}

func (ec *EventClient) Dispatch(params *DispatchEventParams) error {
	req, err := ec.client.NewRequest(context.Background(), "POST", "dispatch", structToReader(params))
	if err != nil {
		return err
	}

	_, err = ec.client.Do(req, &[]byte{})
	if err != nil {
		return err
	}
	return nil
}
