'use strict'
require('dotenv').config()
module.exports = {
  NODE_ENV: '"production"',
  GBOOKSHELF_SERVER_URL: JSON.stringify(process.env.GBOOKSHELF_SERVER_URL)
}
