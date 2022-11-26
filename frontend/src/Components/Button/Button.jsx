import './Button.css';

export const Button = ({ styleName, content, action }) => {
    return (
        <button
            onClick={(e) => {
                action();
                e.stopPropagation();
            }}
            className={`${styleName} btn`}>
            {content}
        </button>
    );
};
