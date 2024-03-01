import { promises as fs } from 'node:fs'
import * as path from 'node:path'
import * as url from 'node:url'

const __dirname = path.dirname(url.fileURLToPath(import.meta.url))
const inputFile = path.resolve(__dirname, './input.txt')
const input = await fs.readFile(inputFile, 'utf-8')
const lines = input.split('\n').filter(Boolean)

console.log(input)
console.dir(lines)
