export const CreateWebSocket = () => {
    return new WebSocket('ws://' + config.api + '/ws');
};
