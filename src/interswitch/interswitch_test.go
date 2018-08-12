package interswitch_test

import (
       "interswitch"
       "testing"
)

var conf = interswitch.Conf{
                Endpoint: "https://api.sandbox.interswitchng.com/v3/quickteller/wallet/",
                SchemeName: "scheme-name",
                ClientId: "client-id",
                SharedSecret: "shared-secret",
                SignatureMethod: "SHA512",
                Channel: "channel",
                Version: "4.6",
        }

var client = interswitch.Client(&conf)

func TestPaymentMethods(t *testing.T) {
	// Different allocations should not be equal.
	
       

        if  _, err := client.PaymentMethods(); err != nil {
                t.Errorf("Payment Methods failed: %v", err)
        }
}
