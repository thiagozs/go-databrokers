# Data Broker with Golang

## Stack of techology

* Golang 1.14.1
* Docker
* Message Broker (Mosquitto, EMQX, Kafka)
* Producer/Consumer clients
* Documentation

## Inspiration

Complexity test of running multiple brokers, code validation, technical feasibility, among other factors.

## Work flow design

* Publish + Subscribe
* Processing
* Store

## Proof of Concept

All consumers and producers are in their respective directories. If you would like more information on how to leverage the services you can consult the `DOC` directory, there is a **README.md** there that will help you with these jobs.

poc folder structure

```sh
├── emqx
│   ├── consumer
│   │   └── main.go
│   └── producer
│       └── main.go
├── kafka
│   ├── consumer
│   │   └── main.go
│   ├── docker-compose.yml
│   └── producer
│       └── main.go
└── mqtt
    ├── consumer
    │   └── main.go
    └── producer
        └── main.go
```

## Versioning and license

We use SemVer for versioning. You can see the versions available by checking the tags on this repository.

For more details about our license model, please take a look at the [LICENSE](LICENCE) file

2020, thiagozs
