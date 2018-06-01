# kafka_sarama_vs_confluent_write
A benchmark between Kafka drivers (Sarama and Confluent)


# Install
```
go get github.com/Shopify/sarama
go get github.com/confluentinc/confluent-kafka-go/kafka
```

# Run 
You should change bootstrap Kafka servers
```
go run sarama.go
go run confluent.go
```

# Example 10000 inserts
```
$ go run confluent.go
2018-06-01 09:45:08.139198119 -0300 -03 m=+0.022129414
2018-06-01 09:45:10.041213364 -0300 -03 m=+1.924202659
1.902073245s

$ go run sarama.go
2018-06-01 09:44:51.249858291 -0300 -03 m=+0.042064779
2018-06-01 09:45:03.954180705 -0300 -03 m=+12.746910193
12.704845414s
```


# Conclusion
Confluent is much faster than Sarama
