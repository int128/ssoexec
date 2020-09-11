# ssoexec ![test](https://github.com/int128/ssoexec/workflows/test/badge.svg)

This is a lightweight command to run a third-party tool with AWS SSO.
It is written in Go and has no dependency.


## Why

Most of third-party tools such as Terraform do not support AWS SSO.
To run a third-party tool, you need to acquire short-term credentials from AWS SSO
and set them to the environment variables.

ssoexec allows you to run a third-party tool with AWS SSO.


## Getting Started

Install the latest release.

```console
% go get github.com/int128/ssoexec
```

If needed, you can set the current profile.

```console
% export AWS_PROFILE=example
```

Log in via AWS SSO.

```console
% aws sso login
Attempting to automatically open the SSO authorization page in your default browser.
If the browser does not open or you wish to use a different device to authorize this request, open the following URL:

https://device.sso.us-east-1.amazonaws.com/

Then enter the code:

****-****
Successully logged into Start URL: https://********.awsapps.com/start
```

Run a command with ssoexec.

```console
% ssoexec terraform
```

As well as you can export the environment variables.

```console
% eval $(ssoexec)
```


## How it works

ssoexec respects the following environment variables:

- `AWS_PROFILE` (current profile name)

It loads the following parameters:

- `~/.aws/config`
  - Region, account ID and role name for SSO
- `~/.aws/sso/cache`
  - Access token retrieved via AWS SSO

Then it acquires short-term credentials by the following command:

```sh
aws sso get-role-credentials --role-name SSORoleName --region SSORegion --account-id SSOAccountID --access-token AccessToken
```

Finally it exports the following environment variables:

- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `AWS_SESSION_TOKEN`


## Contributions

This is an open source software.
Feel free to open issues and pull requests.
