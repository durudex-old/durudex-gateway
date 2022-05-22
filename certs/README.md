# Certificates

You need to move your certificates to this directory, and make configuration changes:
```yml
service:
  name:
    tls:
      ca-cert: "./certs/you-ca-cert.pem"
      cert: "./certs/you-cert.pem"
      key: "./certs/you-key.pem"
```

**If you do not want to use tls connection change**:
```yml
service:
  name:
    tls:
      enable: false
```

## Generate Certificate

1) Install [mkcert](https://github.com/FiloSottile/mkcert).
2) And use `make cert`

#### If an error occurs `$'/r': command not found` use:
```sh
sed -i 's/\r$//' scripts/generate-cert.sh
```
