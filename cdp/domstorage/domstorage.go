// Package domstorage provides the Chrome Debugging Protocol
// commands, types, and events for the Chrome DOMStorage domain.
//
// Query and modify DOM storage.
//
// Generated by the chromedp-gen command.
package domstorage

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	cdp "github.com/knq/chromedp/cdp"
	"github.com/mailru/easyjson"
)

// EnableParams enables storage tracking, storage events will now be
// delivered to the client.
type EnableParams struct{}

// Enable enables storage tracking, storage events will now be delivered to
// the client.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes DOMStorage.enable.
func (p *EnableParams) Do(ctxt context.Context, h cdp.FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandDOMStorageEnable, cdp.Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return cdp.ErrContextDone
	}

	return cdp.ErrUnknownResult
}

// DisableParams disables storage tracking, prevents storage events from
// being sent to the client.
type DisableParams struct{}

// Disable disables storage tracking, prevents storage events from being sent
// to the client.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes DOMStorage.disable.
func (p *DisableParams) Do(ctxt context.Context, h cdp.FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandDOMStorageDisable, cdp.Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return cdp.ErrContextDone
	}

	return cdp.ErrUnknownResult
}

// ClearParams [no description].
type ClearParams struct {
	StorageID *StorageID `json:"storageId"`
}

// Clear [no description].
//
// parameters:
//   storageID
func Clear(storageID *StorageID) *ClearParams {
	return &ClearParams{
		StorageID: storageID,
	}
}

// Do executes DOMStorage.clear.
func (p *ClearParams) Do(ctxt context.Context, h cdp.FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandDOMStorageClear, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return cdp.ErrContextDone
	}

	return cdp.ErrUnknownResult
}

// GetDOMStorageItemsParams [no description].
type GetDOMStorageItemsParams struct {
	StorageID *StorageID `json:"storageId"`
}

// GetDOMStorageItems [no description].
//
// parameters:
//   storageID
func GetDOMStorageItems(storageID *StorageID) *GetDOMStorageItemsParams {
	return &GetDOMStorageItemsParams{
		StorageID: storageID,
	}
}

// GetDOMStorageItemsReturns return values.
type GetDOMStorageItemsReturns struct {
	Entries []Item `json:"entries,omitempty"`
}

// Do executes DOMStorage.getDOMStorageItems.
//
// returns:
//   entries
func (p *GetDOMStorageItemsParams) Do(ctxt context.Context, h cdp.FrameHandler) (entries []Item, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return nil, err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandDOMStorageGetDOMStorageItems, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return nil, cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r GetDOMStorageItemsReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return nil, cdp.ErrInvalidResult
			}

			return r.Entries, nil

		case error:
			return nil, v
		}

	case <-ctxt.Done():
		return nil, cdp.ErrContextDone
	}

	return nil, cdp.ErrUnknownResult
}

// SetDOMStorageItemParams [no description].
type SetDOMStorageItemParams struct {
	StorageID *StorageID `json:"storageId"`
	Key       string     `json:"key"`
	Value     string     `json:"value"`
}

// SetDOMStorageItem [no description].
//
// parameters:
//   storageID
//   key
//   value
func SetDOMStorageItem(storageID *StorageID, key string, value string) *SetDOMStorageItemParams {
	return &SetDOMStorageItemParams{
		StorageID: storageID,
		Key:       key,
		Value:     value,
	}
}

// Do executes DOMStorage.setDOMStorageItem.
func (p *SetDOMStorageItemParams) Do(ctxt context.Context, h cdp.FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandDOMStorageSetDOMStorageItem, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return cdp.ErrContextDone
	}

	return cdp.ErrUnknownResult
}

// RemoveDOMStorageItemParams [no description].
type RemoveDOMStorageItemParams struct {
	StorageID *StorageID `json:"storageId"`
	Key       string     `json:"key"`
}

// RemoveDOMStorageItem [no description].
//
// parameters:
//   storageID
//   key
func RemoveDOMStorageItem(storageID *StorageID, key string) *RemoveDOMStorageItemParams {
	return &RemoveDOMStorageItemParams{
		StorageID: storageID,
		Key:       key,
	}
}

// Do executes DOMStorage.removeDOMStorageItem.
func (p *RemoveDOMStorageItemParams) Do(ctxt context.Context, h cdp.FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandDOMStorageRemoveDOMStorageItem, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return cdp.ErrContextDone
	}

	return cdp.ErrUnknownResult
}