/**
 * Copyright 2016 Confluent Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Example function-based high-level Apache Kafka consumer
package main

// consumer_example implements a consumer using the non-channel Poll() API
// to retrieve messages and events.

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        os.Getenv("BOOTSTRAP_SERVERS"),
		"security.protocol":        "SASL_SSL",
		"sasl.username":            os.Getenv("SASL_USERNAME"),
		"sasl.password":            os.Getenv("SASL_PASSWORD"),
		"sasl.mechanism":           "PLAIN",
		"broker.address.family":    "v4",
		"group.id":                 "local-test-consumer",
		"session.timeout.ms":       6000,
		"auto.offset.reset":        "earliest",
		"go.events.channel.enable": true,
		"enable.auto.offset.store": false,
		"statistics.interval.ms":   5000,
		"enable.auto.commit":       false,
		"max.poll.interval.ms":     10000,
	})

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer 0.3 %v\n", c)

	topics := [...]string{"go-client-test-topic"}
	err = c.SubscribeTopics(topics[:], nil)
	var msgValue string

	var running = true
	for running {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			running = false
		default:
			ev := <-c.Events()
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:

				msgValue = string(e.Value)
				fmt.Printf("Value: %v\n", msgValue)
				_, err := c.StoreMessage(e)
				if err != nil {
					_, _ = fmt.Fprintf(os.Stderr, "%% Error storing offset after message %s:\n", e.TopicPartition)
				}

			case kafka.OffsetsCommitted:
				fmt.Println("offsets committed")

			default:
				fmt.Println("Ignored")
			}
		}
	}

	fmt.Println("Closing consumer")
	_ = c.Close()
}
