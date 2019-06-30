export function postChat(message) {
    const uuidv4 = require('uuid/v4');
    return {
        communityId: "a6bc86ea-bd09-4b62-a1d8-ff6c05558a35",
        postId: uuidv4(),
        postNumber: 10, // server increment last post number. but this func cannot get last post number.
        postUser: {
            userName: "You",
            userId: "dc0ab403-6f47-4cac-abb4-0b82f13843cb",
            userIcon: "/img/you.png",
        },
        postBody: {
            text: message,
        }
    };
}