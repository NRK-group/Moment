import './Groups.css';
import Avatar from '../../components/Avatar';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import Post from '../../features/Post';
import Input from '../../components/Input/Input';
import MiniUserCard from '../../components/MiniUserCard/MiniUserCard';
import {
    ChevronRightIcon,
    ChevronLeftIcon,
    BarsIcon,
    GroupsIcon,
} from '../../components/Icons/Icons';
import { useRef } from 'react';

function Groups({ bodyStyleName, cardStyleName }) {
    const GroupsLeftMenu = useRef(null);
    const GroupsRightMenu = useRef(null);
    const GroupsPostsArea = useRef(null);

    const OpenGroupsLeftMenu = () => {
        GroupsLeftMenu.current.style.width = '670px';
    };

    const CloseGroupsLeftMenu = () => {
        GroupsLeftMenu.current.style.width = '0%';
    };

    const OpenGroupsRightMenu = () => {
        GroupsRightMenu.current.style.width = '670px';
    };

    const CloseGroupsRightMenu = () => {
        GroupsRightMenu.current.style.width = '0%';
    };

    const Events = () => {
        CloseGroupsRightMenu()
        GroupsPostsArea.current.style.width = "100%"
    }

    return (
        <Body styleName={bodyStyleName}>
            <Card styleName={cardStyleName}>
                <div className='Groups'>
                    <div className='GroupsLeftMenu' ref={GroupsLeftMenu}>
                        <span
                            onClick={() => {
                                CloseGroupsLeftMenu();
                            }}
                            className='GroupsLeftMenuIcon'>
                            <ChevronLeftIcon />
                        </span>
                        <div className='GroupsMenuHeader'>
                            <h2>Groups</h2>
                        </div>
                        <div className='GroupsList'>
                            <MiniUserCard name={'Go'} />
                            <MiniUserCard name={'Fishing'} />
                            <MiniUserCard name={'Racing'} />
                        </div>
                        <p
                            style={{ marginTop: '12px', cursor: 'pointer' }}
                            onClick={() => {Events()}}>
                            Events
                        </p>
                    </div>
                    <div ref={GroupsPostsArea} className='GroupsPostsArea'>
                        <Card styleName={'PostHeader'}>
                            <div style={{ display: 'flex' }}>
                                <div
                                    className='GroupHeaderIcons'
                                    onClick={() => {
                                        OpenGroupsLeftMenu();
                                    }}>
                                    <BarsIcon />
                                </div>
                                <Avatar
                                    avatarSrc={
                                        'https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Google_%22G%22_Logo.svg/2048px-Google_%22G%22_Logo.svg.png'
                                    }
                                    styleName={'PostAvatarUsers'}
                                />

                                <p style={{ marginLeft: '4px' }}>{name}</p>
                            </div>
                            <div
                                className='GroupHeaderIcons'
                                onClick={() => {
                                    OpenGroupsRightMenu();
                                }}>
                                <GroupsIcon />
                            </div>
                        </Card>

                        <Post />
                        <Post />
                    </div>

                    <div ref={GroupsRightMenu} className='GroupsRightMenu'>
                        <span
                            onClick={() => {
                                CloseGroupsRightMenu();
                            }}
                            className='GroupsRightMenuIcon'>
                            {' '}
                            <ChevronRightIcon />{' '}
                        </span>
                        <div className='GroupsMenuHeader'>
                            <Input
                                styleName={'search'}
                                type={'search'}
                                placeholder={'Search User'}
                            />
                        </div>
                        <MiniUserCard />
                        <MiniUserCard />
                        <MiniUserCard />
                        <MiniUserCard />
                    </div>
                </div>
            </Card>
        </Body>
    );
}

export default Groups;
