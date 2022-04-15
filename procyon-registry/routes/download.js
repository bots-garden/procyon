const path = require('path')

async function download (fastify, options) {

  fastify.get('/get/:filename', async (request, reply) => {
    let filename = request.params.filename
    reply.sendFile(filename, options.wasmFunctionsFolder)

    await reply
  })
}

module.exports = download
