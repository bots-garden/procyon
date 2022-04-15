const fs = require('fs')
const streamToFile = require('../helpers/stream.helper').streamToFile

async function publish (fastify, options) {

  // route protection
  /*
  fastify.addHook('onRequest', async (request, reply) => {
    let token = request.headers["admin_registry_token"]
    if (options.adminToken==="" || options.adminToken===token) {
      // all good
    } else {
      reply.code(401).send({
        failure: "😡 Unauthorized",
        success: null
      })
    }
  })
  */

  fastify.post(`/publish/:version`, async (request, reply) => {
    let version = request.params.version

    const data = await request.file()

    /* properties of data
      data.file // stream
      data.fields // other parsed parts
      data.fieldname
      data.filename
      data.encoding
      data.mimetype
    */
        
    try {
      // Get the file from the stream
      const wasmFile = await streamToFile({isString:false, stream:data.file})
      // Write the file on disk
      fs.writeFileSync(`${options.wasmFunctionsFolder}/${data.fieldname}.${version}.wasm`, wasmFile)

      reply.send({
        failure: null,
        success: `uploaded: ${options.wasmFunctionsFolder}/${data.fieldname}.${version}.wasm`
      })

    } catch(error) {
      console.log("😡", error)
      reply.send({
        failure: "error when loading wasm file",
        success: null
      })
    }
  
    await reply

  })

}

module.exports = publish
