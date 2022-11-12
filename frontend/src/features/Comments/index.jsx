import './Comments.css';
import Avatar from '../../components/Avatar';
import MiniUserCard from '../../components/MiniUserCard/MiniUserCard';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import { MessagesIcon, UploadIcon } from '../../components/Icons/Icons';
import { useRef, useState, useEffect } from 'react';
import ReadMoreReact from 'read-more-react';
import { useLocation } from 'react-router-dom';
import InputEmoji from 'react-input-emoji';

const Comments = ({ bodyStyleName, cardStyleName }) => {
    const { state } = useLocation();

    const [text, setText] = useState('');
    const [commentS, setCommentS] = useState([]);
    const [flag, setFlag] = useState('');

    useEffect(() => {
        console.log({ state });
       
        window.document
            .querySelectorAll('.CommentsSectionUsers .miniUserCard .contentSep')
            .forEach((ele) => ele.remove());
          
    }, [commentS]);

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



    const PostComments = async ()=> {

        let comments = await fetch(`http://localhost:5070/comment/`, {
            credentials: 'include',
            method: 'POST',
            body: JSON.stringify({postId:state.PostId, content:text})
        }).then(async (response) => {
            let resp = await response.json();
            console.log(resp)
            setText("")
            setCommentS(resp)
           return resp
        })
        setFlag(1)
    }

    function handleOnEnter(text) {
        console.log('enter', text);
    }

    useEffect(() => {
        let comments = fetch(`http://localhost:5070/comment/${state.PostId}`, {
            credentials: 'include',
            method: 'GET',
        }).then(async (response) => {
            let resp = await response.json();
            setCommentS(resp)
           return resp
        })

        window.document
        .querySelectorAll('.CommentsSectionUsers .miniUserCard .contentSep')
        .forEach((ele) => ele.remove());

    }, [flag]);

    return (
        <Body styleName={bodyStyleName}>
            <Card styleName={cardStyleName}>
                <div className='Comments'>
                    <div className='CommentsLeft'>
                        <div className='CommentsImg'>
                            <img src={state.PostBodyImgSrc} />
                        </div>
                        <div className='Comments-Emoji'></div>
                        <div className='CommentsChat'>
                            <InputEmoji
                                value={text}
                                onChange={setText}
                                cleanOnEnter
                                onEnter={handleOnEnter}
                                placeholder='Type a message'
                            />
                            <div className='CommentsIcon'>
                                <i className='fa-solid fa-upload'></i>
                                <i onClick={PostComments}>
                                    {' '}
                                    <MessagesIcon />
                                </i>
                            </div>
                        </div>
                    </div>

                    <div className='CommentsSection'>
                        <Card styleName={'PostHeader'}>
                            <div style={{ display: 'flex' }}>
                                <Avatar
                                    avatarSrc={state.AvatarSrc}
                                    styleName={'PostAvatarUsers'}
                                />

                                <p style={{ marginLeft: '4px' }}>jhfhjfjh</p>
                            </div>

                            <div className='PostHeaderMenu'>
                                <button
                                    className='dropbtn'
                                    onClick={() => OpenDropdownMenu()}>
                                    <i className='fa-solid fa-ellipsis' />
                                </button>
                                <div
                                    ref={dropdown}
                                    className='dropdown-content'>
                                    <a href='#'>option 1</a>
                                    <a href='#'>option 2</a>
                                    <a href='#'>option 3</a>
                                </div>
                            </div>
                        </Card>
                        <div className='CommentsSectionUsers'>
                            {commentS &&
                                commentS.map((ele) => (
                                    <MiniUserCard
                                    key={ele.CommentID}
                                        img={
                                            'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRaLtb_3tNc2GjjuNWX29vbxcdvMGOyGEIKaQ&usqp=CAU'
                                        }
                                        imgStyleName={'miniUserCardImg'}
                                        optContent={
                                            <>
                                                <h3>Name:</h3>
                                                <div className=''>
                                                    <ReadMoreReact
                                                        text={
                                                            ele.content
                                                        }
                                                        readMoreText={
                                                            '...read More'
                                                        }
                                                        min={40}
                                                        ideal={80}
                                                        max={150}
                                                    />
                                                   
                                                </div>
                                            </>
                                        }>
                                        {' '}
                                        {ele.CreatedAt}
                                    </MiniUserCard>
                                ))}
                        </div>
                    </div>
                </div>
            </Card>
        </Body>
    );
};

export default Comments;
