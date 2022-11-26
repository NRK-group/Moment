export const MessageContent = ({ containerStyle, type, content, date }) => {
    return (
        <div className={containerStyle}>
            <div className={type}>
                <div className='message'>
                    <div className='messageContent'>{content}</div>
                    <div className='date'>{date}</div>
                </div>
            </div>
        </div>
    );
};
