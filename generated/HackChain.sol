// SPDX-License-Identifier: MIT
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
pragma solidity ^0.8.0;

contract HackChain is ERC721 {
    uint private tokenCounter;
    mapping(uint => Certificate) public certificates;
    enum CertificateState {
        ISSUED, 
        ISSUED_FOR_SOLUTION,
        ISSUED_FOR_PARTICIPATION
    }
    struct Certificate {
        uint tokenId;
        string name;
        string surname;
        string hackathon;
        uint startDate;
        uint endDate;
        uint hackathonParticipantId;
        uint hackathonId;
        uint solutionId;
        bool solutionValid;
        CertificateState state;
        address owner;
    }
    event CertificateCreated(
        string name, 
        string surname,
        string hackathon,
        uint startDate, 
        uint endDate,       
        uint solutionId, 
        CertificateState state
    );
    constructor() ERC721("Certificate", "HackChain Certificate") {}
    function createCertificate(string memory _name, string memory _surname, string memory _hackathon, uint _startDate, uint _endDate, uint _hackathonParticipantId, uint _hackathonId, bool _solutionValid, uint _solutionId) external {
        tokenCounter++;
        _mint(msg.sender, tokenCounter);
        Certificate memory certificate = Certificate(tokenCounter, _name, _surname, _hackathon, _startDate, _endDate, _hackathonParticipantId, _hackathonId, _solutionId, _solutionValid, CertificateState.ISSUED, msg.sender);
        certificates[tokenCounter] = certificate;
        emit CertificateCreated(_name, _surname, _hackathon, _startDate, _endDate, _solutionId, CertificateState.ISSUED);
    }
    function confirmEvaluationResults(uint tokenId, uint _solutionId, bool _solutionValid) external {
        Certificate storage certificate = certificates[tokenId];
        certificate.solutionValid = _solutionValid;
        certificate.solutionId = _solutionId;
        if (certificate.state == CertificateState.ISSUED && _solutionValid == true &&_solutionId != 0) {
            certificate.state = CertificateState.ISSUED_FOR_SOLUTION;
        } else {
            certificate.state = CertificateState.ISSUED_FOR_PARTICIPATION;
        }
    }
}