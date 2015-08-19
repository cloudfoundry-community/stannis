Dashboard for BOSH deployments
==============================

Once you have multiple BOSH, multiple environments, multiple data centers it can quickly become difficult to visualise "what is running? where is it running?"

![deployments](http://cl.ly/image/1d0F153a271D/Deployments.png)

"Is the same version of Cloud Foundry running in all data centers?"

"Is production Logsearch running the same version as staging Logsearch?"

This dashboard makes it much easier to visualize the versions of software running across your BOSH systems.

This can make visualizing a pipeline of deployments easier.

Running dashboard
-----------------

Create a YAML configuration file to describe the pipelines of deployments. See `config/config.example.yml` for the schema and examples.

```
go run main.go -pipelines config.yml
```

Upload deployment data from BOSHes
----------------------------------

### Manually upload data

Manually, you could import the latest data from each BOSH:

```
curl -X POST http://bosh-deployments.mycf.com/bosh -d '{"uuid": "the-uuid", "deployments": [{"name":"concourse","releases":[{"name":"concourse","version":"0.59.0"},{"name":"garden-linux","version":"0.284.0"},{"name":"slack-notification-resource","version":"3"},{"name":"cf-haproxy","version":"5"}],"stemcells":[{"name":"bosh-warden-boshlite-ubuntu-trusty-go_agent","version":"2776"}],"cloud_config":"none"}]}'
```

You could programmatically extract the deployments from each BOSH (where the app is running locally on port 3000)

```
deployments=$(curl -ks -u admin:admin https://10.20.30.40:25555/deployments)
curl -X POST http://localhost:3000/upload -d "{\"uuid\": \"the-uuid\", \"name\": \"bosh-lite\", \"deployments\": $deployments}"
```

If you are running the server with the example config (`go run main.go -pipelines config/config.example.yml`) then you can upload the example/fixtures data:

```
./bin/upload_fixtures
```

Thanks
------

-	Many staff at [Stark & Wayne](https://starkandwayne.com) who've been working on pipelining BOSH deployments for our clients; which led to the problem being solved "what's actually running everywhere?"
-	https://github.com/jiji262/Bootstrap_Metro_Dashboard was the theme used
