# KmsEnv

kmsenv is a command that reads the .env file and encrypts the environment with `Cloud Key Management Service`.

## Installation

```
go get github.com/locona/kmsenv/cli/kmsenv
```

## Usage

Add your application configuration to your `.env` file in the root of your project:

```
S3_BUCKET=YOURS3BUCKET
SECRET_KEY=YOURSECRETKEYGOESHERE
```

and `.kmsenvrc` file in the root of your project:

```
{
	"project_id": "xxxxx",
	"location": "global",
	"keyring": "xxxxxxx",
	"key": "xxx"
}
```

And then run:
```
kmsenv
```

As a result created the `kms-encrypt.json` file

```
{
	"S3_BUCKET": "CiQA66hgkjw9CbdF1ZKGen3SQNwtFoLOR2paVg+jLFj3B+iDq98SNQDkjMn2WIrF0C8UNgVHAy/knG9DORLvZjaX6DO8ha039AYMGsrFDoizfEzX1em9LPp2CAQT",
	"SECRET_KEY": "CiQA66hgkhwMu1/8cknFFt12XsXuTls+TWuo2/CojJS7YwEoAVESPgDkjMn29Y0UHIUkHQCypRD3Io1BC5ObnC3INKOX/hrUZ9O2ulPjNaRIZRs1w57LiIXPJku4qvUIegkfz33G"
}
```
