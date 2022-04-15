function streamToFile({isString, stream}) {
  const chunks = [];
  return new Promise((resolve, reject) => {
    stream.on('data', (chunk) => chunks.push(Buffer.from(chunk)))
    stream.on('error', (err) => reject(err))
    if(isString) {
      stream.on('end', () => resolve(Buffer.concat(chunks).toString('utf8')))
    } else {
      stream.on('end', () => resolve(Buffer.concat(chunks)))
    }
  })
}

exports.streamToFile = streamToFile
