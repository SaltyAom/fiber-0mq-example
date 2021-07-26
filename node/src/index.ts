import { Router, Context } from 'zeromq'
import PQueue from 'p-queue'

import { cpus } from 'os'

const requestQueue = new PQueue({
	concurrency: Math.pow(2, 16)
})

const socket = new Router({
	sendTimeout: 0,
	context: new Context({
		blocky: false,
		ioThreads: cpus().length
	})
})

const channel = 'tcp://0.0.0.0:5556'

const sendRequest = async (response: [string, string]) => {
	requestQueue.add(async () => {
		await socket.send(response)
	})
}

const handle = async (request: Buffer[]) => {
	let [id, message] = request.toString().split(',')

	sendRequest([id, Date().toLocaleString()])
}

const main = async () => {
	try {
		await socket.bind(channel)

		console.log('Junbi Ok!')

		while (true)
			// ? Process without waiting for handle to be done to acheive asynchronous message
			handle(await socket.receive())
	} catch (error) {
		console.error(error)
		process.exit(0)
	}
}

process.on('exit', async () => {
	await socket.unbind(channel)
	await socket.close()
})

main()
