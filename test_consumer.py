import time
import json
import kafka

# PLAN: have a multithreading version and a multiprocessing version 

# check the different python data structures that are thread-safe, see how data is synchronized in the case of multiprocessing

# have a worker executor class

# first a basic implementation

consumer = kafka.KafkaConsumer(
  'test', # first is the topic name
  bootstrap_servers=['localhost:9092'],
  group_id = 'first_group',
  auto_offset_reset = 'earliest'
)

print(consumer)

for e in consumer:
  print(e)
  print(e.value)