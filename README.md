# Blog_post_server
====================================================================


## Authors

- [@dev-hack95](https://www.github.com/dev-hack95)

## Project Status
- Complete

## Table of Contents

  - [Tech Stack](#tech-stack)
  - [Flowchart](#flowchart)

## Tech Stack
  - Golang, Docker, Postgresql, Microservices

## Flowchart

```mermaid
flowchart LR;
    A(PKG) -----> B(config);
    A -----> C(controllers);
    A -----> D(Models);
    A -----> E(routes);
    A -----> F(utils);
    B -----> G(app.go);
    C -----> H(server.go);
    D -----> I(models.go);
    E -----> J(routes.go);
    F -----> K(utils.go);
```