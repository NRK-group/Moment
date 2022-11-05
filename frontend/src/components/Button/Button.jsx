import './Button.css';

export const Button = ({ styleName, content, action }) => {
    return (
        <button
            onClick={() => {
                action();
            }}
            className={`${styleName} btn`}>
            {content}
        </button>
    );
};
