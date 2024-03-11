package detached

import (
	"github.com/charmbracelet/charm/client"
	charm "github.com/charmbracelet/charm/proto"
	"github.com/charmbracelet/log"
)

func Link(cfg *client.Config, parentName string, linkCode string) error {
	log.Info("attaching link", "code", linkCode)
	cc, err := client.NewClient(cfg)
	if err != nil {
		return err
	}

	lhError := ""
	lh := NewDetachedLinkHandler(
		WithError(func(l *charm.Link) {
			lhError = "error"
		}),
		WithRequestDenied(func(l *charm.Link) {
			lhError = "request denied"
		}),
		WithTimeout(func(l *charm.Link) {
			lhError = "timeout"
		}),
		// WithSameUser(func(l *charm.Link) {
		// 	lhError = "same user" // should this be an error or just warn and exit 0?
		// }),
		WithInvalidToken(func(l *charm.Link) {
			lhError = "invalid token"
		}),
	)

	err = cc.Link(lh, linkCode)

	if err != nil {
		return err
	}

	if lhError != "" {
		log.Fatal("link failed", "error", lhError)
	}
	return nil
}

type DetachedLinkHandler struct {
	tokenCreated  func(l *charm.Link)
	tokenSent     func(l *charm.Link)
	validToken    func(l *charm.Link)
	invalidToken  func(l *charm.Link)
	request       func(l *charm.Link)
	requestDenied func(l *charm.Link)
	sameUser      func(l *charm.Link)
	success       func(l *charm.Link)
	timeout       func(l *charm.Link)
	err           func(l *charm.Link)
}

type DetachedLinkHandlerOption func(*DetachedLinkHandler) *DetachedLinkHandler

func WithTokenCreated(f func(l *charm.Link)) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.tokenCreated = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithTokenCreated(f func(l *charm.Link)) *DetachedLinkHandler {
	lh.tokenCreated = f
	return lh
}

func WithTokenSent(f func(l *charm.Link)) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.tokenSent = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithTokenSent(f func(l *charm.Link)) *DetachedLinkHandler {
	lh.tokenSent = f
	return lh
}

func WithValidToken(f func(l *charm.Link)) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.validToken = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithValidToken(f func(l *charm.Link)) *DetachedLinkHandler {
	lh.validToken = f
	return lh
}

func WithInvalidToken(f func(l *charm.Link)) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.invalidToken = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithInvalidToken(f func(l *charm.Link)) *DetachedLinkHandler {
	lh.invalidToken = f
	return lh
}

func WithRequest(f func(l *charm.Link)) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.request = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithRequest(f func(l *charm.Link)) *DetachedLinkHandler {
	lh.request = f
	return lh
}

func WithRequestDenied(f func(l *charm.Link)) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.requestDenied = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithRequestDenied(f func(l *charm.Link)) *DetachedLinkHandler {
	lh.requestDenied = f
	return lh
}

func WithSameUser(f func(l *charm.Link)) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.sameUser = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithSameUser(f func(l *charm.Link)) *DetachedLinkHandler {
	lh.sameUser = f
	return lh
}

func WithSuccess(f func(l *charm.Link)) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.success = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithSuccess(f func(l *charm.Link)) *DetachedLinkHandler {
	lh.success = f
	return lh
}

func WithTimeout(f func(l *charm.Link)) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.timeout = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithTimeout(f func(l *charm.Link)) *DetachedLinkHandler {
	lh.timeout = f
	return lh
}

func WithError(f func(l *charm.Link)) DetachedLinkHandlerOption {
	return func(lh *DetachedLinkHandler) *DetachedLinkHandler {
		lh.err = f
		return lh
	}
}

func (lh *DetachedLinkHandler) WithError(f func(l *charm.Link)) *DetachedLinkHandler {
	lh.err = f
	return lh
}

func NewDetachedLinkHandler(opts ...DetachedLinkHandlerOption) *DetachedLinkHandler {
	lh := &DetachedLinkHandler{
		tokenCreated:  func(l *charm.Link) {},
		tokenSent:     func(l *charm.Link) {},
		validToken:    func(l *charm.Link) {},
		invalidToken:  func(l *charm.Link) {},
		request:       func(l *charm.Link) {},
		requestDenied: func(l *charm.Link) {},
		sameUser:      func(l *charm.Link) {},
		success:       func(l *charm.Link) {},
		timeout:       func(l *charm.Link) {},
		err:           func(l *charm.Link) {},
	}

	for _, opt := range opts {
		opt(lh)
	}

	return lh
}

func (lh *DetachedLinkHandler) TokenCreated(l *charm.Link) {
	// Not implemented for the link participant
	log.Warn("token created", "link", l)
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
	// Not implemented for the link participant
	log.Warn("request", "link", l)
	lh.request(l)
	return false
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
