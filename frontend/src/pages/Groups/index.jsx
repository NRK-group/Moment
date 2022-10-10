import './Groups.css';

import Avatar from '../../components/Avatar';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import { useRef, useState } from 'react';

function Groups({ bodyStyleName, cardStyleName }) {

    const dropdown = useRef(null);
    const [toggle, setToggle] = useState(true);

    const OpenDropdownMenu = () => {
        setToggle(!toggle);
        if (toggle) {
            console.log('inside');
            dropdown.current.style.display = 'block';
        } else {
            dropdown.current.style.display = 'none';
        }
    };

    return (
        <Body styleName={bodyStyleName}>
            <Card styleName={cardStyleName}>
                <div className='Groups'>
                    <div className='GroupsLeftMenu'></div>
                    <div className='GroupsPostsArea'>
                    <Card styleName={'PostHeader'}>
                    <div style={{ display: 'flex' }}>
                        <Avatar
                            avatarSrc={"https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Google_%22G%22_Logo.svg/2048px-Google_%22G%22_Logo.svg.png"}
                            styleName={'PostAvatarUsers'}
                        />

                        <p style={{ marginLeft: '4px' }}>{name}</p>
                    </div>

                    <div className='PostHeaderMenu'>
                        <button
                            className='dropbtn'
                            onClick={() => OpenDropdownMenu()}>
                            <i className='fa-solid fa-ellipsis' />
                        </button>
                        <div ref={dropdown} className='dropdown-content'>
                            <a href='#'>option 1</a>
                            <a href='#'>option 2</a>
                            <a href='#'>option 3</a>
                        </div>
                    </div>
                </Card>
                        
                    </div>

                    <div className='GroupsRightMenu'></div>
                </div>
            </Card>
        </Body>
    );
}

export default Groups;
