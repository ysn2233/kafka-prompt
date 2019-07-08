# Kafka-prompt
An intertactive auto-complete kafka client usign [c-bata/go-prompt](https://github.com/c-bata/go-prompt)

![demo](https://github.com/ysn2233/kafka-prompt/blob/master/demo.gif)

## Prerequisite
This command-line tool requires kafka console client installed on the filesystem. (In PATH or using absolutely path to call client executable) 

## Installation
### Build from source
```
cd kafka-prompt
make
```
The execuatable binary is located as ./bin/kaprompt

## Feature
The auto-completion of kafka client command is predefined so make sure those client executable(eg. kafka-topic, kafka-console-consumer) have alreadly in your PATH. Kafka-prompt also supports any command ends with ".sh".

### Memorize servers
If you have run a command with "--bootstrap-server/--broker-list/--zookeeper", the prompt will cache these server connected information. Next command without those options can auto add them to connect.

### Support-autocompletion-commands
* kafka-replica-verification
* kafka-log-dirs
* kafka-acls
* kafka-delete-records
* kafka-preferred-replica-election
* kafka-console-producer
* kafka-dump-log
* zookeeper-shell
* kafka-topics
* kafka-console-consumer
* kafka-consumer-groups
* kafka-reassign-partitions
* kafka-broker-api-versions
* kafka-configs

### Built-in commands
* **exit/quit** Exit the shell
* **cluster boostrap/server/zookeeper** Show current server connected