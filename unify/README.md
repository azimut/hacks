# unify
gets domains from stdin and extract all the subdomains and returns an unique list of them. Without a main domain, only subdomains. And with each level of subdomains.

```
-*- mode: compilation; default-directory: "~/go/src/github.com/azimut/purify/unify/" -*-
Compilation started at Fri Jun 19 08:27:02

echo -en 'www.starbucks.com\nftp.google.com\ndev.intel.cia.gov\n' | go run /home/sendai/go/src/github.com/azimut/purify/unify/
www
ftp
intel
dev.intel

Compilation finished at Fri Jun 19 08:27:02
```

