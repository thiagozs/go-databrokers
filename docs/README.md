# Study of DATA Broker

## Options for test

* EMQX ✔
* Mosquitto ✔
* Kafka ✔

## EMQX

Repo: [https://github.com/emqx/emqx](https://github.com/emqx/emqx)

Runing with Docker image:

```sh
sudo docker run -d --name emqx -p 1883:1883 -p 8083:8083 -p 8883:8883 -p 8084:8084 -p 18083:18083 emqx/emqx
```

Start with repo binary compiled:

```sh
# Clone repo
git clone -b v4.0.0 https://github.com/emqx/emqx-rel.git

# Make command
cd emqx-rel && make

# Console
cd _build/emqx/rel/emqx && ./bin/emqx console

# Start emqx
./bin/emqx start

# Check Status
./bin/emqx_ctl status

# Stop emqx
./bin/emqx stop
```

## Mosquitto

Repo : [https://github.com/eclipse/mosquitto](https://github.com/eclipse/mosquitto)

Runing with Docker image:

```sh
sudo docker run -d --name mosquitto -p 1883:1883 -p 9001:9001 eclipse-mosquitto
```

If you have installed a binary package the broker should have been started
automatically. If not, it can be started with a basic configuration:

Then use `mosquitto_sub` to subscribe to a topic:

```sh
mosquitto_sub -t 'test/topic' -v
```

And to publish a message:

```sh
mosquitto_pub -t 'test/topic' -m 'hello world'
```

## Kafka

Repo : [https://github.com/apache/kafka](https://github.com/apache/kafka)

Start with docker composer. Just run `sudo docker-compose up` or deattached mode `sudo docker-compose up -d`

```yml
---
version: '2'
services:
  zookeeper:
    image: confluentinc/cp-zookeeper:latest
    network_mode: host
    environment:
      ZOOKEEPER_CLIENT_PORT: 32181
      ZOOKEEPER_TICK_TIME: 2000
    extra_hosts:
      - "moby:127.0.0.1"

  kafka:
    image: confluentinc/cp-kafka:latest
    network_mode: host
    depends_on:
      - zookeeper
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: localhost:32181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:29092
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
    extra_hosts:
      - "moby:127.0.0.1"
```

## Versioning and license

We use SemVer for versioning. You can see the versions available by checking the tags on this repository.

For more details about our license model, please take a look at the [LICENSE](../LICENSE) file

---

2020, thiagozs
