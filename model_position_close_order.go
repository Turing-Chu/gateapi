/*
 * Gate API v4
 *
 * APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * API version: 4.6.1
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Current close order if any, or `null`
type PositionCloseOrder struct {
	// Close order ID
	Id int64 `json:"id,omitempty"`
	// Close order price
	Price string `json:"price,omitempty"`
	// Is the close order from liquidation
	IsLiq bool `json:"is_liq,omitempty"`
}
