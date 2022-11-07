export const CreateWebSocket = () => {
    return new WebSocket('ws://' + 'localhost:5070' + '/ws');
};
