var redis = require("redis"),
    client = redis.createClient();

const fixture = require('./episode')

client.set('episodes:pid1:web', JSON.stringify(fixture))

process.exit(1);
