# Smart Contract for Ownership Transfer (SCOT)

This repository implements the **CLI component**, a software that imports the cryptographic library and extends its properties by creating interactive user functions for the SCOT system.

SCOT is a proof of concept that implements the abstract model and the cryptographic construction devised in the thesis [IMPROVED DISTRIBUTED LEDGER TRANSACTIONS WITH HOMOMORPHIC COMPUTATIONS](https://uccs.idm.oclc.org/login?url=https://www.proquest.com/dissertations-theses/improved-distributed-ledger-transactions-with/docview/2469874959/se-2?accountid=25388).

The concept is explained at page 86, in chapter 5 (i.e., Software Architecture of SCOT).

## How To Run The Example

The complete system uses 3 repositories:

1. A cryptographic library [phe](https://github.com/hanesbarbosa/phe);
2. A Command Line Interface (CLI) [phe-cli](https://github.com/hanesbarbosa/phe-cli) that implements the cryptographic functions for the user;
3. A small Hyperledger Fabric blockchain [fabric-samples](https://github.com/hanesbarbosa/fabric-samples).

Section 5.3, from chapter 5 (i.e., Application) describes all necessary steps to run the example. For troubleshooting and installation tips, follow the instructions at the APPENDIX A (i.e., Setup And Configuration Guides).
