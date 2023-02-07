# Storm

Storm is a template for a Go backend project.

## Steps to initiate a project

* Create go.mod
* Create cmd
* cmd should contain entry points of different executable units of the project. e.g. http server, grpc server, graphql server, seeder cli, and others.
* Create model
* model should contain domain data structures, validators on these structures, storage layer interfaces, service layer interfaces, and constants.
* Create internal
* internal should contain defined error package.
* Create pkg
* pkg should contain any general code which doesn't relate directly to the project's business. e.g. libraries, general types, ...