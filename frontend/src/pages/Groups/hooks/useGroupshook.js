export const GetAllGroupPosts = async (id) => {
    let fetchAllgroupPosts = await fetch(
        `http://localhost:5070/getGroupPost?groupId=${id}`,
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
    let fetchAllgroupEvents = await fetch(
        `http://localhost:5070/event?groupId=${id}`,
        {
            credentials: 'include',
            method: 'GET',
        }
    )
        .then(async (resp) => await resp.json())
        .then((data) => data);
    return fetchAllgroupEvents;
};

export const GetAllNonMembers = async (id) => {
    let fetchAllNonMembers = await fetch(
        `http://localhost:5070/groupNonMembers?groupId=${id}`,
        {
            credentials: 'include',
            method: 'GET',
        }
    )
        .then(async (resp) => await resp.json())
        .then((data) => data);
    return fetchAllNonMembers;
};


export const RequestTogroup = async (id, receiverId, socket, type) => {
    socket.send(JSON.stringify({type: "groupInvitation"+type, senderId: id, receiverId:receiverId}))
};
