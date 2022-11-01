import './Comments.css';
import Avatar from '../../components/Avatar';
import MiniUserCard from '../../components/MiniUserCard/MiniUserCard';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import { FaceSmileIcon } from '../../components/Icons/Icons';
import { useRef, useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';

const Comments = ({ bodyStyleName, cardStyleName }) => {


    const { state } = useLocation();
    useEffect(() => {
        console.log({ state });
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

    return (
        <Body styleName={bodyStyleName}>
            <Card styleName={cardStyleName}>
                <div className='Comments'>
                    <div className='CommentsLeft'>
                        <div className='CommentsImg'>
                            <img src={state.postBodyImgSrc} />
                        </div>
                        <div className='CommentsChat'>
                            <div className='CommentsChatIcons'>
                                <FaceSmileIcon />
                                <input></input>
                            </div>
                            <p>Post</p>
                        </div>
                    </div>

                    <div className='CommentsSection'>
                        <Card styleName={'PostHeader'}>
                            <div style={{ display: 'flex' }}>
                                <Avatar
                                    avatarSrc={
                                        'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRaLtb_3tNc2GjjuNWX29vbxcdvMGOyGEIKaQ&usqp=CAU'
                                    }
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
                            {[
                                1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 64, 5, 3, 4, 3, 5,
                                56, 54, 34, 43,
                            ].map((ele) => (
                                <MiniUserCard
                                    img={
                                        'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRaLtb_3tNc2GjjuNWX29vbxcdvMGOyGEIKaQ&usqp=CAU'
                                    }
                                    imgStyleName='miniUserCardImg'
                                    name={'vjkjkbkj'}
                                    content={'gjkgkgk'}
                                />
                            ))}
                        </div>
                    </div>
                </div>
            </Card>
        </Body>
    );
};

export default Comments;
