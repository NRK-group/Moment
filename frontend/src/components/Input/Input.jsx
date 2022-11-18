import Icon from '../Icon/Icon';
import './Input.css';
import { useState } from 'react';

const Input = ({ styleName, placeholder, type, onChange }) => {
    const [focused, setFocused] = useState(false);
    const onFocus = () => setFocused(true);
    const onBlur = () => setFocused(false);
    return (
        <div className={styleName}>
            {type === 'search' ? (
                <span className='searchContainer'>
                    {!focused && (
                        <Icon>
                            <svg
                                xmlns='http://www.w3.org/2000/svg'
                                width='16'
                                height='16'
                                viewBox='0 0 512 512'>
                                <path d='M416 208c0 45.9-14.9 88.3-40 122.7L502.6 457.4c12.5 12.5 12.5 32.8 0 45.3s-32.8 12.5-45.3 0L330.7 376c-34.4 25.2-76.8 40-122.7 40C93.1 416 0 322.9 0 208S93.1 0 208 0S416 93.1 416 208zM208 352c79.5 0 144-64.5 144-144s-64.5-144-144-144S64 128.5 64 208s64.5 144 144 144z' />
                            </svg>
                        </Icon>
                    )}
                    <input
                        onFocus={onFocus}
                        onBlur={onBlur}
                        className='input'
                        type='text'
                        onChange={onChange}
                        placeholder={'Search'}
                    />
                </span>
            ) : (
                <input
                    className='input'
                    type='text'
                    placeholder={placeholder}
                    onChange={onChange}
                />
            )}
        </div>
    );
};
export default Input;
