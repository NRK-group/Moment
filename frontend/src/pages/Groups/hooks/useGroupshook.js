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
        console.log({fetchAllgroupPosts})
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
    console.log({ fetchAllgroupEvents });
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
        console.log({fetchAllNonMembers})
    return fetchAllNonMembers;
};


export const RequestToJoin = async (id, receiverId, socket) => {
    socket.send(JSON.stringify({type: "groupInvitation", senderId: id, receiverId:receiverId}))
};
