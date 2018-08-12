# Go Interswitch

This library provides an API written in Go programming language 
for transacting payments using the Interswitch REST API.

# Usage

## Get Payment Methodss

    conf := interswitch.Conf{
        endpoint: "https://sandbox.interswitchng.com/",
        schemeName: "scheme-name",
        clientId: "client-id",
        sharedSecret: "shared-secret",
        signatureMethod: "SHA512",
        channel: "channel",
        version: "4.6"
    }

    client := interswitch.Client(&conf)
    paymentMethods := client.PaymentMethods()



# Contribution guidelines

See [CONTRIBUTING.md](./CONTRIBUTING.md)