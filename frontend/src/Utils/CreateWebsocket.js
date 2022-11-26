import config from '../../config';

export const CreateWebSocket = () => {
    return new WebSocket(config.socket);
};
