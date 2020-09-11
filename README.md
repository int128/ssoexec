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

ssoexec executes the following command to acquire short-term credentials:

```sh
aws sso get-role-credentials --role-name SSORoleName --region SSORegion --account-id SSOAccountID --access-token AccessToken
```

It exports the following environment variables:

- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `AWS_SESSION_TOKEN`

It loads the SSO related parameters of the current profile from `~/.aws/config`.
It loads the access token from `~/.aws/sso/cache`.


## Contributions

This is an open source software.
Feel free to open issues and pull requests.
