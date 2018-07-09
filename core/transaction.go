// Copyright (C) 2018  MediBloc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package core

import (
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/medibloc/go-medibloc/common"
	"github.com/medibloc/go-medibloc/core/pb"
	"github.com/medibloc/go-medibloc/crypto"
	"github.com/medibloc/go-medibloc/crypto/signature"
	"github.com/medibloc/go-medibloc/crypto/signature/algorithm"
	"github.com/medibloc/go-medibloc/util"
	"github.com/medibloc/go-medibloc/util/byteutils"
	"golang.org/x/crypto/sha3"
	"github.com/medibloc/go-medibloc/consensus/dpos/pb"
	"encoding/json"
)

// Transaction struct represents transaction
type Transaction struct {
	hash      []byte
	from      common.Address
	to        common.Address
	value     *util.Uint128
	timestamp int64
	data      *corepb.Data
	nonce     uint64
	chainID   uint32
	alg       algorithm.Algorithm
	sign      []byte
	payerSign []byte
}

// Transactions is just multiple txs
type Transactions []*Transaction

// ToProto converts Transaction to corepb.Transaction
func (tx *Transaction) ToProto() (proto.Message, error) {
	value, err := tx.value.ToFixedSizeByteSlice()
	if err != nil {
		return nil, err
	}

	return &corepb.Transaction{
		Hash:      tx.hash,
		From:      tx.from.Bytes(),
		To:        tx.to.Bytes(),
		Value:     value,
		Timestamp: tx.timestamp,
		Data:      tx.data,
		Nonce:     tx.nonce,
		ChainId:   tx.chainID,
		Alg:       uint32(tx.alg),
		Sign:      tx.sign,
		PayerSign: tx.payerSign,
	}, nil
}

// FromProto converts corepb.Transaction to Transaction
func (tx *Transaction) FromProto(msg proto.Message) error {
	if msg, ok := msg.(*corepb.Transaction); ok {
		tx.hash = msg.Hash
		tx.from = common.BytesToAddress(msg.From)
		tx.to = common.BytesToAddress(msg.To)

		value, err := util.NewUint128FromFixedSizeByteSlice(msg.Value)
		if err != nil {
			return err
		}
		tx.value = value
		tx.timestamp = msg.Timestamp
		tx.data = msg.Data
		tx.nonce = msg.Nonce
		tx.chainID = msg.ChainId
		alg := algorithm.Algorithm(msg.Alg)
		err = crypto.CheckAlgorithm(alg)
		if err != nil {
			return err
		}
		tx.alg = alg
		tx.sign = msg.Sign
		tx.payerSign = msg.PayerSign

		return nil
	}

	return ErrCannotConvertTransaction
}

// BuildTransaction generates a Transaction instance with entire fields
func BuildTransaction(
	chainID uint32,
	from, to common.Address,
	value *util.Uint128,
	nonce uint64,
	timestamp int64,
	data *corepb.Data,
	hash []byte,
	alg uint32,
	sign []byte,
	payerSign []byte) (*Transaction, error) {
	return &Transaction{
		from:      from,
		to:        to,
		value:     value,
		timestamp: timestamp,
		data:      data,
		nonce:     nonce,
		chainID:   chainID,
		hash:      hash,
		alg:       algorithm.Algorithm(alg),
		sign:      sign,
		payerSign: payerSign,
	}, nil
}

// NewTransaction generates a Transaction instance
func NewTransaction(
	chainID uint32,
	from, to common.Address,
	value *util.Uint128,
	nonce uint64,
	payloadType string,
	payload []byte) (*Transaction, error) {
	return NewTransactionWithSign(chainID, from, to, value, nonce, payloadType, payload, []byte{}, []byte{}, []byte{})
}

// NewTransactionWithSign generates a Transaction instance with sign
func NewTransactionWithSign(
	chainID uint32,
	from, to common.Address,
	value *util.Uint128,
	nonce uint64,
	payloadType string,
	payload []byte,
	hash []byte,
	sign []byte,
	payerSign []byte) (*Transaction, error) {

	return &Transaction{
		from:      from,
		to:        to,
		value:     value,
		timestamp: time.Now().Unix(),
		data:      &corepb.Data{Type: payloadType, Payload: payload},
		nonce:     nonce,
		chainID:   chainID,
		hash:      hash,
		sign:      sign,
		payerSign: payerSign,
	}, nil
}

// CalcHash calculates transaction's hash.
func (tx *Transaction) CalcHash() ([]byte, error) {
	hasher := sha3.New256()

	value, err := tx.value.ToFixedSizeByteSlice()
	if err != nil {
		return nil, err
	}
	txHashTarget := &corepb.TransactionHashTarget{
		From:      tx.from.Bytes(),
		To:        tx.to.Bytes(),
		Value:     value,
		Timestamp: tx.timestamp,
		Data: &corepb.Data{
			Type:    tx.data.Type,
			Payload: tx.data.Payload,
		},
		Nonce:   tx.nonce,
		ChainId: tx.chainID,
	}
	data, err := proto.Marshal(txHashTarget)
	if err != nil {
		return nil, err
	}
	hasher.Write(data)

	hash := hasher.Sum(nil)
	return hash, nil
}

// SignThis signs tx with given signature interface
func (tx *Transaction) SignThis(signer signature.Signature) error {
	tx.alg = signer.Algorithm()
	hash, err := tx.CalcHash()
	if err != nil {
		return err
	}

	sig, err := signer.Sign(hash)
	if err != nil {
		return err
	}
	tx.hash = hash
	tx.sign = sig
	return nil
}

func (tx *Transaction) getPayerSignTarget() []byte {
	hasher := sha3.New256()

	hasher.Write(tx.hash)
	hasher.Write(tx.sign)

	hash := hasher.Sum(nil)
	return hash
}

func (tx *Transaction) recoverPayer() (common.Address, error) {
	if tx.payerSign == nil || len(tx.payerSign) == 0 {
		return common.Address{}, ErrPayerSignatureNotExist
	}
	msg := tx.getPayerSignTarget()

	sig, err := crypto.NewSignature(tx.alg)
	if err != nil {
		return common.Address{}, err
	}

	pubKey, err := sig.RecoverPublic(msg, tx.payerSign)
	if err != nil {
		return common.Address{}, err
	}

	return common.PublicKeyToAddress(pubKey)
}

// SignByPayer puts payer's sign in tx
func (tx *Transaction) SignByPayer(signer signature.Signature) error {
	target := tx.getPayerSignTarget()

	sig, err := signer.Sign(target)
	if err != nil {
		return err
	}
	tx.payerSign = sig
	return nil
}

// VerifyIntegrity returns transaction verify result, including Hash and Signature.
func (tx *Transaction) VerifyIntegrity(chainID uint32) error {
	// check ChainID.
	if tx.chainID != chainID {
		return ErrInvalidChainID
	}

	// check Hash.
	wantedHash, err := tx.CalcHash()
	if err != nil {
		return err
	}
	if !byteutils.Equal(wantedHash, tx.hash) {
		return ErrInvalidTransactionHash
	}

	// check Signature.
	return tx.verifySign()
}

func (tx *Transaction) verifySign() error {
	signer, err := tx.recoverSigner()
	if err != nil {
		return err
	}
	if !tx.from.Equals(signer) {
		return ErrInvalidTransactionSigner
	}
	return nil
}

func (tx *Transaction) recoverSigner() (common.Address, error) {
	sig, err := crypto.NewSignature(tx.alg)
	if err != nil {
		return common.Address{}, err
	}

	pubKey, err := sig.RecoverPublic(tx.hash, tx.sign)
	if err != nil {
		return common.Address{}, err
	}

	return common.PublicKeyToAddress(pubKey)
}

// From returns from
func (tx *Transaction) From() common.Address {
	return tx.from
}

// To returns to
func (tx *Transaction) To() common.Address {
	return tx.to
}

// Value returns value
func (tx *Transaction) Value() *util.Uint128 {
	return tx.value
}

// Timestamp returns tx timestamp
func (tx *Transaction) Timestamp() int64 {
	return tx.timestamp
}

// SetTimestamp sets tx timestamp
func (tx *Transaction) SetTimestamp(t int64) {
	tx.timestamp = t
}

// Type returns tx type
func (tx *Transaction) Type() string {
	return tx.data.Type
}

// Data returns tx payload
func (tx *Transaction) Data() []byte {
	return tx.data.Payload
}

// Nonce returns tx nonce
func (tx *Transaction) Nonce() uint64 {
	return tx.nonce
}

// Hash returns hash
func (tx *Transaction) Hash() []byte {
	return tx.hash
}

// Signature returns tx signature
func (tx *Transaction) Signature() []byte {
	return tx.sign
}

// PayerSignature returns tx payer signature
func (tx *Transaction) PayerSignature() []byte {
	return tx.payerSign
}

// String returns string representation of tx
func (tx *Transaction) String() string {
	return fmt.Sprintf(`{chainID:%v, hash:%v, from:%v, to:%v, value:%v, type:%v, alg:%v, nonce:%v'}`,
		tx.chainID,
		tx.hash,
		tx.from,
		tx.to,
		tx.value.String(),
		tx.Type(),
		tx.alg,
		tx.nonce,
	)
}

//SendTx is a structure for sending MED
type SendTx struct {
	*Transaction
}

//NewSendTx returns SendTx
func NewSendTx(tx *Transaction) (ExecutableTx, error) {
	if tx.value.Cmp(util.Uint128Zero()) == 0 {
		return nil, ErrVoidTransaction
	}
	return &SendTx{tx}, nil
}

//Execute SendTx
func (tx *SendTx) Execute(b *Block) error {
	as := b.state.AccState()

	if err := as.SubBalance(tx.from.Bytes(), tx.value); err != nil {
		return err
	}
	return as.AddBalance(tx.to.Bytes(), tx.value)
}

//AddRecordTx is a structure for adding record
type AddRecordTx struct {
	owner     common.Address
	timestamp int64
	payload   []byte
}

//NewAddRecordTx returns AddRecordTx
func NewAddRecordTx(tx *Transaction) (ExecutableTx, error) {
	return &AddRecordTx{
		owner:     tx.From(),
		timestamp: tx.Timestamp(),
		payload:   tx.Data(),
	}, nil
}

//Execute AddRecordTx
func (tx *AddRecordTx) Execute(b *Block) error {
	rs := b.state.recordsState

	payload, err := BytesToAddRecordPayload(tx.payload)
	if err != nil {
		return err
	}

	pbRecord := &corepb.Record{
		Hash:      payload.Hash,
		Owner:     tx.owner.Bytes(),
		Timestamp: tx.timestamp,
	}
	recordBytes, err := proto.Marshal(pbRecord)
	if err != nil {
		return err
	}

	return rs.Put(payload.Hash, recordBytes)
}

//VestTx is a structure for withdrawing vesting
type VestTx struct {
	user   common.Address
	amount *util.Uint128
}

//NewVestTx returns NewTx
func NewVestTx(tx *Transaction) (ExecutableTx, error) {
	return &VestTx{
		user:   tx.From(),
		amount: tx.Value(),
	}, nil
}

//Execute VestTx
func (tx *VestTx) Execute(b *Block) error {
	as := b.state.AccState()
	cs := b.state.DposState().CandidateState()

	if err := as.SubBalance(tx.user.Bytes(), tx.amount); err != nil {
		return err
	}
	if err := as.AddVesting(tx.user.Bytes(), tx.amount); err != nil {
		return err
	}

	account, err := as.getAccount(tx.user.Bytes())
	if err != nil {
		return err
	}

	// Check user voted
	candidate := account.Voted()
	if candidate == nil {
		return nil
	}

	// Add user's vesting to candidate's votePower
	candidateBytes, err := cs.Get(candidate)
	if err != nil {
		return nil
	}
	pbCandidate := new(dpospb.Candidate)
	if err := proto.Unmarshal(candidateBytes, pbCandidate); err != nil {
		return err
	}
	votePower, err := util.NewUint128FromFixedSizeByteSlice(pbCandidate.VotePower)
	if err != nil {
		return err
	}
	newVotePower, err := votePower.Add(tx.amount)
	if err != nil {
		return err
	}
	pbCandidate.VotePower, err = newVotePower.ToFixedSizeByteSlice()
	if err != nil {
		return err
	}
	newCandidateBytes, err := proto.Marshal(pbCandidate)
	if err != nil {
		return err
	}

	return cs.Put(candidate, newCandidateBytes)
}

//WithdrawalVestTx is a structure for withdrawing vesting
type WithdrawalVestTx struct {
	user   common.Address
	amount *util.Uint128
}

//NewWithdrawalVestTx returns WithdrawalVestTx
func NewWithdrawalVestTx(tx *Transaction) (ExecutableTx, error) {
	return &WithdrawalVestTx{
		user:   tx.From(),
		amount: tx.Value(),
	}, nil
}
//Todo Naming
//Execute WithdrawalVestTx
func (tx *WithdrawalVestTx) Execute(b *Block) error {
	as := b.state.AccState()
	cs := b.state.DposState().CandidateState()

	//if err := as.AddBalance(tx.user.Bytes(), tx.amount); err != nil {
	//	return err
	//}
	//if err := as.SubBalance(tx.user.Bytes(), tx.amount); err != nil {
	//	return err
	//}

	account, err := as.getAccount(tx.user.Bytes())
	if err != nil {
		return err
	}

	if tx.amount.Cmp(account.Vesting()) > 0 {
		return ErrVestingNotEnough
	}

	splitAmount, err := tx.amount.Div(util.NewUint128FromUint(RtWithdrawNum))
	if err != nil {
		return err
	}

	amountLeft := tx.amount.DeepCopy()

	payload := new(RtWithdraw)
	for i := 0; i < RtWithdrawNum; i++ {
		if amountLeft.Cmp(splitAmount) <= 0 {
			payload, err = NewRtWithdraw(amountLeft)
			if err != nil {
				return err
			}
		} else {
			payload, err = NewRtWithdraw(splitAmount)
			if err != nil {
				return err
			}
		}
		task := NewReservedTask(RtWithdrawType, tx.user, payload, b.Timestamp()+int64(i+1)*RtWithdrawInterval)
		if err := b.state.AddReservedTask(task); err != nil {
			return err
		}
		amountLeft, _ = amountLeft.Sub(splitAmount)
	}

	// Check user voted
	candidate := account.Voted()
	if candidate == nil {
		return nil
	}

	// Add user's vesting to cadidate's votePower
	candidateBytes, err := cs.Get(candidate)
	if err != nil {
		return nil
	}
	pbCandidate := new(dpospb.Candidate)
	if err := proto.Unmarshal(candidateBytes, pbCandidate); err != nil {
		return err
	}
	votePower, err := util.NewUint128FromFixedSizeByteSlice(pbCandidate.VotePower)
	if err != nil {
		return err
	}
	newVotePower, err := votePower.Sub(tx.amount)
	if err != nil {
		return err
	}
	pbCandidate.VotePower, err = newVotePower.ToFixedSizeByteSlice()
	if err != nil {
		return err
	}
	newCandidateBytes, err := proto.Marshal(pbCandidate)
	if err != nil {
		return err
	}

	return cs.Put(candidate, newCandidateBytes)
}

//AddCertificationTx is a structure for adding certification
type AddCertificationTx struct {
	Issuer    common.Address
	Certified common.Address
	Payload   *AddCertificationPayload
}

// AddCertificationPayload is payload type for AddCertificationTx
type AddCertificationPayload struct {
	IssueTime       int64
	ExpirationTime  int64
	CertificateHash []byte
}

//NewAddCertificationTx returns AddCertificationTx
func NewAddCertificationTx(tx *Transaction) (ExecutableTx, error) {
	payload := new(AddCertificationPayload)
	if err := json.Unmarshal(tx.Data(), payload); err != nil {
		return nil, ErrInvalidTxPayload
	}
	//TODO: certification payload Verify: drsleepytiger
	return &AddCertificationTx{
		Issuer:    tx.From(),
		Certified: tx.To(),
		Payload:   payload,
	}, nil
}

//Execute AddCertificationTx
func (tx *AddCertificationTx) Execute(b *Block) error {
	as := b.state.AccState()
	cs := b.state.certificationState

	if err := as.AddCertReceived(tx.Certified.Bytes(), tx.Payload.CertificateHash); err != nil {
		return err
	}
	if err := as.AddCertIssued(tx.Issuer.Bytes(), tx.Payload.CertificateHash); err != nil {
		return err
	}
	pbCertification := &corepb.Certification{
		CertificateHash: tx.Payload.CertificateHash,
		Issuer:          tx.Issuer.Bytes(),
		Certified:       tx.Certified.Bytes(),
		IssueTime:       tx.Payload.IssueTime,
		ExpirationTime:  tx.Payload.ExpirationTime,
		RevocationTime:  int64(0),
	}
	certificationBytes, err := proto.Marshal(pbCertification)
	if err != nil {
		return err
	}
	return cs.Put(tx.Payload.CertificateHash, certificationBytes)
}

//RevokeCertificationTx is a structure for revoking certification
type RevokeCertificationTx struct {
	Revoker    common.Address
	Payload   *RevokeCertificationPayload
}

// RevokeCertificationPayload is payload type for RevokeCertificationTx
type RevokeCertificationPayload struct {
	CertificateHash []byte
	RevocationTime int64
}

//NewRevokeCertificationTx returns RevokeCertificationTx
func NewRevokeCertificationTx(tx *Transaction) (ExecutableTx, error) {
	payload := new(RevokeCertificationPayload)
	if err := json.Unmarshal(tx.Data(), payload); err != nil {
		return nil, ErrInvalidTxPayload
	}
	//TODO: certification payload Verify: drsleepytiger
	return &RevokeCertificationTx{
		Revoker:    tx.From(),
		Payload:   payload,
	}, nil
}

//Execute RevokeCertificationTx
func (tx *RevokeCertificationTx) Execute(b *Block) error {
	//as := b.state.AccState()
	cs := b.state.certificationState

	certificationBytes, err := cs.Get(tx.Payload.CertificateHash)
	if err != nil {
		return nil
	}

	pbCertification := new(corepb.Certification)
	if err := proto.Unmarshal(certificationBytes, pbCertification); err != nil {
		return err
	}
	if common.BytesToAddress(pbCertification.Issuer) != tx.Revoker {
		return ErrInvalidCertificationRevoker
	}
	if pbCertification.RevocationTime > int64(0) {
		return ErrCertAlreadyRevoked
	}

	pbCertification.RevocationTime = tx.Payload.RevocationTime
	newBytesCertification, err := proto.Marshal(pbCertification)
	if err != nil {
		return err
	}

	//Todo: Remove certificate from AccountState(both issuer and certified)
	return cs.Put(tx.Payload.CertificateHash, newBytesCertification)
}

//Todo : move to another file