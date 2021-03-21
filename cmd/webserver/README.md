# webserver

This demonstrates:
* a very simple web api created using Gin
* the repository pattern
* simple TLS usage in Gin

## Repository pattern
The web server can store contact data in a json file or a sqlite database. The code uses the Repository interface provided by the contacts package. You can choose the actual repository inmplementation by changing the `getrepo` function defined at the end of *webserver.go*.

## TLS Usage

For TLS, we will need a certificate and a corresponding key. A self signed certificate 
can be generated using the `openssl` utility. Check the version of openssl you have using the command `openssl version`. If your version of openssl is 1.1.1 or later, the following command should be sufficient:

```bash
openssl req -newkey rsa:2048 -nodes -keyout ws.key -x509 -days 365 -subj "/C=IN/O=AASampleOrg/OU=GoTraining/CN=localhost" -addext "subjectAltName = DNS:localhost" -out ws.crt
```

This command will create a key pair (`-newkey`), save it without requiring a passphrase (`-nodes`) to the file *ws.key*, then generate a self signed certificate (`-x509`) using that key and save it to the file *ws.crt*.

The option `-subj` fills in the Country, Organization and Common Name fields in the certificate. If not provided, openssl would prompt for this information.

The option `-addext` adds a Subject Alternative Name (SAN) to the certificate. This lets browsers perform an additional validation - the hostname used for the request is checked against the certificate. _This option is only available if the version of openssl is 1.1.1 or greater_. 

If you have a version of openssl earlier than 1.1.1, or if you have LibreSSL, those options can be provided in a configuration file. The included file *ssl.cnf* contains the equivalent options. To use that, issue the following command instead:

```bash
openssl req -newkey rsa:2048 -nodes -keyout ws.key -config ssl.conf -x509 -days 365 -out ws.crt
```

Finally, uncomment and comment the relevant lines at the end of the `main` function in *webserver.go*.