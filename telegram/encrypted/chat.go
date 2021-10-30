package encrypted

import (
	"encoding/binary"
	"math/big"

	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/telegram/internal/dh"
	"github.com/gotd/td/tg"
)

// getKeyFingerprint computes key fingerprint.
func getKeyFingerprint(key crypto.AuthKey) int64 {
	return int64(binary.LittleEndian.Uint64(key.ID[:]))
}

type (
	// ExchangeState contains key exchange state.
	//
	// See https://core.telegram.org/api/end-to-end/pfs.
	ExchangeState struct {
		// G is a g DH parameter.
		G int
		// P is a p DH parameter.
		P *big.Int
		// Key is message encryption key.
		Key crypto.AuthKey
		// Originator denotes current user is creator.
		Originator bool
	}

	// Chat is an encrypted chat metadata structure.
	Chat struct {
		// Chat ID.
		ID int
		// AccessHash is a chat access hash.
		AccessHash int64
		// Layer is a TL encrypted schema layer version.
		Layer int
		// Date chat was created.
		Date int
		// Chat creator ID
		AdminID int64
		// ID of the second chat participant
		ParticipantID int64

		// InSeq is an incoming message sequence.
		// InSeq stored as total number of all consumed messages.
		InSeq int
		// OutSeq is an outgoing message sequence.
		// InSeq stored as total number of all sent messages.
		OutSeq int
		// HisInSeq is an incoming message sequence of other party.
		// Need for security checks.
		HisInSeq int

		ExchangeState
	}
)

func (c *Chat) init(
	obj *tg.EncryptedChat,
	originator bool,
	key crypto.AuthKey,
	dhCfg dh.Config,
) {
	c.ID = obj.ID
	c.AccessHash = obj.AccessHash
	c.Layer = minLayer
	c.Date = obj.Date
	c.AdminID = obj.AdminID
	c.ParticipantID = obj.ParticipantID
	c.InSeq = 0
	c.OutSeq = 0
	c.HisInSeq = 0
	c.ExchangeState = ExchangeState{
		G:          dhCfg.G,
		P:          dhCfg.P,
		Key:        key,
		Originator: originator,
	}
}

// seqNo returns a pair of incoming and outgoing messages sequence numbers.
func (c Chat) seqNo() (in, out int) {
	if c.Originator {
		in = 2 * c.InSeq
		out = 2*c.OutSeq + 1
	} else {
		in = 2*c.InSeq + 1
		out = 2 * c.OutSeq
	}
	return in, out
}

func (c *Chat) nextMessage() (in, out int) {
	in, out = c.seqNo()
	c.OutSeq++
	return in, out
}

type consumeResult int

const (
	skipMessage    consumeResult = -1
	consumeMessage consumeResult = 0
	fillGap        consumeResult = 1
	abortChat      consumeResult = 2
)

func (c *Chat) consumeMessage(hisInSeq, hisOutSeq int) consumeResult {
	// See https://core.telegram.org/api/end-to-end/seq_no#checking-out-seq-no.
	//
	// If the received out_seq_no<=C, the local client must drop the message (repeated message).
	// The client should not check the contents of the message because the original message could have
	// been deleted (see Deleting unacknowledged messages).
	//
	// We store C+1, so check < instead of <=.
	myInSeq, myOutSeq := c.seqNo()
	if hisOutSeq < myInSeq {
		return skipMessage
	}

	// If the received out_seq_no>C+1, it most likely means that the server left out some messages due
	// to a technical failure or due to the messages becoming obsolete. A temporary solution to this is
	// to simply abort the secret chat. But since this may cause some existing older secret chats to be aborted,
	// it is strongly recommended for the client to properly handle such seq_no gaps. Note that in_seq_no is not
	// increased upon receipt of such a message; it is advanced only after all preceding gaps are filled.
	if hisOutSeq > myInSeq {
		return fillGap
	}
	c.InSeq++

	// See https://core.telegram.org/api/end-to-end/seq_no#checking-and-handling-in-seq-no.
	//
	// - in_seq_no must form a non-decreasing sequence of non-negative integer numbers.
	if hisInSeq < 0 || (c.HisInSeq > 0 && hisInSeq < c.HisInSeq) {
		// If in_seq_no contradicts these criteria, the local client is required
		// to immediately abort the secret chat.
		return abortChat
	}
	c.HisInSeq = hisInSeq
	// - in_seq_no must be valid at the moment of receiving the message, that is, if D
	// is the out_seq_no of last message we sent, the received in_seq_no should not
	// be greater than D + 1.
	if hisInSeq > myOutSeq {
		// If in_seq_no contradicts these criteria, the local client is required
		// to immediately abort the secret chat.
		return abortChat
	}

	return consumeMessage
}
