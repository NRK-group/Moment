export const CalculateTimeDiff = (date) => {
    const now = new Date();
    const notifDate = new Date(date);
    const diff = now.getTime() - notifDate.getTime();
    const diffDays = Math.floor(diff / (1000 * 3600 * 24));
    const diffHours = Math.floor(diff / (1000 * 3600));
    const diffMinutes = Math.floor(diff / (1000 * 60));
    if (diffDays > 0) return diffDays + 'd ago';
    if (diffHours > 0) return diffHours + 'h ago';
    if (diffMinutes > 0) return diffMinutes + 'm ago';
    return 'now';
};
