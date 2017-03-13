# ksd
### Env
`KSD_PORT` defaults `8080`
`KSD_TOKEN`

### Deploy
`/deploy?token={KSD_TOKEN}`
```JSON
{
    "namespace": "my-namespace",
    "deployment": "my-app",
    "containers": [
        {"name": "my-container", "image": "nginx:latest"}
    ]
}
```