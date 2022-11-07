import './Comments.css';
import Avatar from '../../components/Avatar';
import MiniUserCard from '../../components/MiniUserCard/MiniUserCard';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import { FaceSmileIcon } from '../../components/Icons/Icons';
import { useRef, useState, useEffect } from 'react';
import ReadMoreReact from 'read-more-react';
import { useLocation } from 'react-router-dom';
import InputEmoji from 'react-input-emoji';

const Comments = ({ bodyStyleName, cardStyleName }) => {
    const { state } = useLocation();

    const [test, setTest] = useState(''); 

    useEffect(() => {
        console.log({ state });
        console.log(state.Comments);
        window.document
            .querySelectorAll('.CommentsSectionUsers .miniUserCard .contentSep')
            .forEach((ele) => ele.remove());
       
    }, []);

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

    const [text, setText] = useState('');

    function handleOnEnter(text) {
        console.log('enter', text);
    }

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

                            <p>Comm</p>
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
                            {state.Comments &&
                                state.Comments.map((ele) => (
                                    <MiniUserCard
                                        img={
                                            'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRaLtb_3tNc2GjjuNWX29vbxcdvMGOyGEIKaQ&usqp=CAU'
                                        }
                                        imgStyleName={'miniUserCardImg'}
                                        optContent={
                                            <>
                                                <h3>Name:</h3>
                                                <div className=''>
                                                    <ReadMoreReact
                                                        text={ele.Content}
                                                        readMoreText={
                                                            '...read More'
                                                        }
                                                        ideal={100}
                                                    />
                                                </div>
                                            </>
                                        }>
                                        {' '}
                                        {ele.CreatedAt &&
                                            Date(ele.CreatedAt)
                                                .slice(0, -34)
                                                .slice(0, 15) +
                                                ' -' +
                                                Date(ele.CreatedAt)
                                                    .slice(0, -34)
                                                    .slice(15)}
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
