# Parago

Design, generate, and share application definition across multiple providers.

![Parago](docs/images/app-definition.png?raw=true "Parago")

Parago helps teams define applications by using a drag and drop interface and generate the application definitions to be deployed on Kubernetes using different providers.

While application definitions can be kept private, users can also save and share their definitions with users within their organizations or publicly.

## Architecture

![Architecture](docs/images/parago-arch.png?raw=true "Architecture")

- Definitions are saved using MongoDB
- Parago uses Shipa for authentication - You can create a [free account here](https://apps.shipa.cloud)
- Parago initially used Shipa for deployment, so files generated for the different providers use Shipa as the deployment engine
- Current providers include: Crossplane, Ansible, CloudFormation, Terraform, GitLab, GitHub Actions, and Pulumi
- Roadmap: Add Helm Chart generation (this does not assume Shipa will be the deployment engine but rather a Helm chart to deploy your app)

## Development

To build and run everything 

```shell
docker-compose build
docker-compose up
```

To working with backend you can start frontend using docker-compose and doing manual build and run for backend part and vice versa.

To start frontend from a container
```shell
    docker-compose up frontend
```

## Documentation

More details can be found in [docs](docs/index.md) section.

## Contributing

We welcome contributions to Parago in the form of Pull Requests or submitting Issues. 

More information available [here](CONTRIBUTING.md)
