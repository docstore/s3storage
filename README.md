# s3storage

Aws S3 storage implementation of [docstore.Storer](https://github.com/docstore/service)


## Installation
```
go get github.com/dmashuda/service
go get github.com/docstore/s3storage
```


### Storage Creation
```
storer := BasicAws("us-east-1", "dmashuda-dev")
```


## Configuring Credentials For AWS

Before using the SDK, ensure that you've configured credentials. The best
way to configure credentials on a development machine is to use the
`~/.aws/credentials` file, which might look like:

```
[default]
aws_access_key_id = AKID1234567890
aws_secret_access_key = MY-SECRET-KEY
```

You can learn more about the credentials file from this
[blog post](http://blogs.aws.amazon.com/security/post/Tx3D6U6WSFGOK2H/A-New-and-Standardized-Way-to-Manage-Credentials-in-the-AWS-SDKs).

Alternatively, you can set the following environment variables:

```
AWS_ACCESS_KEY_ID=AKID1234567890
AWS_SECRET_ACCESS_KEY=MY-SECRET-KEY
```

## Contributing
https://github.com/docstore/s3storage/graphs/contributors
 - Pull requests welcome
 - Feel free to add new docstore.Storer Implementations


## License

Released under the MIT License
