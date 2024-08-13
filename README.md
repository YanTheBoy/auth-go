Just an example and educational layout

Generate Pem-encoded Edwards curve priv/pub keys
priv: openssl genpkey -algorithm ed25519 -out private.pem
pub: openssl pkey -in private.pem -pubout -out public.pem
