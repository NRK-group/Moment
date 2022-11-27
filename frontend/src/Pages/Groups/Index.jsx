import './Groups.css';
import { GetCookie } from '../Profile/ProfileData';
import { useNavigate } from 'react-router-dom';
import Body from '../../Components/Body/Body';
import Card from '../../Components/Card/Card';
import Post from '../../Features/Post/Index';
import AddGroup from './Components/AddGroup/Index';
import Modal from '../../Features/Modal/Index';
import MiniUserCard from '../../Components/MiniUserCard/MiniUserCard';
import GroupList from './Components/GroupList/Index';
import GroupPost from './Components/GroupPost/Index';
import GroupEvent from './Components/GroupEvents/Index';
import Event from '../../Features/Event/Index';

import {
    RequestToS,
    GetAllGroupPosts,
    GetAllGroupEvents,
    GetAllNonMembers,
} from './Hooks/UseGroupshook';

import {
    ChevronRightIcon,
    ChevronLeftIcon,
    BarsIcon,
    GroupsIcon,
} from '../../Components/Icons/Icons';
import { useRef, useState, useEffect } from 'react';

function Groups({ isMobile, socket }) {
    let bodyStyleName = isMobile ? 'mobile' : 'desktop';
    let cardStyleName = isMobile ? 'mobileCard' : 'desktopCard';

    const navigate = useNavigate();

    const GetAllUsergroups = async () => {
        let fetchAllUsergroups = await fetch(
            'http://localhost:5070/getUserGroups',
            {
                credentials: 'include',
                method: 'GET',
            }
        )
            .then(async (resp) => await resp.json())
            .then((data) => data);
        if (fetchAllUsergroups !== null) {
            console.log(fetchAllUsergroups[0]);
            setGroupS(fetchAllUsergroups);
            setGroupSelect(fetchAllUsergroups[0]);
            let holder = await GetAllGroupPosts(fetchAllUsergroups[0].GroupID);
            setGroupPosts(holder);
            holder = await GetAllGroupEvents(fetchAllUsergroups[0].GroupID);
            setGroupE(holder);
        }
    };

    const GroupsLeftMenu = useRef(null);
    const GroupsRightMenu = useRef(null);
    const GroupsPostsArea = useRef(null);

    const [groupPosts, setGroupPosts] = useState(null);
    const [flag, setFlag] = useState(false);
    const [groupS, setGroupS] = useState(null);
    const [getallNonMembers, setGetallNonMembers] = useState(null);
    const [displayAllNonMembers, setDisplayAllNonMembers] = useState(null);
    const [groupE, setGroupE] = useState(null);
    const [groupSelect, setGroupSelect] = useState(null);
    const [openModal, setOpenModal] = useState(false);
    const [ele, setEle] = useState(null);

    useEffect(() => {
        GetAllUsergroups();
    }, [flag]);

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

    const Getgroups = async () => {
        let fetchGroups = await fetch('http://localhost:5070/group', {
            credentials: 'include',
            method: 'GET',
        })
            .then(async (resp) => await resp.json())
            .then((data) => data);

        setEle(<GroupList socket={socket} data={fetchGroups} />);
        setOpenModal(true);
    };

    const switchGroup = async (element) => {
        setGroupSelect(element);
        let holder = await GetAllGroupPosts(element.GroupID);
        setGroupPosts(holder);
        holder = await GetAllGroupEvents(element.GroupID);
        setGroupE(holder);
    };

    const dropdown = useRef(null);
    const InputRef = useRef(null);
    const [toggle, setToggle] = useState(false);

    const OpenDropdownMenu = async () => {
        setToggle(!toggle);
        if (toggle) {
            let resp = await GetAllNonMembers(groupSelect.GroupID);
            setGetallNonMembers(resp);
            setDisplayAllNonMembers(resp);
            dropdown.current.style.display = 'block';
        } else {
            dropdown.current.style.display = 'none';
        }
    };

    window.onclick = function (e) {
        if (
            !toggle &&
            dropdown.current.style.display === 'block' &&
            e.target !== InputRef.current
        ) {
            dropdown.current.style.display = 'none';
            setToggle(!toggle);
        }
    };

    const SearchUsers = (searchText) => {
        if (searchText !== '') {
            setDisplayAllNonMembers(
                getallNonMembers.filter((ele) =>
                    ele.firstName
                        .toLowerCase()
                        .includes(searchText.toLowerCase())
                )
            );
        } else {
            setDisplayAllNonMembers(getallNonMembers);
        }
    };

    return (
        <Body styleName={bodyStyleName}>
            <Card styleName={cardStyleName}>
                <div className='Groups'>
                    {openModal && (
                        <Modal setOpenModal={setOpenModal}> {ele}</Modal>
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
                            {groupS !== null ? (
                                groupS.map((ele) => (
                                    <span
                                        key={ele.GroupID}
                                        onClick={() => switchGroup(ele)}>
                                        <MiniUserCard
                                            imgStyleName={'miniUserCardImg'}
                                            optContent={ele.Name}>
                                            {' '}
                                        </MiniUserCard>
                                    </span>
                                ))
                            ) : (
                                <div className='Join-Group'>
                                    {' '}
                                    <span
                                        onClick={() => {
                                            Getgroups();
                                        }}>
                                        Join a group
                                    </span>{' '}
                                </div>
                            )}
                        </div>
                        <div className='GroupPageMenu'>
                            {groupS !== null ? (
                                <>
                                    <p
                                        style={{
                                            marginTop: '12px',
                                            cursor: 'pointer',
                                        }}
                                        onClick={() => {
                                            setEle(
                                                <GroupEvent
                                                    groupId={
                                                        groupSelect.GroupID
                                                    }
                                                    setOpenModal={setOpenModal}
                                                    socket={socket}
                                                    flag={flag}
                                                    setFlag={setFlag}
                                                />
                                            );
                                            setOpenModal(true);
                                        }}>
                                        Create Events
                                    </p>

                                    <p
                                        style={{
                                            marginTop: '12px',
                                            cursor: 'pointer',
                                        }}
                                        onClick={() => {
                                            setEle(
                                                <GroupPost
                                                    groupId={
                                                        groupSelect.GroupID
                                                    }
                                                    setOpenModal={setOpenModal}
                                                    flag={flag}
                                                    setFlag={setFlag}
                                                />
                                            );
                                            setOpenModal(true);
                                        }}>
                                        {' '}
                                        Create group Post
                                    </p>
                                    <p
                                        style={{
                                            marginTop: '12px',
                                            cursor: 'pointer',
                                        }}
                                        onClick={() => {
                                            setEle(
                                                <AddGroup
                                                    setOpenModal={setOpenModal}
                                                    flag={flag}
                                                    setFlag={setFlag}
                                                    socket={socket}
                                                />
                                            );
                                            setOpenModal(true);
                                        }}>
                                        {' '}
                                        Create a group
                                    </p>

                                    <p
                                        style={{
                                            marginTop: '12px',
                                            cursor: 'pointer',
                                        }}
                                        onClick={() => {
                                            Getgroups();
                                        }}>
                                        Join a group
                                    </p>
                                </>
                            ) : (
                                <>
                                    <p
                                        style={{
                                            marginTop: '12px',
                                            cursor: 'pointer',
                                        }}
                                        onClick={() => {
                                            setEle(
                                                <AddGroup
                                                    setOpenModal={setOpenModal}
                                                    flag={flag}
                                                    setFlag={setFlag}
                                                    socket={socket}
                                                />
                                            );
                                            setOpenModal(true);
                                        }}>
                                        {' '}
                                        Create a group
                                    </p>
                                </>
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

                                <p style={{ marginLeft: '4px' }}>
                                    {groupSelect && groupSelect.Name}
                                </p>
                            </div>
                            <div
                                className='GroupHeaderIcons'
                                onClick={() => {
                                    OpenGroupsRightMenu();
                                }}>
                                <GroupsIcon />
                            </div>
                        </Card>
                        <div className='Group-Post'>
                            {groupPosts &&
                                groupPosts.map((data) => (
                                    <Post
                                        key={data.PostID}
                                        avatarSrc={`http://localhost:5070/${data.Image}`}
                                        name={data.NickName}
                                        postContent={data.Content}
                                        userID={data.UserID}
                                        likes={data.NumLikes}
                                        commentsnum={data.NumOfComment}
                                        postBodyImgSrc={data.ImageUpload}
                                        postId={data.PostID}
                                    />
                                ))}
                            {groupE &&
                                groupE.map((data) => (
                                    // console.log(data)
                                    <Event
                                        key={data.EventId}
                                        eventContent={data.Description}
                                        location={data.Location}
                                        start={data.StartTime}
                                        end={data.EndTime}
                                        attending={data.Status}
                                        eventBodyImgSrc={data.ImageUpload}
                                        name={data.Name}
                                        eventId={data.EventId}
                                        eventObj={data}
                                        setFlag={setFlag}
                                        flag={flag}
                                        setEle={setEle}
                                        setOpenModal={setOpenModal}
                                    />
                                ))}
                        </div>
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
                            <button
                                className='profileDetailBtn grey'
                                onClick={() => OpenDropdownMenu()}>
                                Add a user
                            </button>
                            <div ref={dropdown} className='dropdown-content'>
                                <div className='UserSearch'>
                                    <input
                                        ref={InputRef}
                                        className='search'
                                        type='text'
                                        onChange={({ target }) => {
                                            SearchUsers(target.value);
                                        }}
                                        placeholder='Search User'
                                    />
                                </div>
                                {displayAllNonMembers &&
                                    displayAllNonMembers.map((ele) => (
                                        <a
                                            key={ele.id}
                                            onClick={() => {
                                                RequestToS(
                                                    GetCookie(
                                                        'session_token'
                                                    ).split('&')[0],
                                                    ele.id,
                                                    socket,
                                                    'groupInvitationRequest',
                                                    groupSelect.GroupID
                                                );
                                            }}>
                                            {ele.firstName}
                                        </a>
                                    ))}
                            </div>
                        </div>
                        {(groupSelect &&
                            groupSelect.Members.reverse().map((ele, i) => (
                                <span
                                    key={ele.UserId}
                                    onClick={() =>
                                        navigate(`/profile?id=${ele.UserId}`)
                                    }>
                                    <MiniUserCard
                                        imgStyleName={'miniUserCardImg'}>
                                        <h4 style={{ color: 'black' }}>
                                            {ele.UserName}
                                        </h4>
                                        {i === 0 ? (
                                            <h4 style={{ color: 'black' }}>
                                                Admin
                                            </h4>
                                        ) : (
                                            <h4 style={{ color: 'black' }}>
                                                Member
                                            </h4>
                                        )}
                                    </MiniUserCard>
                                </span>
                            ))) || <> </>}
                    </div>
                </div>
            </Card>
        </Body>
    );
}

export default Groups;
