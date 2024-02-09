<p align="center">
	<img src="purpl.png" width=50" />
</p>

# jwt-go-ecdsa
This Go module generates JWTs using ECDSA encryption, tailored to specific services and purposes as defined in a policy file.
It dynamically adjusts token claims based on the policy, including permissions and conditions for the token's use, then
signs it with a private EC key. The token's expiration is set according to the specified duration.

The function structure is:
```
GenerateToken(policyPath string, serviceName string, purpose string, keyPath string, expirationInHours time.Duration)
```

The key should be a ECDSA-256 private key and the JSON structure of the policy should be:

```
{
  "services": [
    {
      "service1": {
        "purpose1": {
          "allowed":
          {
            "field1": [
              "string"
            ],
            "field2": [
              "string"
            ],
            "field3": [
              "string"
            ]
          },
          "generalized":
          {
            "field1": [
              "string",
              "parameter if necessary"
            ],
            "field2": [
              "string",
              "parameter if necessary"
            ],
            "field3": [
              "string",
              "parameter if necessary"
            ]
          },
          "noised":
          {
            "field1": [
              "string",
              "parameter if necessary"
            ],
            "field2": [
              "string",
              "parameter if necessary"

            ],
            "field3": [
              "string",
              "parameter if necessary"
            ]
          },
          "reduced":
          {
            "field1": [
              "string",
              "parameter if necessary"
            ],
            "field2": [
              "string",
              "parameter if necessary"
            ],
            "field3": [
              "string",
              "parameter if necessary"
            ]
          }
        },
        "purpose2": {
          ...
        }
      },
      "service2": {
        ...
      },
      ...
  ]
}
```

Example:
You can find an example [here](https://github.com/PEngG7/jwt-go-ecdsa/blob/main/policy.json).


# Usage

To use this module run:
```shell
go get -u github.com/PEngG7/jwt-go-ecdsa@v0.1.0
``` 

and add this import statement to your Go file:
```go
import ("github.com/PEngG7/jwt-go-ecdsa")
```

# Testing
The test.go file contains a test for the GenerateToken function. It uses the policy.json file and the private key
provided in this repo. The provided test generates the following token:

```
eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJwb2xpY3kiOnsiYWxsb3dlZCI6e30sImdlbmVyYWxpemVkIjp7ImNpdHkiOlsic3RyaW5nIiwiMiJdLCJjcmVkaXRfY2FyZF9jdnYiOlsiaW50IiwiMyJdLCJjcmVkaXRfY2FyZF9leHBpcmF0aW9uX3llYXIiOlsiaW50IiwiMTAiXSwiY3JlZGl0X2NhcmRfbnVtYmVyIjpbInN0cmluZyIsIjUiXSwiemlwX2NvZGUiOlsiaW50IiwiOCJdfSwibm9pc2VkIjp7ImFnZSI6WyJpbnQiLCJMYXBsYWNlIl0sImNyZWRpdF9jYXJkX2V4cGlyYXRpb25fbW9udGgiOlsiaW50IiwiTGFwbGFjZSJdLCJzdHJlZXRfbmFtZSI6WyJzdHJpbmciLCJMYXBsYWNlIl0sInN0cmVldF9udW1iZXIiOlsiaW50IiwiTGFwbGFjZSJdfSwicmVkdWNlZCI6eyJjb3VudHJ5IjpbInN0cmluZyIsIjMiXSwiZW1haWwiOlsic3RyaW5nIiwiNCJdLCJuYW1lIjpbInN0cmluZyIsIjQiXSwicGhvbmUiOlsic3RyaW5nIiwiMyJdfX0sImlzcyI6InRva2VuR2VuZXJhdG9yIiwiZXhwIjoxNzA3NTA5OTE5fQ.F0nzF6clbMLxyOSFfeBHCXOEpHC1nWQGRYThA_vt3_nsI8gaYW8slupbAhc4EwFVHkx1knleX14Vj2UQyJHVOw
```

The content can be decoded using a JWT decoder, such as [jwt.io](https://jwt.io/).

In this case it looks like this:
HEADER
```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

PAYLOAD
```json
{
  "policy": {
    "allowed": {},
    "generalized": {
      "city": [
        "string",
        "2"
      ],
      "credit_card_cvv": [
        "int",
        "3"
      ],
      "credit_card_expiration_year": [
        "int",
        "10"
      ],
      "credit_card_number": [
        "string",
        "5"
      ],
      "zip_code": [
        "int",
        "8"
      ]
    },
    "noised": {
      "age": [
        "int",
        "Laplace"
      ],
      "credit_card_expiration_month": [
        "int",
        "Laplace"
      ],
      "street_name": [
        "string",
        "Laplace"
      ],
      "street_number": [
        "int",
        "Laplace"
      ]
    },
    "reduced": {
      "country": [
        "string",
        "3"
      ],
      "email": [
        "string",
        "4"
      ],
      "name": [
        "string",
        "4"
      ],
      "phone": [
        "string",
        "3"
      ]
    }
  },
  "iss": "tokenGenerator",
  "exp": 1707509919
}
```

