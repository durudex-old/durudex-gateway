# Certificates

To generate certificates you need:
1) Install [mkcert](https://github.com/FiloSottile/mkcert).
2) Add local domains:
```
authservice.durudex.local
notifservice.durudex.local
```
3) And use `make cert`

#### If an error occurs `$'/r': command not found` use:
```sh
sed -i 's/\r$//' scripts/generate-cert.sh
```
