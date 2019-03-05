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
libpath="/usr/local/lib/softhsm/libsofthsm2.so"

#Initialize the token
softhsm2-util --init-token --slot ${slot_no} --label "${token_name}" \
    --pin ${upin} --so-pin ${upin}

#Import the private key and cert into SoftHSM
pkcs11-tool --module $libpath -l --pin ${upin} \
    --write-object $1 --type privkey --id ${cert_id} --label ${key_label}

pkcs11-tool --module $libpath -l --pin ${upin} \
--write-object $2 --type cert --id ${cert_id}

#Create config file
cat <<EOF > /tmp/config
{
  "Path" : $libpath,
  "TokenLabel": $token_name,
  "Pin" : $upin 
}
EOF

