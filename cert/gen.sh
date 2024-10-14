rm *.pem

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=DK/ST=Zealand/L=Copenhagen/O=ITU/OU=Education/CN=localhost/emailAddress=ca@itu.com"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2. Generate hospital server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout hospital-server-key.pem -out hospital-server-req.pem -subj "/C=DK/ST=Zealand/L=Copenhagen/O=Hospital/OU=Hospital/CN=localhost/emailAddress=hospital@hospital.com"

# 3. Use CA's private key to sign hospital server's CSR and get back the signed certificate
openssl x509 -req -in hospital-server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out hospital-server-cert.pem -extfile server-ext.cnf

echo "Server's signed certificate"
openssl x509 -in hospital-server-cert.pem -noout -text

# 4. Generate Alices's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout alice-key.pem -out alice-req.pem -subj "/C=DK/ST=Zealand/L=Copenhagen/O=Patient/OU=Patient/CN=localhost/emailAddress=alice@gmail.com"

# 5. Use CA's private key to sign Alice's CSR and get back the signed certificate
openssl x509 -req -in alice-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out alice-cert.pem -extfile client-ext.cnf

echo "Alice's signed certificate"
openssl x509 -in alice-cert.pem -noout -text

# 6. Generate Bob's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout bob-key.pem -out bob-req.pem -subj "/C=DK/ST=Zealand/L=Copenhagen/O=Patient/OU=Patient/CN=localhost/emailAddress=bob@gmail.com"

# 7. Use CA's private key to sign Bob's CSR and get back the signed certificate
openssl x509 -req -in bob-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out bob-cert.pem -extfile client-ext.cnf

echo "Bob's signed certificate"
openssl x509 -in bob-cert.pem -noout -text

# 8. Generate Charlie's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout charlie-key.pem -out charlie-req.pem -subj "/C=DK/ST=Zealand/L=Copenhagen/O=Patient/OU=Patient/CN=localhost/emailAddress=charlie@gmail.com"

# 9. Use CA's private key to sign Charlie's CSR and get back the signed certificate
openssl x509 -req -in charlie-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out charlie-cert.pem -extfile client-ext.cnf

echo "Charlie's signed certificate"
openssl x509 -in charlie-cert.pem -noout -text