#!/bin/bash

#input args
# $1 - key file path 
# $2 - cert file path 
# $3 - key label 
# $4 - key id 
# $5 - token pin 

#User pin
upin=$5
#key id in hex
key_id=$(printf '%x' $4)
#Key_label used for key import
key_label=$3
#Initial slot number
slot_no="0"
#Token name
token_name="CAToken"
#SoftHSM2 lib location
libpath="/usr/lib/x86_64-linux-gnu/softhsm/libsofthsm2.so"

#Initialize the token
softhsm2-util --init-token --slot ${slot_no} --label "${token_name}" \
    --pin ${upin} --so-pin ${upin}

openssl pkcs8 -inform PEM -outform DER -in $1 -out /tmp/prikey.der -nocrypt
#Import the private key and cert into SoftHSM
pkcs11-tool --module $libpath -l --pin ${upin} \
    --write-object /tmp/prikey.der --type privkey --id ${key_id} --label ${key_label} 
rm -f /tmp/prikey.der

openssl x509 -inform PEM -in $2 -outform PEM -pubkey -noout >& /tmp/pubkey.pem
openssl rsa -pubin -inform PEM -in /tmp/pubkey.pem -outform DER -out /tmp/pubkey.der 

pkcs11-tool --module $libpath -l --pin ${upin} \
    --write-object /tmp/pubkey.der --type pubkey --id ${key_id} --label ${key_label}

openssl x509 -in $2 -outform der -out /tmp/ca.der
pkcs11-tool --module $libpath -l --pin ${upin} \
    --write-object /tmp/ca.der --type cert --id ${key_id} --label ${key_label}

rm -f /tmp/out
rm -f /tmp/pubkey.*
rm -f /tmp/ca.der

#Create config file
cat <<EOF > /tmp/config
{
  "Path" : "$libpath",
  "TokenLabel": "$token_name",
  "Pin" : "$upin" 
}
EOF

