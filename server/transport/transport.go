/*
 * Copyright (c) 2018 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package transport

import (
	"crypto/tls"
	"errors"
	"io"

	"github.com/ortuman/jackal/config"
	"github.com/ortuman/jackal/xml"
)

var ErrTooLargeStanza = errors.New("too large stanza")

// Transport represents a stream transport mechanism.
type Transport interface {
	io.Closer

	// ReadElement reads next available XML element.
	ReadElement() (xml.XElement, error)

	// WriteString writes a raw string to the transport.
	WriteString(string) error

	// WriteElement writes an element to the transport
	// serializing it to it's XML representation.
	WriteElement(elem xml.XElement, includeClosing bool) error

	// StartTLS secures the transport using SSL/TLS
	StartTLS(*tls.Config)

	// EnableCompression activates a compression
	// mechanism on the transport.
	EnableCompression(config.CompressionLevel)

	// ChannelBindingBytes returns current transport
	// channel binding bytes.
	ChannelBindingBytes(config.ChannelBindingMechanism) []byte
}
