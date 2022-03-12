# Certificates

To generate certificates you need:
1) Install [mkcert](https://github.com/FiloSottile/mkcert).
2) Add local domains:
```
api.durudex.local
auth.service.durudex.local
email.service.durudex.local
user.service.durudex.local
code.service.durudex.local
```
3) And use `make cert`

**If you do not want to use tls connection change [service.name.tls](https://github.com/durudex/durudex-gateway/blob/main/configs/main.yml) configuration to `false`**:
```yml
service:
    name:
        addr: service.local
        tls: false
```

#### If an error occurs `$'/r': command not found` use:
```sh
sed -i 's/\r$//' scripts/generate-cert.sh
```
