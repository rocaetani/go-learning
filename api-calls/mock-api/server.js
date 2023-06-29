const express = require('express')

const Chance = require('chance')

const server = express()

// Information to be displayed on the cards
const minValue = 0
const maxValue = 500
const overflowValue = 99999
const random = new Chance();

const port = 4001
server.use(express.json())

server.post('/info', (req, res) => {

    console.log(req.body); // JavaScript object containing the parse JSON
    //res.json(req.body);
  
    console.log(req.headers.authorization)

    res.json(
        {
            "users": [
                {
                "name": "Rodrigo",
                "code": random.integer({ min: overflowValue, max: 2 * overflowValue }),
                "other": "Heyyy",
                },
                {
                    "name": "Carla",
                    "code": random.integer({ min: overflowValue, max: 2 * overflowValue }),
                    "other": "Hello",
                },
            ]
        }
    )
   
})

server.listen(port, () => {
    console.log(`Mock API is running on port ${port}.`)
})
