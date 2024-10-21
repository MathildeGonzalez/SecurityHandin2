# Security Handin 2

## Additive secret sharing using TLS and certificates

### How to run

To run the program, you need to have Go installed on your machine.

To create the certificates, run the following command in the terminal:
```bash
make cert
```

To run the program, you need to open 4 terminals.

Terminal 1:

```bash
make hospital
```

Terminal 2:

```bash
make alice
```

Terminal 3:

```bash
make bob
```

Terminal 4:

```bash
make charlie
```

There is some waiting time encoded in the program to give enough time to start all the processes, so please be patient.

### Example output for Alice

```bash
2024/10/14 13:30:39 Alice is running on localhost:5001
2024/10/14 13:30:49 Calculated shares: 8721147990643752020, 7339549505873501197, 2386046577192298400
2024/10/14 13:30:54 Received share from other patient: 6427883394960236832
2024/10/14 13:30:59 Sending share to Bob: 7339549505873501197
2024/10/14 13:30:59 Acknowledgment from Bob: Share received
2024/10/14 13:30:59 Sending share to Charlie: 2386046577192298400
2024/10/14 13:30:59 Acknowledgment from Charlie: Share received
2024/10/14 13:31:01 Received share from other patient: 1201670889690606359
2024/10/14 13:31:01 Sum of shares: -2096041798414956405
2024/10/14 13:31:01 Sending aggregated share to hospital: -2096041798414956405
2024/10/14 13:31:01 Acknowledgment from Hospital: Received aggregated share
```
