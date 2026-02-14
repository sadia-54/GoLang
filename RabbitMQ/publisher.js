const amqp = require("amqplib");

const msg = {name: process.argv[2]};

connect()

async function connect() {
    try{

        const connection = await amqp.connect("amqp://localhost:5672");
        const channel = await connection.createChannel();
        const result = await channel.assertQueue("jobs");

        channel.sendToQueue("jobs", Buffer.from(JSON.stringify(msg)));
        console.log(`Message sent ${JSON.stringify(msg)}`);

    } catch (err) {
        console.log(err);
    }
}