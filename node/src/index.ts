import { Reply } from 'zeromq'

const socket = new Reply()

const channel = 'tcp://0.0.0.0:5556'

const main = async () => {
	try {
		await socket.bind(channel)

        console.log("Junbi Ok!")

		for await (const [message] of socket) {
			socket.send(Buffer.from(new Date().toLocaleString()))
		}
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
