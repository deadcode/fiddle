import { client } from './client';
import { randomBytes } from 'crypto';

export const withLock = async (key: string, cb: () => any) => {
	const retryDelayMs = 100;
	let retries = 20;

	// Generate random value to store at lock key
	const token = randomBytes(6).toString('hex');
	// Create the lock key
	const lockKey = `lock:${key}`;

	// While loop to implement the retry
	while (retries >= 0) {
		retries--;
		// Try SET NX operation
		const locked = await client.set(lockKey, token, {
			NX: true,
			PX: 2000 // Delete lock in case of crash/error-throw
		});

		if (!locked) {
			// ELSE retry after retry delay pause
			await pause(retryDelayMs);
			continue;
		}

		// IF set is successful, run the callback
		// After callback is done, unset the lock key (unlock)
		try {
			const result = cb();
			return result;
		} finally {
			await client.unlock(lockKey, token);
		}
	}
};

const buildClientProxy = () => {};

const pause = (duration: number) => {
	return new Promise((resolve) => {
		setTimeout(resolve, duration);
	});
};
