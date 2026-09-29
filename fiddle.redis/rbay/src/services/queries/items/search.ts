import { client } from '$services/redis';
import { deserialize } from './deserialize';
import { itemsIndexKey } from '$services/keys';

export const searchItems = async (term: string, size: number = 5) => {
	const cleaned = term
		.replaceAll('/[^a-zA-Z0-9 /g', '') // remove all special symbols
		.trim() // remove spaces at start and end of sentence
		.split(' ') // break sentence into array of words
		.map((word) => (word ? `%${word}%` : '')) // remove spaces between words, wrap words in '%' to trigger fuzzy search
		.join(' '); // join all words into sentence with single space to trigger 'AND' search

	// Validate cleaned search term
	if (!cleaned) {
		return []; // return empty search results
	}

	// execute search
	const results = await client.ft.search(itemsIndexKey(), cleaned, {
		LIMIT: {
			from: 0,
			size: size
		}
	});
	//console.log(results);

	// deserialze and return results
	return results.documents.map(({ id, value }) => deserialize(id, value as any));
};
