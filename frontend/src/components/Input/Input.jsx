import Icon from '../Icon/Icon';
import './Input.css';
import { useState } from 'react';

const Input = ({ styleName, placeholder, type }) => {
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
                        placeholder={placeholder}
                    />
                    {focused && (
                        <Icon>
                            <svg
                                width='16'
                                height='16'
                                xmlns='http://www.w3.org/2000/svg'
                                viewBox='0 0 512 512'>
                                <path
                                    d='M175 175C184.4 165.7 199.6 165.7 208.1
                                175L255.1 222.1L303 175C312.4 165.7 327.6 165.7
                                336.1 175C346.3 184.4 346.3 199.6 336.1
                                208.1L289.9 255.1L336.1 303C346.3 312.4 346.3
                                327.6 336.1 336.1C327.6 346.3 312.4 346.3 303
                                336.1L255.1 289.9L208.1 336.1C199.6 346.3 184.4
                                346.3 175 336.1C165.7 327.6 165.7 312.4 175
                                303L222.1 255.1L175 208.1C165.7 199.6 165.7
                                184.4 175 175V175zM512 256C512 397.4 397.4 512
                                256 512C114.6 512 0 397.4 0 256C0 114.6 114.6 0
                                256 0C397.4 0 512 114.6 512 256zM256 48C141.1 48
                                48 141.1 48 256C48 370.9 141.1 464 256 464C370.9
                                464 464 370.9 464 256C464 141.1 370.9 48 256
                                48z'
                                />
                            </svg>
                        </Icon>
                    )}
                </span>
            ) : (
                <input
                    className='input'
                    type='text'
                    placeholder={placeholder}
                />
            )}
        </div>
    );
};
export default Input;
