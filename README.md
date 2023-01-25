# Modular Monolith - Boilerplate

This essay has been published at > https://essays.nesovic.dev/posts/modular-monolith-boilerplate/
## From Theory to Practice

It is recommended that readers familiarize themselves with the principles outlined in the article **[Modular Monoliths - Simplified](https://essays.nesovic.dev/posts/modular-monoliths/)** before delving into the practical example provided in this follow-up piece. This article will explore a specific implementation of a modular monolith architecture, utilizing a clear separation of handlers, services, and repository layers. The accompanying [GitHub Repository](https://github.com/kaynetik/modular-monolith-example/) serves as a reference and starting point, providing a boilerplate structure that can be easily adapted to suit the specific needs of your project.

-----

<details><summary>Click to Expand for the Recommended Project Structure</summary>
<p>

#### Example Project Structure for a Modular Monolith

```bash
.
├── Makefile
├── README.md
├── build
│   ├── certs
│   ├── ci
│   │   ├── docker
│   │   │   ├── api
│   │   │   │   └── Dockerfile
│   │   │   └── zapmodule
│   │   │       └── Dockerfile
│   │   └── kube
│   │       ├── api
│   │       │   └── deployment.yaml
│   │       └── zapmodule
│   │           └── deployment.yaml
│   ├── config
│   │   └── redis.conf
│   └── package
│       ├── api
│       └── zapmodule
├── cmd
│   ├── api
│   │   └── main.go
│   └── zapmodule
│       └── main.go
├── docker-compose.yml
├── docs
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   ├── api
│   │   │   ├── handlers
│   │   │   │   └── attach.go
│   │   │   └── users
│   │   │       ├── create-user.go
│   │   │       ├── get-user.go
│   │   │       ├── handlers.go
│   │   │       └── service.go
│   │   ├── errors
│   │   │   └── errors.go
│   │   └── zapmodule
│   │       ├── handlers
│   │       │   └── attach.go
│   │       └── users
│   │           ├── create-user.go
│   │           ├── get-user.go
│   │           ├── handlers.go
│   │           └── service.go
│   └── pkg
│       ├── auth
│       │   └── authorize.go
│       ├── config
│       │   ├── env.go
│       │   └── server.go
│       ├── email
│       │   └── sender.go
│       ├── env
│       │   ├── env.go
│       │   └── vars.go
│       ├── jwt
│       │   ├── jwt.go
│       │   ├── jwt_test.go
│       │   └── suite_setup_test.go
│       ├── middleware
│       │   └── auth.go
│       ├── models
│       │   ├── base.go
│       │   ├── migration.go
│       │   └── user.go
│       ├── server
│       │   └── server.go
│       └── storage
│           ├── migrations
│           │   ├── 000-table-migrations.sql
│           │   └── 001-api-tables.sql
│           ├── postgres
│           │   └── db.go
│           ├── storage.go
│           ├── storage_suite_setup_test.go
│           ├── user.go
│           └── user_test.go
└── tests
    └── test_case.go
```

</p>
</details>

-----

The example presented in this article features two modules, `api` and `zapmodule`, with the latter serving as a clone of the former. The `cmd` package acts as the main entry point, and a clear separation is evident throughout the project structure, including in the `internal/app`, `build/package`, `build/ci/docker`, and `build/ci/kube` packages. The primary benefit of this approach is the ability to share code in the `internal/pkg` package across the entire system without worrying about circular dependencies. The repository layer and models are also located in this package.

It is important to note that there are downsides to this approach, and it is up to the development team to weigh the trade-offs between complexity in service discovery and infrastructure versus the potential for increased entanglement and coupling. To mitigate this, it is crucial to enforce a clear separation between modules and strictly prohibit cross-over imports. This responsibility falls on the shoulders of the engineers and their peers during pull request reviews.

Particular attention should be paid to the precise handling of requests within each module. Each module boasts its own distinct entities, in this example, the entity being `users`. These entities expose specific routes, such as the creation and retrieval of users, located in the file `internal/app/api/users/handlers.go`. All relevant routes for that particular entity are contained within the `Attach(Router, Repo)` function, which is then seamlessly integrated into the appropriate route stack on the server, as demonstrated in `internal/app/api/handlers/attach.go`. It is worth noting that this approach is not a rigid one, and should be tailored to suit the unique needs of your project. While the structure provides a clear separation and organization, it is not rigidly enforced, and deviations from it may be necessary to achieve the desired outcome. In other words, **it is a guide, not a rule**.

-----

🎨 Crafting software is an art, and our canvas is **simplicity**. We believe in creating solutions that are not only elegant in design but also robust and tested to withstand the test of time. Our approach is to provide a solution that meets stakeholders' requirements and ensures long-term maintainability and scalability. Our ultimate aim is to deliver efficient, effective, and adaptable software to the ever-evolving needs of businesses without succumbing to the allure of unnecessary complexity.

If that is what you seek, then contact us at [contact@decantera.dev](mailto:contact@decantera.dev) or via our site [decantera.dev](https://decantera.dev). 🚀