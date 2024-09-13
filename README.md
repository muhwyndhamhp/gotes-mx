# What is this?
This is a teeny tiny template repo for my own personal use. This setup meant for a fullstack Go project with templating support with HTMX. 

This repo's goal is to provide opinionated, but tiny template that ready to be used to develop go web apps. I would only provide what's important (for me) and omit what's not. 

## What this repo can handle out of the box
- Hot Reload via [Air](https://github.com/cosmtrek/air) just `make-run`.
- Quick Local dev setup including DB via `make setup-local`.
- ORM ready using [GORM](https://gorm.io/).
- Routing ready using [Labstack Echo](https://github.com/labstack/echo).
- TempL is used to templating purpose, see: [TempL Getting Started](https://templ.guide).
- [TailwindCSS](https://tailwindcss.com/) Ready! 
- Exposed directory for resources in the `public` directory (ex. `http://localhost:4040/public/favicon.ico`).
- Auth via [Zitadel](https://zitadel.com/docs/examples/login/go).

## What this repo will never handle
- Deployment beyond simple Dockerfile & Docker Compose
- Testing

## Prerequisites
- [Air](https://github.com/cosmtrek/air)
- [Docker](https://docs.docker.com/get-started/)

## Getting Started
- Create `.env` file (see `env.sample`).
- Run `make local-setup`. It will setup Postgres DB & Zitadel Self-hosted.
- Open up Zitadel (default is in `http://localhost:5756`), login with default login and update the default password.
```
username: zitadel-admin@zitadel.localhost
password: Password1!
```
- Follow create frontend project from [this tutorial](https://zitadel.com/docs/examples/login/go) until you receive the client id. Follow it to a T except when asked for callback URL. Change the localhost domain to `localhost:4040` instead of `localhost:8089`.
- Put the client id into the `.env` file on `ZITADEL_CLIENT_ID=` section
- (Optional) Run `make seed` to migrate and seed the example Todo App.
- Run `make run` and enjoy :)


## Todo List
- Make it easily deployable to Fly.io
- Add ESBuild support for client side JS
- Helpers for Server-Side HTMX Control Flow (Redirects, Out-of-Band Swap, etc.)
- Cache Control helpers
