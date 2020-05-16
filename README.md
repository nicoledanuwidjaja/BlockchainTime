# Blockchain Implementation

https://bitcoin.org/bitcoin.pdf

### Definition
A Block is a data structure that comprises of the previous block hash (PrevHash), its own hash (Hash), and information relevant to the transaction (Data).

Each Block is linked together like a back-linked list, where each Block references to the previous Block with the PrevHash field.

Hashes are generated with the SHA-256 hash algorithm.
