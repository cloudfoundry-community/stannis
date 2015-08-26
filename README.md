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

Create a YAML configuration file to describe the pipelines of deployments. See `config/webserver.config.example.yml` for the schema and examples.

```
go run main.go webserver --config config/webserver.config.example.yml
```

Or if `stannis` is installed:

```
go get github.com/cloudfoundry-community/stannis
stannis webserver --config config/webserver.config.example.yml
```

Deploying Stannis to Cloud Foundry
----------------------------------

This repository can be pushed to any Cloud Foundry. Stark & Wayne runs all its Stannis dashboards (one per client, one for itself) on [Pivotal Web Services](http://run.pivotal.io/).

First, create `config.yml` in the root folder so that it is uploaded with the entire Stannis codebase. See `config/webserver.config.example.yml` for an example.

To deploy to any Cloud Foundry:

```
cf push stannis
```

Change `Procfile` to pass in any additional flags.

Use Stannis to upload BOSH data
-------------------------------

This Stannis CLI is also an agent that fetches data from a BOSH and uploads it to a central Stannis dashboard.

```
go get github.com/cloudfoundry-community/stannis
stannis agent --config agent-config.yml
```

See `config/agent.config.example.yml` for an example of this configuration file.

Run Stannis agent via BOSH
--------------------------

Since you already have BOSH, you can run one Stannis agent per BOSH using the [stannis-boshrelease](https://bosh.io/releases/github.com/cloudfoundry-community/stannis-boshrelease)

See the README for upload and deployment instructions for your infrastructure.

Run Stannis agent within the BOSH
---------------------------------

This seems like a cheap idea - run the `stannis agent` inside the BOSH VM itself.

This could be easy using `bosh-init` where it is easy to merge multiple releases - such as `bosh` and `stannis` into one VM. Something to be investigated for sure.

Thanks
------

-	Many staff at [Stark & Wayne](https://starkandwayne.com) who've been working on pipelining BOSH deployments for our clients; which led to the problem being solved "what's actually running everywhere?"
-	GE's Predix team at San Ramon who have huge BOSH universes and tested some edge cases of Golang standard libraries
-	https://github.com/jiji262/Bootstrap_Metro_Dashboard was the theme used
