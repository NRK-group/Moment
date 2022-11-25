import config from '../../../../config';

export const GetAllGroupPosts = async (id) => {
    let fetchAllgroupPosts = await fetch(
        `${config.api}/getGroupPost?groupId=${id}`,
        {
            credentials: 'include',
            method: 'GET',
        }
    )
        .then(async (resp) => await resp.json())
        .then((data) => data);
    return fetchAllgroupPosts;
};

export const GetAllGroupEvents = async (id) => {
    let fetchAllgroupEvents = await fetch(`${config.api}/event?groupId=${id}`, {
        credentials: 'include',
        method: 'GET',
    })
        .then(async (resp) => await resp.json())
        .then((data) => data);
    return fetchAllgroupEvents;
};

export const GetAllNonMembers = async (id) => {
    let fetchAllNonMembers = await fetch(
        `${config.api}/groupNonMembers?groupId=${id}`,
        {
            credentials: 'include',
            method: 'GET',
        }
    )
        .then(async (resp) => await resp.json())
        .then((data) => data);
    return fetchAllNonMembers;
};

//groupInvitation

export const RequestToS = (id, receiverId, socket, type, groupId) => {
    socket.send(
        JSON.stringify({
            type: type,
            senderId: id,
            receiverId: receiverId,
            groupId: groupId,
        })
    );
};
