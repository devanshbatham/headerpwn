<h1 align="center">
    headerpwn
  <br>
</h1>

<h4 align="center">A fuzzer for analyzing how servers respond to different HTTP headers.</h4>


<p align="center">
  <a href="#install">🏗️ Install</a>
  <a href="#usage">⛏️ Usage</a>
  <a href="#proxying-requests-through-burp-suite">📡 Proxying HTTP Requests</a>
  <br>
</p>


![headerpwn](https://github.com/devanshbatham/headerpwn/blob/main/static/banner.png?raw=true)

# Install
To install `headerpwn`, run the following command:

```
go install github.com/devanshbatham/headerpwn@v0.0.3
```

# Usage
headerpwn allows you to test various headers on a target URL and analyze the responses. Here's how to use the tool:

1. Provide the target URL using the `-url` flag.
2. Create a file containing the headers you want to test, one header per line. Use the `-headers` flag to specify the path to this file.

Example usage:
```sh
headerpwn -url https://example.com -headers my_headers.txt
```

- Format of `my_headers.txt` should be like below:

```sh
Proxy-Authenticate: foobar
Proxy-Authentication-Required: foobar
Proxy-Authorization: foobar
Proxy-Connection: foobar
Proxy-Host: foobar
Proxy-Http: foobar
```

## Proxying requests through Burp Suite: 

Follow following steps to proxy requests through Burp Suite: 


- Export Burp's Certificate:

    - In Burp Suite, go to the "Proxy" tab.
    - Under the "Proxy Listeners" section, select the listener that is configured for `127.0.0.1:8080`
    - Click on the "Import/ Export CA Certificate" button.
    - In the certificate window, click "Export Certificate" and save the certificate file (e.g., burp.der).


- Install Burp's Certificate:

    - Install the exported certificate as a trusted certificate on your system. How you do this depends on your operating system.
    - On Windows, you can double-click the .cer file and follow the prompts to install it in the "Trusted Root Certification Authorities" store.
    - On macOS, you can double-click the .cer file and add it to the "Keychain Access" application in the "System" keychain.
    - On Linux, you might need to copy the certificate to a trusted certificate location and configure your system to trust it.


You should be all set: 


```sh
headerpwn -url https://example.com -headers my_headers.txt -proxy 127.0.0.1:8080
```



![proxy](https://github.com/devanshbatham/headerpwn/blob/main/static/proxy-cli.png?raw=true)


![proxy-burp](https://github.com/devanshbatham/headerpwn/blob/main/static/proxy-burp.png?raw=true)


## Credits
The `headers.txt` file is compiled from various sources, including the [Seclists project](https://github.com/danielmiessler/SecLists). These headers are used for testing purposes and provide a variety of scenarios for analyzing how servers respond to different headers.

