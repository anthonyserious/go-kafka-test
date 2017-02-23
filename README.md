go-kafka-test
====

*go-kafka-test* is something I'm building to learn Go.  Here I use Kafka as a bus, and I want to build some contrived microservices, with the ability to just deploy everything with docker-compose.  I use Spotify's `spotify/kafka` image for its simplicity (bundled zookeeper), and `sheepkiller/kafka-manager` so I can manage Kafka from a web UI.

Effectively, this application will periodically read station data from the public [NYC CitiBike API](https://feeds.citibikenyc.com/stations/stations.json). A producer will pull this data and publish one message per station onto a Kafka topic.

As for consumers...who knows?  One can dump the data into a database, another can run some analytics.  For now, it'll just dump formatted messages to stdout.
