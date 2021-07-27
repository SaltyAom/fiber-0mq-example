import { Router } from 'zeromq'
import PQueue from 'p-queue'

const requestQueue = new PQueue({
	concurrency: 1
})

const socket = new Router()

const channel = 'tcp://0.0.0.0:5556'

const handle = async (request: Buffer[]) => {
	let [id, message] = request.toString().split(',')

	requestQueue.add(async () => {
		await socket.send([id, Date().toLocaleString()])
	})
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
