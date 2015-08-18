Dashboard for BOSH deployments
==============================

Once you have multiple BOSH, multiple environments, multiple data centers it can quickly become difficult to visualise "what is running? where is it running?"

![deployments](http://cl.ly/image/1d0F153a271D/Deployments.png)

"Is the same version of Cloud Foundry running in all data centers?"

"Is production Logsearch running the same version as staging Logsearch?"

This dashboard makes it much easier to visualize the versions of software running across your BOSH systems.

This can make visualizing a pipeline of deployments easier.

### Uploading latest data

Manually, you could import the latest data from each BOSH:

```
curl -X POST http://bosh-deployments.mycf.com/bosh -d '{"uuid": "the-uuid", "deployments": [{"name":"concourse","releases":[{"name":"concourse","version":"0.59.0"},{"name":"garden-linux","version":"0.284.0"},{"name":"slack-notification-resource","version":"3"},{"name":"cf-haproxy","version":"5"}],"stemcells":[{"name":"bosh-warden-boshlite-ubuntu-trusty-go_agent","version":"2776"}],"cloud_config":"none"}]}'
```

You could programmatically extract the deployments from each BOSH (where the app is running locally on port 3000)

```
deployments=$(curl -ks -u admin:admin https://10.20.30.40:25555/deployments)
curl -X POST http://localhost:3000/bosh -d "{\"uuid\": \"the-uuid\", \"deployments\": $deployments}"
```

Thanks
------

-	https://github.com/jiji262/Bootstrap_Metro_Dashboard was the theme used
