import type { CreateUserAttrs } from '$services/types';
import { genId } from '$services/utils';
import { client } from '$services/redis';
import { usersKey } from '$services/keys';
import { usernamesUniqueKey, userNamesKey } from '$services/keys';

export const getUserByUsername = async (username: string) => {
    // Get 'username's user id
    const decimalId = await client.zScore(userNamesKey(), username);

    // Check userId is valid
    if (!decimalId) {
        throw new Error('User does not exist');
    }

    // Convert userId to hex and fetch user hash
    const id = decimalId.toString(16);
    const user = await client.hGetAll(usersKey(id));

    // desrialize hash and return
    return deserialize(id, user);
};

export const getUserById = async (id: string) => {
    const user = await client.hGetAll(usersKey(id));

    return deserialize(id, user);
};

export const createUser = async (attrs: CreateUserAttrs) => {
    const id = genId();

    // Check if the username is already in already created usernames
    const exists = await client.sIsMember(usernamesUniqueKey(), attrs.username);
    if (exists) {
        throw new Error('Username is taken');
    }

    await client.hSet(usersKey(id), serialize(attrs));
    await client.sAdd(usernamesUniqueKey(), attrs.username);
    await client.zAdd(userNamesKey(), {
        value: attrs.username,
        score: parseInt(id, 16)
    });

    return id;
};

const serialize = (user: CreateUserAttrs) => {
    return {
        username: user.username,
        password: user.password
    };
};

const deserialize = (id: string, user: { [key: string], string }) => {
    return {
        id: id,
        username: user.username,
        password: user.password
    };
};