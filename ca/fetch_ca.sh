FETCH_URL=$(grep 'origin_url' ./config/blocklist.yaml | awk '{print $2}')

if [ -z "$FETCH_URL" ]; then
    echo "Error: Failed to fetch the origin_url from blocklist.yaml" >&2
    exit 1
fi

if ! echo | openssl s_client -servername $FETCH_URL -connect rapla.dhbw-karlsruhe.de:443 2>/dev/null \
    | sed -n '/BEGIN CERTIFICATE/,/END CERTIFICATE/p' > dhbw_rapla_cert.pem; then
    echo "Error: Failed to fetch or process the certificate" >&2
    exit 1
fi