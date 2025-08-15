# Signal Gateway

Signal Gateway is a small service that allows sending [Signal](https://signal.org/) messages via HTTP POST requests.
Messages are sent using [signal-cli](https://github.com/AsamK/signal-cli), accessed via JSON-RPC over HTTP.

Messages can be sent either as JSON-encoded data or as URL-encoded form data.

The service handles the conversion of POST requests into JSON-RPC calls. Furthermore, it is possible to restrict the allowed recipients.

## Configuration (`config.yaml`)

```yaml
port: 80
endpoint: /message/send
signalCliEndpoint: http://signal-cli/api/v1/rpc
account: +491701234567
allowedRecipients:
- +491712345678
- +491723456789
```

| Parameter             | Description                                                        |
|-----------------------|--------------------------------------------------------------------|
| `port`                | The port on which the server listens.                              |
| `endpoint`            | The HTTP endpoint for sending messages.                            |
| `formEndpoint`        | An optional HTTP endpoint for a test form.                         |
| `signalCliEndpoint`   | The URL of the signal-cli JSON-RPC server.                         |
| `account`             | The Signal account for sending messages.                           |
| `allowedRecipients`   | A list of allowed recipients. If empty all recipients are allowed. |

## Example HTTP Requests

### JSON

```http
POST /message/send HTTP/1.1
Host: signal-gateway
Content-Type: application/json
Content-Length: 54

{"recipient":"+491712345678","message":"Hello World!"}
```

### XML

```http
POST /message/send HTTP/1.1
Host: signal-gateway
Content-Type: application/xml
Content-Length: 107

<?xml version="1.0"?><message><recipient>+491712345678</recipient><message>Hello World!</message></message>
```

### Form Data

```http
POST /message/send HTTP/1.1
Host: signal-gateway
Content-Length: 48
Content-Type: application/x-www-form-urlencoded

recipient=%2B491723456789&message=Hello+World%21
```

### Multipart Form Data

```http
POST /message/send HTTP/1.1
Host: signal-gateway
Content-Length: 157
Content-Type: multipart/form-data; boundary=data

--data
Content-Disposition: form-data; name="recipient"

+491723456789
--data
Content-Disposition: form-data; name="message"

Hello World!
--data--
```

## Example curl Commands

```bash
curl -X POST -H 'Content-Type: application/json' -d '{"recipient":"+491712345678","message":"Hello World!"}' signal-gateway/message/send
```

```bash
curl -X POST -H 'Content-Type: application/xml' -d '<?xml version="1.0"?><message><recipient>+491712345678</recipient><message>Hello World!</message></message>' signal-gateway/message/send
```

```bash
curl --data-urlencode 'recipient=+491723456789' --data-urlencode 'message=Hello World!' signal-gateway/message/send
```

```bash
curl -F 'recipient=+491723456789' -F 'message=Hello World!' signal-gateway/message/send
```
