This was originally forked from the [Sample Function: Go "Random"](https://github.com/digitalocean/sample-functions-golang-random) repo. It now has some tools for testing the serverless go functions locally.

# Deployment

- Fork this repo
- create devcontainer (you can use github codespaces for this!)
- create free neondb database and get the connection string (or some other way to get a postgres url)
- Put in .env: `DB_URL=your-connection-string`
- `doctl serverless deploy .`
- manually create the schedule trigger

## Create a free database with neon

To get started, we need to create a new Postgres instance for us to connect to.
For this tutorial, we will be using [Neon Postgres](https://neon.tech/) as they
provide free, managed Postgres instances. If you like to host your database
somewhere else, you can do that too.

1. Visit https://neon.tech/ and click **Sign up** to sign up with an email,
   Github, Google, or partner account. After signing up, you are directed to the
   Neon Console to create your first project.
2. Enter a name for your project, select a Postgres version, provide a database
   name, and select a region. Generally, you'll want to select the region
   closest to your application. When you're finished, click **Create project**.
3. You are presented with the connection string for your new project, which you
   can use to connect to your database. Save the connection string, which looks
   something like this:

   ```sh
   postgres://alex:AbC123dEf@ep-cool-darkness-123456.us-east-2.aws.neon.tech/dbname?sslmode=require
   ```

   You will need the connection string in the next step.

# Development

A [devcontainer.json](./.devcontainer/devcontainer.json) defines the expected environment.

- **Run locally** - from `local/`, do `go run .`
- **Deploy** - from project root, do `doctl serverless deploy .`
    - **NOTE:** the trigger doesn't get created automatically because of a [bug](https://github.com/digitalocean/doctl/issues/1474), you need to create the `gong` trigger in the web console. (maybe this won't be true if deployed with the DO template?)
- **View runs** - `doctl serverless activations list`
- **View logs** - `doctl serverless activations logs` (gets only most recent run)

## Using the Function

```
doctl serverless functions invoke discord-social-media-bot/main
```

Use this command to retrieve the URL for your function and use it as an API.
```
doctl sls fn get discord-social-media-bot/main --url
```

You can use that API directly in your browser, with `curl` or with an API platform such as Postman.
Parameters may be passed as query parameters, or as JSON body. Here is an example using `curl`.

```
curl `doctl sls fn get discord-social-media-bot/main --url`
```

### Learn More

You can learn more about Functions by reading the [Functions Documentation](https://docs.digitalocean.com/products/functions). 
