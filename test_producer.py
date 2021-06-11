import time
import json
import kafka
import sys

try:
  producer = kafka.KafkaProducer(
    bootstrap_servers = ['localhost:9092'],
  )
except kafka.errors.NoBrokersAvailable:
  print('kafka server not found')
  sys.exit(1)

for index in range(100):
  data = json.dumps({'counter': index}).encode('utf-8')
  producer.send('test', value=data)
  time.sleep(0.5)
  