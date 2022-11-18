import './Groups.css';

import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import Post from '../../features/Post';
import Modal from '../../features/Modal';
import Input from '../../components/Input/Input';
import MiniUserCard from '../../components/MiniUserCard/MiniUserCard';
import GroupList from './components/GroupList';
import {
    ChevronRightIcon,
    ChevronLeftIcon,
    BarsIcon,
    GroupsIcon,
} from '../../components/Icons/Icons';
import { useRef, useState } from 'react';

function Groups({ isMobile }) {
    let bodyStyleName = isMobile ? 'mobile' : 'desktop';
    let cardStyleName = isMobile ? 'mobileCard' : 'desktopCard';

    const GroupsLeftMenu = useRef(null);
    const GroupsRightMenu = useRef(null);
    const GroupsPostsArea = useRef(null);

    const [toggle, setToggle] = useState(true);
    const [groupS, setGroupS] = useState([]);
    const [groupSM, setGroupSM] = useState([]);
    const [openModal, setOpenModal] = useState(false);

    const OpenGroupsLeftMenu = () => {
        GroupsLeftMenu.current.style.width = '670px';
        GroupsLeftMenu.current.style.display = 'inline-table';
    };

    const CloseGroupsLeftMenu = () => {
        GroupsLeftMenu.current.style.width = '0%';
        GroupsLeftMenu.current.style.display = 'none';
    };

    const OpenGroupsRightMenu = () => {
        GroupsRightMenu.current.style.width = '670px';
        GroupsRightMenu.current.style.display = 'inline-table';
        GroupsPostsArea.current.style.width = '50%';
    };

    const CloseGroupsRightMenu = () => {
        GroupsRightMenu.current.style.width = '0%';
        GroupsRightMenu.current.style.display = 'none';
        GroupsPostsArea.current.style.width = '100%';
    };

    const Events = () => {
        CloseGroupsRightMenu();
        GroupsPostsArea.current.style.width = '100%';
        setToggle(false);
    };

    const Users = () => {
        GroupsRightMenu.current.style.display = 'block';
        GroupsRightMenu.current.style.width = '25%';
        GroupsPostsArea.current.style.width = '50%';
        setToggle(true);
    };

    const Getgroups = async () => {
        let fetchGroups = await fetch('http://localhost:5070/group', {
            credentials: 'include',
            method: 'GET',
        })
            .then(async (resp) => await resp.json())
            .then((data) => data);
        console.log({ fetchGroups });
        setGroupSM(fetchGroups);
        setOpenModal(true);
    };

    return (
        <Body styleName={bodyStyleName}>
            <Card styleName={cardStyleName}>
                <div className='Groups'>
                    {openModal && (
                        <Modal setOpenModal={setOpenModal}>
                            {' '}
                            <GroupList data={groupSM} />
                        </Modal>
                    )}
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
                            {groupS.length > 0 ? (
                                groupS.map((ele) => {})
                            ) : (
                                <div className='Join-Group'>
                                    {' '}
                                    <span onClick={() => Getgroups()}>
                                        Join a group
                                    </span>{' '}
                                </div>
                            )}
                        </div>
                        <div className='GroupPageMenu'>
                            {groupS.length > 0 ? (
                                <>
                                    <p
                                        style={{
                                            marginTop: '12px',
                                            cursor: 'pointer',
                                        }}
                                        onClick={() => {}}>
                                        Create Events
                                    </p>

                                    <p
                                        style={{
                                            marginTop: '12px',
                                            cursor: 'pointer',
                                        }}
                                        onClick={() => {}}>
                                        {' '}
                                        Create group Post
                                    </p>
                                </>
                            ) : (
                                <></>
                            )}
                        </div>
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
                        <div className='Group-Post'></div>
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
                        {'   Group members'}
                    </div>
                </div>
            </Card>
        </Body>
    );
}

export default Groups;
