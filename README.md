<h1 align="center">
    headerpwn
  <br>
</h1>

<h4 align="center">A fuzzer for analyzing how servers respond to different HTTP headers.</h4>


<p align="center">
  <a href="#install">🏗️ Install</a>
  <a href="#usage">⛏️ Usage</a>
  <br>
</p>


![headerpwn](https://github.com/devanshbatham/headerpwn/blob/main/static/headerpwn.png?raw=true)

# Install
To install `headerpwn``, run the following command:

```
go install github.com/devanshbatham/headerpwn@latest
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

## Credits
The `headers.txt` fileis compiled from various sources, including the Seclists project (https://github.com/danielmiessler/SecLists). These headers are used for testing purposes and provide a variety of scenarios for analyzing how servers respond to different headers.

