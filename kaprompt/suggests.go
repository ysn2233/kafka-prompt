package kaprompt

import "github.com/c-bata/go-prompt"

var (
	commandSuggest = []prompt.Suggest{
		{Text: "kafka-delete-records", Description: "Delete records of the given partitions by offset."},
		{Text: "kafka-console-producer", Description: "Read data from standard input and publish it to Kafka."},
		{Text: "kafka-preferred-replica-election", Description: "Cause leadership for each partition to be transferred back to the 'preferred replica'."},
		{Text: "kafka-replica-verification", Description: "Validate that all replicas for a set of topics have the same data."},
		{Text: "kafka-broker-api-versions", Description: "Retrieve broker version information."},
		{Text: "kafka-dump-log", Description: "Parse a log file and dump its contents to the console."},
		{Text: "kafka-consumer-groups", Description: "List all consumer groups, describe a consumer group, delete consumer group info, or reset consumer group offsets."},
		{Text: "kafka-configs", Description: "Add/Remove entity config for a topic, client, user or broker"},
		{Text: "kafka-console-consumer", Description: "The console consumer is a tool that reads data from Kafka and outputs it to standard output"},
		{Text: "kafka-reassign-partitions", Description: "Moves topic partitions between replicas."},
		{Text: "kafka-log-dirs", Description: "Find log directory usage on specified brokers."},
		{Text: "kafka-topics", Description: "Create, delete, describe, or change a topic"},
		{Text: "zookeeper-shell", Description: "Access zookeeper shell."},
		{Text: "kafka-acls", Description: "Manage access control lists."},
	}
	optionSuggest = map[string][]prompt.Suggest{
		"kafka-preferred-replica-election": {
			{Text: "--admin.config", Description: "Property file containing configs to be passed to the admin client."},
			{Text: "--bootstrap-server", Description: "The Kafka broker(s) to connect to."},
			{Text: "--help", Description: "Print usage information."},
			{Text: "--path-to-json-file", Description: "The JSON file with the list of partitions to perform preferred replica leader election on."},
			{Text: "--zookeeper", Description: "DEPRECATED. Zookeeper connection string."},
		},
		"kafka-replica-verification": {
			{Text: "--max-wait-ms", Description: "The max amount of time each fetch request waits. Default: 1000."},
			{Text: "--report-interval-ms", Description: "The reporting interval. Default: 30000."},
			{Text: "--time", Description: "Timestamp for getting the initial offsets. Default: -1."},
			{Text: "--topic-white-list", Description: "Whitelist of topics to verify replica consistency using Java regex. Default: '.*'."},
			{Text: "--broker-list", Description: "The Kafka broker(s) to connect to. NOTE: This'll overwrite the selected cluster value."},
			{Text: "--fetch-size", Description: "The fetch size of each request in bytes. Default: 1048576."},
		},
		"kafka-broker-api-versions": {
			{Text: "--command-config", Description: "Property file containing configs to be passed to the admin client."},
			{Text: "--bootstrap-server", Description: "The Kafka broker(s) to connect to. NOTE: This'll overwrite the selected cluster value."},
		},
		"kafka-consumer-groups": {
			{Text: "--execute", Description: "Execute operation. Supported for '--reset-offsets'."},
			{Text: "--state", Description: "Describe the group state."},
			{Text: "--export", Description: "Export operation execution to a CSV file. Supported for '--reset-offsets'."},
			{Text: "--group", Description: "The consumer group to act on."},
			{Text: "--list", Description: "List all consumer groups."},
			{Text: "--shift-by", Description: "Reset offsets shifting current offset by 'n' where 'n' is positive or negative."},
			{Text: "--by-duration", Description: "Reset offsets to offset by duration from current timestamp. Format: 'PnDTnHnMnS'."},
			{Text: "--command-config", Description: "Property file containing configs to be passed to admin client and consumer."},
			{Text: "--describe", Description: "Describe consumer group and list offset lag for the consumer group."},
			{Text: "--dry-run", Description: "Show results without executing changes. Supported for '--reset-offsets'."},
			{Text: "--to-current", Description: "Reset offsets to current offset."},
			{Text: "--to-earliest", Description: "Reset offsets to earliest offset."},
			{Text: "--to-latest", Description: "Reset offsets to latest offset."},
			{Text: "--to-offset", Description: "Reset offsets to a specific offset."},
			{Text: "--topic", Description: "The topic whose consumer group information should be acted upon."},
			{Text: "--verbose", Description: "Provide additional information, if any, when describing the group."},
			{Text: "--bootstrap-server", Description: "The Kafka broker(s) to connect to. NOTE: This'll overwrite the selected cluster value."},
			{Text: "--delete", Description: "Delete topic partition offsets and ownership info over the entire consumer group."},
			{Text: "--timeout", Description: "The timeout that can be set for some use cases."},
			{Text: "--to-datetime", Description: "Reset offsets to offset from datetime. Format: 'YYYY-MM-DDTHH:mm:SS.sss'."},
			{Text: "--reset-offsets", Description: "Reset offsets of consumer group. Supports one group at a time."},
			{Text: "--all-topics", Description: "Consider all topics assigned to a group in the 'reset-offsets' process."},
			{Text: "--from-file", Description: "Reset offsets to values defined in CSV file."},
			{Text: "--members", Description: "Describe members of the given consumer group."},
			{Text: "--offsets", Description: "Describe group and list all topic partitions in the group along with their offset lag."},
		},
		"kafka-configs": {
			{Text: "--bootstrap-server", Description: "The Kafka broker(s) to connect to. NOTE: This'll overwrite the selected cluster value."},
			{Text: "--command-config", Description: "Property file containing configs to be passed to the admin client."},
			{Text: "--describe", Description: "List configs for the given entity."},
			{Text: "--entity-default", Description: "Default entity name for clients, users, and brokers."},
			{Text: "--entity-name", Description: "Name of the entity (topic name, client id, user principal name, broker id)."},
			{Text: "--entity-type", Description: "Type of the entity (topic, client, user, broker)."},
			{Text: "--help", Description: "Print usage information."},
			{Text: "--add-config", Description: "Key-value pairs of configs to add. Format: key=value."},
			{Text: "--delete-config", Description: "Key-value pairs of configs to remove. Format: key=value."},
			{Text: "--force", Description: "Assume yes for all queries and do not prompt."},
			{Text: "--zookeeper", Description: "Zookeeper connection string. NOTE: This'll overwrite the selected cluster value."},
			{Text: "--alter", Description: "Alter the configuration for the entity."},
		},
		"kafka-console-consumer": {
			{Text: "--consumer.config", Description: "Consumer config properties file. Note that --consumer-property takes precedence over this config."},
			{Text: "--formatter", Description: "The name of a class to use for formatting kafka messages for display."},
			{Text: "--isolation-level", Description: "Set to read_committed in order to filter out transactional messages which are not committed."},
			{Text: "--key-deserializer", Description: "Deserializer for the key."},
			{Text: "--topic", Description: "The topic to consume from."},
			{Text: "--value-deserializer", Description: "Deserializer for the value."},
			{Text: "--whitelist", Description: "Whitelist of topics to include for consumption."},
			{Text: "--bootstrap-server", Description: "The Kafka broker(s) to connect to. NOTE: This'll overwrite the selected cluster value."},
			{Text: "--consumer-property", Description: "A way to pass user-defined properties in the form key=value to the consumer."},
			{Text: "--group", Description: "The consumer group id of the consumer."},
			{Text: "--max-messages", Description: "The maximum number of messages to consume before exiting. Defaults to continual consumption."},
			{Text: "--from-beginning", Description: "Start with the earliest message present in the log rather than the latest message."},
			{Text: "--partition", Description: "The partition to consume from. Starts from the end unless --offset is set."},
			{Text: "--property", Description: "The properties to initialize the message formatter."},
			{Text: "--skip-message-on-error", Description: "If there is an error when processing a message, skip it rather than halting."},
			{Text: "--timeout-ms", Description: "If specified, exit if no message is available for consumption for the specified interval."},
			{Text: "--enable-systest-events", Description: "Log lifecycle events of the consumer in addition to logging consumed messages."},
			{Text: "--offset", Description: "The offset id to consume from. A non-negative integer, 'earliest', or 'latest'."},
		},
		"kafka-reassign-partitions": {
			{Text: "--throttle", Description: "The movement of partitions between brokers will be throttled by this value. Default: -1."},
			{Text: "--zookeeper", Description: "Zookeeper connection string. WARNING: This'll overwrite the selected cluster value."},
			{Text: "--broker-list", Description: "The list of brokers used in reassignment in the form '0,1,2'. Default: all."},
			{Text: "--disable-rack-aware", Description: "Disable rack aware replica assignment."},
			{Text: "--replica-alter-logs-dirs-throttle", Description: "The movement of replicas between log directories will be throttled by this value. Default: -1."},
			{Text: "--generate", Description: "Generate a candidate partition reassignment configuration. This does not execute it."},
			{Text: "--reassignment-json-file", Description: "The JSON file with the partition reassignment configuration."},
			{Text: "--bootstrap-server", Description: "The Kafka broker(s) to connect to. NOTE: This'll overwrite the selected cluster value."},
			{Text: "--command-config", Description: "Property file containing configs to be passed to the admin client."},
			{Text: "--execute", Description: "Start the reassignment as specified by the '--reassignment-json-file' options."},
			{Text: "--help", Description: "Print usage information."},
			{Text: "--timeout", Description: "The maximum time in ms allowed to wait for execution to start. Default: 10000."},
			{Text: "--topics-to-move-json-file", Description: "Generate a configuration to move partitions of the specified topics to the given list of brokers."},
			{Text: "--verify", Description: "Verify if the reassignment has completed."},
		},
		"kafka-console-producer": {
			{Text: "--sync", Description: "If set message send requests to the brokers are synchronously, one at a time as they arrive."},
			{Text: "--topic", Description: "The topic to produce messages to."},
			{Text: "--max-block-ms", Description: "The max time that the producer will block for during a send request."},
			{Text: "--message-send-max-retries", Description: "The number of retires before the producer will give up and drop a message."},
			{Text: "--producer.config", Description: "Producer config properties file. Note that --producer-property takes precedence over this config."},
			{Text: "--property", Description: "A mechanism to pass user-defined properties in the form key=value to the message reader."},
			{Text: "--request-required-acks", Description: "The required acks of the producer request."},
			{Text: "--socket-buffer-size", Description: "The size of the tcp RECV size."},
			{Text: "--timeout", Description: "In asynchronous mode, the maximum amount of time a message will queue awaiting sufficient batch size."},
			{Text: "--compression-codec", Description: "The compression codec. Default: gzip."},
			{Text: "--max-partition-memory-bytes", Description: "The buffer size allocated for a partition."},
			{Text: "--line-reader", Description: "The class name of the class to use for reading lines from standard in."},
			{Text: "--max-memory-bytes", Description: "The total memory used by the producer to buffer records waiting to be sent to the server."},
			{Text: "--metadata-expiry-ms", Description: "Number of milliseconds after which we force a refresh of metadata even if we haven't seen any leadership changes."},
			{Text: "--producer-property", Description: "A mechanism to pass user-defined properties in the form key=value to the producer."},
			{Text: "--request-timeout-ms", Description: "The ack timeout of the producer requests. Value must be non-negative."},
			{Text: "--retry-backoff-ms", Description: "The amount of time the producer waits before refreshing the metadata of the topic."},
			{Text: "--batch-size", Description: "Number of messages to send in a single batch if they are not being sent synchronously."},
			{Text: "--broker-list", Description: "The Kafka broker(s) to connect to. NOTE: This'll overwrite the selected cluster value."},
		},
		"kafka-acls": {
			{Text: "--resource-pattern-type", Description: "The type of the resource pattern or pattern filter. Default: LITERAL."},
			{Text: "--topic", Description: "Topic for which ACLs should be added or removed."},
			{Text: "--bootstrap-server", Description: "The Kafka broker(s) to connect to. NOTE: This'll overwrite the selected cluster value."},
			{Text: "--idempotent", Description: "Enable idempotence for the producer. Used with '--producer'."},
			{Text: "--remove", Description: "Indicates you are trying to remove ACLs."},
			{Text: "--authorizer", Description: "Fully qualified class name of the authorizer."},
			{Text: "--command-config", Description: "A property file containing configs to be passed to the admin client."},
			{Text: "--delegation-token", Description: "Delegation token to which ACLs should be added or removed."},
			{Text: "--deny-principal", Description: "Principal to deny in the principalType:name format."},
			{Text: "--help", Description: "Print usage information."},
			{Text: "--add", Description: "Indicates you want to add ACLs."},
			{Text: "--allow-host", Description: "Host from which principals listed in --allow-principal will have access."},
			{Text: "--allow-principal", Description: "Principal to allow in the principalType:name format."},
			{Text: "--operation", Description: "Operation that is being allowed or denied."},
			{Text: "--principal", Description: "List ACLs for the specified principal."},
			{Text: "--transactional-id", Description: "The transactionalId to which ACLs should be added or removed."},
			{Text: "--force", Description: "Assume yes for all queries and do not prompt."},
			{Text: "--group", Description: "Consumer Group to which the ACLs should be added or removed."},
			{Text: "--producer", Description: "Convenience options to add or remove ACLs for producer role."},
			{Text: "--cluster", Description: "Add or remove cluster ACLs."},
			{Text: "--consumer", Description: "Convenience options to add or remove ACLs for the consumer role."},
			{Text: "--deny-host", Description: "Host from which principals listed in --deny-principal will be denied access."},
			{Text: "--authorizer-properties", Description: "Properties required to configure an instance of the authorizer. Format: key=value."},
			{Text: "--list", Description: "List ACLs for the specified resource."},
		},
		"kafka-topics": {
			{Text: "--create", Description: "Create a new topic."},
			{Text: "--delete-config", Description: "A topic configuration override to be removed for an existing topic."},
			{Text: "--exclude-internal", Description: "Exclude internal topics when running list or describe command."},
			{Text: "--help", Description: "Print usage information."},
			{Text: "--list", Description: "List all available topics."},
			{Text: "--topic", Description: "The topic to be create, alter, describe, or delete."},
			{Text: "--unavailable-partitions", Description: "If set when describing topics, only show partitions whose leader is not available."},
			{Text: "--alter", Description: "Alter the number of partitions, replica assignment, and/or configuration for a topic."},
			{Text: "--config", Description: "A topic configuration override for the topic being created or altered."},
			{Text: "--disable-rack-aware", Description: "Disable rack aware replica assignment."},
			{Text: "--if-exists", Description: "If set when altering or deleting topics, the action will only execute if the topic exists."},
			{Text: "--partitions", Description: "The number of partitions for the topic being created or altered."},
			{Text: "--zookeeper", Description: "Zookeeper connection string. WARNING: This'll overwrite the selected cluster value."},
			{Text: "--delete", Description: "Delete a topic."},
			{Text: "--force", Description: "Assume yes for all queries and do not prompt."},
			{Text: "--if-not-exists", Description: "If set when creating topics, the action will only execute if the topic does not already exist."},
			{Text: "--replica-assignment", Description: "A list of manual partition-to-broker assignments for the topic being created or altered."},
			{Text: "--topics-with-overrides", Description: "If set when describing topics, only show topics that have overridden configs."},
			{Text: "--describe", Description: "List details for the given topics."},
			{Text: "--replication-factor", Description: "The replication factor for each partition in the topic being created."},
			{Text: "--under-replicated-partitions", Description: "If set when describing topics, only show under replicated partitions."},
		},
		"kafka-dump-log": {
			{Text: "--key-decoder-class", Description: "Used to deserialize the keys. Default: kafka.serializer.StringDecoder."},
			{Text: "--offsets-decoder", Description: "Parse log data as offset data from the __consumer_offsets topic."},
			{Text: "--print-data-log", Description: "Print the messages content when dumping data log. Automatically set when using any decoder options."},
			{Text: "--transaction-log-decoder", Description: "Parse log data as transaction metadata from the __transaction_state topic."},
			{Text: "--value-decoder-class", Description: "Used to deserialize the messages. Default: kafka.serializer.StringDecoder."},
			{Text: "--verify-index-only", Description: "Verify the index log without printing its content."},
			{Text: "--deep-iteration", Description: "Use deep instead of shallow iteration. Automatically set when using '--print-data-log'."},
			{Text: "--files", Description: "Comma-separated list of data and index log files to be dumped."},
			{Text: "--help", Description: "Print usage information."},
			{Text: "--index-sanity-check", Description: "Check the index sanity without printing its content."},
			{Text: "--max-message-size", Description: "Size of largest message. Default: 5242880."},
		},
		"kafka-log-dirs": {
			{Text: "--command-config", Description: "Property file containing configs to be passed to the admin client."},
			{Text: "--describe", Description: "Describe the specified log directories on the specified brokers."},
			{Text: "--topic-list", Description: "The list of topics to be queried in the form 'topic1,topic2,topic3'. Default: all."},
			{Text: "--bootstrap-server", Description: "The Kafka broker(s) to connect to. NOTE: This'll overwrite the selected cluster value."},
			{Text: "--broker-list", Description: "The list of brokers to be queried in the form '0,1,2'. Default: all."},
		},
		"kafka-delete-records": {
			{Text: "--bootstrap-server", Description: "The Kafka broker(s) to connect to. NOTE: This'll overwrite the selected cluster value."},
			{Text: "--command-config", Description: "Property file containing configs to be passed to the admin client."},
			{Text: "--offset-json-file", Description: "The JSON file with offset per partition."},
		},
	}
)
