# REST API
## Update Proxy
PROXY_GROUP is selector name.
curl -v -X PUT -H "Content-Type: application/json" -d '{"name":"PROXY_NAME" }' "http://localhost:9999/proxies/PROXY_GROUP"