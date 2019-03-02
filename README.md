# Go Interswitch

This library provides an API written in Go programming language 
for transacting payments using the Interswitch REST API.

# Usage

## Get Payment Methods

    conf := interswitch.Conf{
        endpoint: "https://sandbox.interswitchng.com/",
        schemeName: "scheme-name",
        clientId: "client-id",
        sharedSecret: "shared-secret",
        signatureMethod: "SHA512",
        channel: "channel",
        version: "4.6"
    }

    c, err := interswitch.ClientUsingConf(&conf)
    p := client.PaymentMethods()


## Testing

In your terminal run

    make test


# Contribution guidelines

See [CONTRIBUTING.md](./CONTRIBUTING.md)