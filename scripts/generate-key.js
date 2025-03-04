// @ts-check
import crypto from "crypto"
import fs from "fs"
import path from "path"

const rawToken = Buffer.from(crypto.randomBytes(48)).toString("hex")
const tokenToFeed = `om_${rawToken}`

const __pwd = process.cwd()

const envFilePath = path.join(__pwd, ".env")

if (!fs.existsSync(envFilePath)) {
  fs.writeFileSync(envFilePath, "")
}

const envFile = fs.readFileSync(envFilePath, "utf8")

const newEnvFileContent = envFile.includes("API_KEY")
  ? envFile.replace(/API_KEY=.*/, `API_KEY=${tokenToFeed}`)
  : `${envFile}\API_KEY=${tokenToFeed}`

fs.writeFileSync(envFilePath, newEnvFileContent.trim())

console.log(`Token generated! This token has already populated from an .env file. \n\n${tokenToFeed}\n`)
