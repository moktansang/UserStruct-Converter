1. This is a simple implementation of emitter. There are three folders in the repository.
2. Client folder is the publisher, the server.go file inside this folder is a publisher to the channel provided in config.yml file.
3.  Configure the emitter server. Update the emitter.conf file with the license_key and use the secret_key, provide a channel and select permissions at  `http://127.0.0.1:8080/keygen` link in a browser to get the new secret key. Set Time-To-Live to be a high number approx 1000 so that secret_key doesn't expire soon. Update the secret key and channel in config.yml file. More info at https://github.com/emitter-io/emitter.
4. Server folder is the subscriber, the server.go file inside this folder is subscriber to the channel provided in config.yml file and waits for messages.
5. Once the emitter server is up and running, run the subscriber by navigating into server folder and running 'go run server.go' command in a different shell.
6. Once subscriber is up run publisher by navigating into client folder and running 'go run client.go' command in a different shell.
7. When client starts publishing to channel, subscriber shell should start printing the published messages. 
8. Messages along with the number of times to publish can be configured from config.yml file.