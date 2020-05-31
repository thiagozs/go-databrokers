# Study of DATA Broker

## Options for test

* EMQX (poc done! ps: mqtt protocol)
* Kafka (in progress...)
* Mosquitto (poc done! ps: mqtt protocol)

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

Pros:

Cons:

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

Pros:

Cons:

## Versioning and license

We use SemVer for versioning. You can see the versions available by checking the tags on this repository.

For more details about our license model, please take a look at the [LICENSE](../LICENSE) file

---

2020, thiagozs
