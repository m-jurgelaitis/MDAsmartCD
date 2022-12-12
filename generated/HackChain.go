package chaincode

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type HackChain struct {
	contractapi.Contract
}

const nameKey = "Certificate"
const symbolKey = "HackChain Certificate"

var tokenCounter uint

type Certificate struct {
	TokenId                uint   `json:"tokenId"`
	Name                   string `json:"name"`
	Surname                string `json:"surname"`
	Hackathon              string `json:"hackathon"`
	StartDate              uint   `json:"startDate"`
	EndDate                uint   `json:"endDate"`
	HackathonParticipantId uint   `json:"hackathonParticipantId"`
	HackathonId            uint   `json:"hackathonId"`
	SolutionId             uint   `json:"solutionId"`
	SolutionValid          bool   `json:"solutionValid"`
	State                  string `json:"state"`
	Owner                  string `json:"owner"`
}
type CertificateCreated struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Hackathon  string `json:"hackathon"`
	StartDate  uint   `json:"startDate"`
	EndDate    uint   `json:"endDate"`
	SolutionId uint   `json:"solutionId"`
	State      string `json:"state"`
}

func (s *HackChain) createCertificate(ctx contractapi.TransactionContextInterface, _name string, _surname string, _hackathon string, _startDate uint, _endDate uint, _hackathonParticipantId uint, _hackathonId uint, _solutionValid bool, _solutionId uint) error {
	tokenCounter++
	MintWithTokenURI(ctx, tokenCounter, "http://example.org/Certificate")
	certificate := Certificate{tokenCounter, _name, _surname, _hackathon, _startDate, _endDate, _hackathonParticipantId, _hackathonId, _solutionId, _solutionValid, "ISSUED", ""}
	certificateJSON, err := json.Marshal(certificate)
	if err != nil {
		return err
	}
	var stub = ctx.GetStub()
	return stub.PutState(tokenCounter, certificateJSON)
	var certificateCreated = CertificateCreated{_name, _surname, _hackathon, _startDate, _endDate, _solutionId, "ISSUED"}
	certificateCreatedJSON, err := json.Marshal(certificateCreated)
	if err != nil {
		return err
	}
	return stub.PutState(tokenCounter, certificateCreatedJSON)
}
func (s *HackChain) confirmEvaluationResults(ctx contractapi.TransactionContextInterface, tokenId uint, _solutionId uint, _solutionValid bool) error {
	var stub = ctx.GetStub()
	certificateJSON, err := stub.GetState(tokenId)
	if err != nil {
		return nil, err
	}
	var certificate Certificate
	err = json.Unmarshal(certificateJSON, certificate)
	if err != nil {
		return nil, err
	}
	certificate.SolutionValid = _solutionValid
	certificate.SolutionId = _solutionId
	if certificate.State == "ISSUED" && _solutionValid == true && _solutionId != 0 {
		certificate.State = "ISSUED_FOR_SOLUTION"
	} else {
		certificate.State = "ISSUED_FOR_PARTICIPATION"
	}
	certificateJSON, err = json.Marshal(certificate)
	if err != nil {
		return err
	}
	return stub.PutState(tokenId, certificateJSON)
}
