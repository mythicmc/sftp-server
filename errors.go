package main

type fxerr uint32

const (
	// ErrSSHQuotaExceeded ...
	// Extends the default SFTP server to return a quota exceeded error to the client.
	//
	// @see https://tools.ietf.org/id/draft-ietf-secsh-filexfer-13.txt
	ErrSSHQuotaExceeded = fxerr(15)
)

func (e fxerr) Error() string {
	switch e {
	case ErrSSHQuotaExceeded:
		return "Quota Exceeded"
	default:
		return "Failure"
	}
}
