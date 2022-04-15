const path = require('path')
const fs = require('fs')

var httpPort = null
// export LOGGING=true; ./serve.sh
let fastifyServerOptions = {
  logger: process.env.LOGGING || true
}

if(process.env.REGISTRY_CRT && process.env.REGISTRY_KEY) {
  fastifyServerOptions.https = {
    cert: fs.readFileSync(path.join(`${__dirname}/`, process.env.REGISTRY_CRT)),
    key: fs.readFileSync(path.join(`${__dirname}/`, process.env.REGISTRY_KEY))
  }
  httpPort = process.env.REGISTRY_HTTPS || 7070
} else {
  httpPort = process.env.REGISTRY_HTTP || 7070
}

const fastify = require('fastify')(fastifyServerOptions)

// folder where to store wasm file
const wasmFunctionsFolder = process.env.REGISTRY_FUNCTIONS_PATH || './wasm.functions'

// 🚧 (not used)
const adminToken = process.env.ADMIN_REGISTRY_TOKEN || ""

fastify.register(require('fastify-formbody'))
fastify.register(require('fastify-multipart'))


let routesOptions = {
  wasmFunctionsFolder: wasmFunctionsFolder,
  adminToken: adminToken
}

fastify.register(require('./routes/publish.js'), routesOptions)
fastify.register(require('./routes/download.js'), routesOptions)


// Serve the static assets
fastify.register(require('fastify-static'), {
  root: path.join(__dirname, 'public'),
  prefix: '/'
})


// Run the server!
const start = async _ => {
  try {
    await fastify.listen(httpPort, "0.0.0.0")
    fastify.log.info(`server listening on ${fastify.server.address().port}`)
  } catch (error) {
    fastify.log.error(error)
  }
}
start()

