This was originally forked from the [Sample Function: Go "Random"]() repo. It now has some tools for testing the serverless go functions locally.

A [devcontainer.json](./.devcontainer/devcontainer.json) defines the expected environment.

- **Run locally** - from `local/`, do `go run .`
- **Deploy** - from project root, do `doctl serverless deploy .`
- **View runs** - `doctl serverless activations list`
- **View logs** - `doctl serverless activations logs` (gets only most recent run)

## Using the Function

```
doctl serverless functions invoke bot/main
```

Use this command to retrieve the URL for your function and use it as an API.
```
doctl sls fn get bot/main --url
```

You can use that API directly in your browser, with `curl` or with an API platform such as Postman.
Parameters may be passed as query parameters, or as JSON body. Here is an example using `curl`.

```
curl `doctl sls fn get bot/main --url`
```

### Learn More

You can learn more about Functions by reading the [Functions Documentation](https://docs.digitalocean.com/products/functions). 
