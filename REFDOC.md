Here is a list of URLs that document the original REST API.

- https://connect.interswitchng.com/docs/quickteller-3/payment-with-our-rest-apis/

- https://connect.interswitchng.com/docs/api-security/

- https://developer.interswitchng.com/console/index


# Consuming the REST API

Consuming the REST API involves three steps namely:

## Signing the request

### Prerequisites

- Trim whitespace from signature fields
- Percent encode string in uppercase letters

### Algorithm

    String baseStringToBeSigned = http_verb + "&" +
    percent_encode(url) + "&" +
    timestamp + "&" +
    nonce + "&" +
    client_id + "&" +
    shared_secret_key;

    String parameterStringToBeSigned = percent_encode(value1) + "&" + percent_encode(value2);

    String stringToBeSigned = baseStringToBeSigned + "&" + parameterStringToBeSigned;
    String signature = Base64(Hash(stringToBeSigned));

### Sample Request

    POST http://172.25.20.21/quicktellerservice/api/v3/transactions/doinquiry

    Authorization: InterswitchAuth Zkp3dmJ4dHA1VU12QkludFUxVTg3OExLQXU2TXdWWW0=
    timestamp: 1361281946
    nonce: 634968823463411609
    signature:EzU5qY9eCDL+0J/e2sYildgiBA8TgPsXHrkm2dx8kU12a1a7u813/pTRk52BymCVDPaOxX9p4Qn4ZxjAfFHhcQ==
    signatureMethod: SHA1
    terminalid: 3api0001
    Content-Type: application/json
    Host: 172.25.20.21
    Content-Length: 158

    { “paymentcode”: “10402″, “customerid”: “0000000001″, “cardpan”: “627629020217176055″, “customeremail”: “eawagu@gmail.com”, “customermobile”: “08035782209″ }

## Package your request

E.g.
- set paymentitemcode
- set amount

## Verify your response

    Confirm that signature in response header matches signature computed from response body


# Authorization
Set header `Authorization: InterswitchAuth Base64(client_id)`
