import { uniqueNamesGenerator, adjectives, animals } from 'unique-names-generator';

export const generateName = (seed) => {
    return uniqueNamesGenerator({
        dictionaries: [adjectives, animals],
        length: 2,
        seed: seed,
    });
}

