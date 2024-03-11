package detached

import (
	"errors"
	"os"
	"strings"

	"github.com/charmbracelet/charm/client"
	charm "github.com/charmbracelet/charm/proto"
	"github.com/charmbracelet/log"
)

func LinkGen(cfg *client.Config, parentName string, outFilePath string, keysString string) error {
	commaParts := strings.Split(keysString, ",")
	var trimmedCommaParts []string
	for _, key := range commaParts {
		trimmedParts := strings.TrimSpace(key)
		if trimmedParts != "" {
			trimmedCommaParts = append(trimmedCommaParts, trimmedParts)
		}
	}
	var keys []string
	if len(trimmedCommaParts) == 1 {
		spaceParts := strings.Split(trimmedCommaParts[0], " ")
		var trimmedSpaceParts []string
		for _, keyPart := range spaceParts {
			trimmedKeyPart := strings.TrimSpace(keyPart)
			if trimmedKeyPart != "" {
				trimmedSpaceParts = append(trimmedSpaceParts, trimmedKeyPart)
			}
		}
		if len(trimmedSpaceParts)%2 != 0 {
			return errors.New("invalid keys format: expected even number of entries")
		}
		for i := 0; i < len(trimmedSpaceParts); i += 2 {
			keys = append(keys, trimmedSpaceParts[i]+" "+trimmedSpaceParts[i+1])
		}
	} else {
		keys = trimmedCommaParts
	}
	log.Info("parsed keys", "len(keys)", len(keys), "keys", keys, "len(keys)", len(keys))

	cc, err := client.NewClient(cfg)
	if err != nil {
		return err
	}
	lhError := ""
	lh := NewDetachedLinkHandler(
		WithError(func(*charm.Link) {
			lhError = "error"
		}),
		WithRequestDenied(func(*charm.Link) {
			lhError = "request denied"
		}),
		WithTimeout(func(*charm.Link) {
			lhError = "timeout"
		}),
		// WithSameUser(func(*charm.Link) {
		// 	lhError = "same user" // should this be an error or just warn and exit 0?
		// }),
		WithInvalidToken(func(*charm.Link) {
			lhError = "invalid token"
		}),
		WithRequest(func(l *charm.Link) bool {
			valid := false
			for _, k := range keys {
				if k == l.RequestPubKey {
					log.Info("request approved", "reason", "valid, matching key", "link", l, "request-key", l.RequestPubKey)
					valid = true
					break
				}
			}
			if len(keys) == 0 {
				log.Info("request approved", "reason", "no keys provided for matching", "request-key", l.RequestPubKey)
				valid = true
			} else {
				if !valid {
					log.Error("request denied", "reason", "invalid, no matching key", "request-key", l.RequestPubKey, "l", l, "keys", keys)
				}
			}
			return valid
		}),
		WithTokenCreated(func(l *charm.Link) {
			f, err := os.Create(outFilePath)
			if err != nil {
				log.Error("error creating file", "error", err)
			}
			defer f.Close() // nolint:errcheck
			_, err = f.WriteString(string(l.Token))
			if err != nil {
				log.Error("error writing token to file", "error", err)
			}
			log.Info("token written to file", "link", l, "file", outFilePath, "token", l.Token)
		}),
	)

	err = cc.LinkGen(lh)

	if err != nil {
		return err
	}

	if lhError != "" {
		log.Fatal("link failed", "error", lhError)
	}
	return nil
}

type DetachedLinkHandler struct {
	tokenCreated  func(*charm.Link)
	tokenSent     func(*charm.Link)
	validToken    func(*charm.Link)
	invalidToken  func(*charm.Link)
	request       func(*charm.Link) bool
	requestDenied func(*charm.Link)
	sameUser      func(*charm.Link)
	success       func(*charm.Link)
	timeout       func(*charm.Link)
	err           func(*charm.Link)
}

type DetachedLinkHandlerOption func(*DetachedLinkHandler) *DetachedLinkHandler

func WithTokenCreated(f func(*charm.Link)) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.tokenCreated = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithTokenCreated(f func(*charm.Link)) *DetachedLinkHandler {
	lh.tokenCreated = f
	return lh
}

func WithTokenSent(f func(*charm.Link)) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.tokenSent = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithTokenSent(f func(*charm.Link)) *DetachedLinkHandler {
	lh.tokenSent = f
	return lh
}

func WithValidToken(f func(*charm.Link)) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.validToken = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithValidToken(f func(*charm.Link)) *DetachedLinkHandler {
	lh.validToken = f
	return lh
}

func WithInvalidToken(f func(*charm.Link)) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.invalidToken = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithInvalidToken(f func(*charm.Link)) *DetachedLinkHandler {
	lh.invalidToken = f
	return lh
}

func WithRequest(f func(*charm.Link) bool) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.request = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithRequest(f func(*charm.Link) bool) *DetachedLinkHandler {
	lh.request = f
	return lh
}

func WithRequestDenied(f func(*charm.Link)) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.requestDenied = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithRequestDenied(f func(*charm.Link)) *DetachedLinkHandler {
	lh.requestDenied = f
	return lh
}

func WithSameUser(f func(*charm.Link)) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.sameUser = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithSameUser(f func(*charm.Link)) *DetachedLinkHandler {
	lh.sameUser = f
	return lh
}

func WithSuccess(f func(*charm.Link)) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.success = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithSuccess(f func(*charm.Link)) *DetachedLinkHandler {
	lh.success = f
	return lh
}

func WithTimeout(f func(*charm.Link)) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.timeout = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithTimeout(f func(*charm.Link)) *DetachedLinkHandler {
	lh.timeout = f
	return lh
}

func WithError(f func(*charm.Link)) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.err = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithError(f func(*charm.Link)) *DetachedLinkHandler {
	lh.err = f
	return lh
}

func NewDetachedLinkHandler(opts ...DetachedLinkHandlerOption) *DetachedLinkHandler {
	lh := &DetachedLinkHandler{
		tokenCreated:  func(*charm.Link) {},
		tokenSent:     func(*charm.Link) {},
		validToken:    func(*charm.Link) {},
		invalidToken:  func(*charm.Link) {},
		request:       func(*charm.Link) bool { return false },
		requestDenied: func(*charm.Link) {},
		sameUser:      func(*charm.Link) {},
		success:       func(*charm.Link) {},
		timeout:       func(*charm.Link) {},
		err:           func(*charm.Link) {},
	}

	for _, opt := range opts {
		opt(lh)
	}

	return lh
}

func (lh *DetachedLinkHandler) TokenCreated(l *charm.Link) {
	log.Info("token created", "link", l, "token", l.Token)
	lh.tokenCreated(l)
}

func (lh *DetachedLinkHandler) TokenSent(l *charm.Link) {
	log.Info("token sent", "link", l)
	lh.tokenSent(l)
}

func (lh *DetachedLinkHandler) ValidToken(l *charm.Link) {
	log.Info("valid token", "link", l)
	lh.validToken(l)
}

func (lh *DetachedLinkHandler) InvalidToken(l *charm.Link) {
	log.Error("invalid token", "link", l)
	lh.invalidToken(l)
}

func (lh *DetachedLinkHandler) Request(l *charm.Link) bool {
	log.Info("request", "link", l)
	return lh.request(l)
}

func (lh *DetachedLinkHandler) RequestDenied(l *charm.Link) {
	log.Error("request denied", "link", l)
	lh.requestDenied(l)
}

func (lh *DetachedLinkHandler) SameUser(l *charm.Link) {
	log.Warn("same user", "link", l)
	lh.sameUser(l)
}

func (lh *DetachedLinkHandler) Success(l *charm.Link) {
	log.Info("success", "link", l)
	lh.success(l)
}

func (lh *DetachedLinkHandler) Timeout(l *charm.Link) {
	log.Error("timeout", "link", l)
	lh.timeout(l)
}

func (lh *DetachedLinkHandler) Error(l *charm.Link) {
	log.Error("error", "link", l)
	lh.err(l)
}
