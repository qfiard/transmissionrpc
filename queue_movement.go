package transmissionrpc

import (
	"context"
	"fmt"
)

/*
	Queue Movement Requests
	https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L601
*/

// QueueMoveTop moves IDs to the top of the queue list.
func (c *Client) QueueMoveTop(ctx context.Context, IDs []int64) (err error) {
	payload := &queueMovePayload{IDs: IDs}
	if err = c.rpcCall(ctx, "queue-move-top", payload, nil); err != nil {
		err = fmt.Errorf("'queue-move-top' rpc method failed: %v", err)
	}
	return
}

// QueueMoveUp moves IDs of one position up on the queue list.
func (c *Client) QueueMoveUp(ctx context.Context, IDs []int64) (err error) {
	payload := &queueMovePayload{IDs: IDs}
	if err = c.rpcCall(ctx, "queue-move-up", payload, nil); err != nil {
		err = fmt.Errorf("'queue-move-up' rpc method failed: %v", err)
	}
	return
}

// QueueMoveDown moves IDs of one position down on the queue list.
func (c *Client) QueueMoveDown(ctx context.Context, IDs []int64) (err error) {
	payload := &queueMovePayload{IDs: IDs}
	if err = c.rpcCall(ctx, "queue-move-down", payload, nil); err != nil {
		err = fmt.Errorf("'queue-move-down' rpc method failed: %v", err)
	}
	return
}

// QueueMoveBottom moves IDs to the bottom of the queue list.
func (c *Client) QueueMoveBottom(ctx context.Context, IDs []int64) (err error) {
	payload := &queueMovePayload{IDs: IDs}
	if err = c.rpcCall(ctx, "queue-move-bottom", payload, nil); err != nil {
		err = fmt.Errorf("'queue-move-bottom' rpc method failed: %v", err)
	}
	return
}

type queueMovePayload struct {
	IDs []int64 `json:"ids"`
}
