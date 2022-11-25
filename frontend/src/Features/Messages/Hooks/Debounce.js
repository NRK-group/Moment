let timeOutId;
export const debounce = (call, time) => {
    if (timeOutId) {
        clearTimeout(timeOutId);
    }
    timeOutId = setTimeout(() => {
        call();
    }, time);
};
