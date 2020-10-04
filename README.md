
# Postbox

Mail delivery service, via http.

Postbox is a service for sending emails over http. Under the hood, it uses the smtp protocol. The service allows you to use smtp users by connecting to other smtp servers. You can enter your google / yandex / yahoo account details.

You can connect postbox as a microservice to your project. Postbox is easy to run with docker or docker-compose.

Postbox uses MongoDB as storage. Therefore, env variables need to specify data for connecting to MongoDB. You can use docker to set up MongoDB, or you can use ready-made MongoDB.
