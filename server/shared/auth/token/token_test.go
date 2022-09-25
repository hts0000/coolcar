package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAu1SU1LfVLPHCozMxH2Mo
4lgOEePzNm0tRgeLezV6ffAt0gunVTLw7onLRnrq0/IzW7yWR7QkrmBL7jTKEn5u
+qKhbwKfBstIs+bMY2Zkp18gnTxKLxoS2tFczGkPLPgizskuemMghRniWaoLcyeh
kd3qqGElvW/VDL5AaWTg0nLVkjRo9z+40RQzuVaE8AkAFmxZzow3x+VJYKdjykkJ
0iT9wCS0DRTXu269V264Vf/3jvredZiKRkgwlL9xNAwxXFg0x/XFw005UWVRIkdg
cKWTjpBP2dPwVZ4WWC+9aGVd+Gyn1o0CLelf4rEjGoXbAAEgAqeGUxrcIlbjXfbc
mwIDAQAB
-----END PUBLIC KEY-----`

func TestVerify(t *testing.T) {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Fatalf("cannot parse public key: %v", err)
	}
	v := &JWTTokenVerifier{
		PublicKey: pubKey,
	}

	cases := []struct {
		name    string
		tkn     string
		now     time.Time
		want    string
		wantErr bool
	}{
		{
			name:    "valid_token",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjM5MjcyMDAsImlhdCI6MTY2MzkyMDAwMCwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjMyYjFjNmUxMzBmNTBjMjc0ODEzN2FkIn0.cq3U52YLD5QlyDaB0J2qkUggbWvLeuf4DR6S7HJ1p1ddj4qxRzwfFSAqbbzSXPm9wGTb3FXNCRPMJM6fi0OvQFewiZ2IhiVVPqLl_HPAHMT8jHWjVKezh1ZY2TNL5f1x1TUkeGQapu0rc4kN2_WDStAnoKVbo7MvOsXAwnTNjOBCMEjc8axk2lJfvb25dHEEgQsoy9l0G8OH2PVWwViGtj-qrSz5AOfgVTCBHB6w02PUvP0zEvcV_v-7Aa_0y_RywhzmnQdh1ROQNVEPAIQP9-dNNWLQy8VvISCf6hOqph5jouGheWsQqFRoX3X-rv2EjZx_egC5dZ6gGenaBmiDkA",
			now:     time.Date(2022, time.September, 23, 18, 0, 0, 0, time.Local),
			want:    "632b1c6e130f50c2748137ad",
			wantErr: false,
		},
		{
			name:    "expired_token",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjM5MjcyMDAsImlhdCI6MTY2MzkyMDAwMCwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjMyYjFjNmUxMzBmNTBjMjc0ODEzN2FkIn0.cq3U52YLD5QlyDaB0J2qkUggbWvLeuf4DR6S7HJ1p1ddj4qxRzwfFSAqbbzSXPm9wGTb3FXNCRPMJM6fi0OvQFewiZ2IhiVVPqLl_HPAHMT8jHWjVKezh1ZY2TNL5f1x1TUkeGQapu0rc4kN2_WDStAnoKVbo7MvOsXAwnTNjOBCMEjc8axk2lJfvb25dHEEgQsoy9l0G8OH2PVWwViGtj-qrSz5AOfgVTCBHB6w02PUvP0zEvcV_v-7Aa_0y_RywhzmnQdh1ROQNVEPAIQP9-dNNWLQy8VvISCf6hOqph5jouGheWsQqFRoX3X-rv2EjZx_egC5dZ6gGenaBmiDkA",
			now:     time.Date(2022, time.September, 23, 18, 0, 1, 0, time.Local),
			want:    "",
			wantErr: true,
		},
		{
			name:    "bad_token",
			tkn:     "bad_token",
			now:     time.Date(2022, time.September, 23, 18, 0, 1, 0, time.Local),
			want:    "",
			wantErr: true,
		},
		{
			name:    "wrong_signature",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjM5MjcyMDAsImlhdCI6MTY2MzkyMDAwMCwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjMyYjFjNmUxMzBmNTBjMjc0ODEzN2FkIn0.DES7DKzbDP_OJyDUwG4vuzZERXqDmyB5SElKADXf0pkZMMM0frSYe9pG-bVAZ1eWsrky3UlIYK9JsQzoTEvmc4znxwQIaZwnQzJyn8ZGVafREcQj91sZO9Tp_jiOVlQcHBQUNgdAJRqVBp7gu6haXf_Tz5RrLcjDJEcnnWVjQs2F0-lzhq-zTgPG3U4U0uZAL_Tq7drFuTwPHBhI83YMNYCXDiau0DhYfCw3x6-Nb1ZNdnDWf9zw8UIk5zu1vpLULu_Mzkg0RX0tmy7abNYwNiRbn_UMiA4Cw75sVW7BRIbygXLwZ-2o-GEP6vVTBjMQtIBCXPij2_ZOQqefJ0GUKQ",
			now:     time.Date(2022, time.September, 23, 18, 0, 0, 0, time.Local),
			want:    "",
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			jwt.TimeFunc = func() time.Time {
				return c.now
			}
			accountID, err := v.Verify(c.tkn)
			if !c.wantErr && err != nil {
				t.Errorf("verification failed: %v", err)
			}
			if c.wantErr && err == nil {
				t.Errorf("want error, got no error")
			}
			if accountID != c.want {
				t.Errorf("wrong account id.\n want: %q,\ngot: %q", c.want, accountID)
			}
		})
	}
}
