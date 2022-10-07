import Card from '../../components/card/Card';
import Avatar from '../../components/Avatar';
import './Post.css';
import { useRef, useState } from 'react';

export default function Post(props) {


    const dropdown = useRef(null)
    const [toggle, setToggle] = useState(true);

const OpenDropdownMenu = ()=> {
    setToggle(!toggle)
    if(toggle){
        console.log("inside")
    dropdown.current.style.display = "block"
    } else {
        dropdown.current.style.display = "none"
    }
}


    return (
        <>
            <Card styleName={'PostContainer'}>
                <Card styleName={'PostHeader'}>
                    <>
                        <Avatar
                            avatarSrc={
                                'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRaLtb_3tNc2GjjuNWX29vbxcdvMGOyGEIKaQ&usqp=CAU'
                            }
                            styleName={'PostAvatarUsers'}
                        />

                        <p>name</p>
                    </>

                    <div className='PostHeaderMenu'>
                        <button className='dropbtn' onClick={()=> OpenDropdownMenu()}>
                            <i className='fa-solid fa-ellipsis' />
                        </button>
                        <div ref={dropdown} className='dropdown-content'>
                            <a href='#'>option 1</a>
                            <a href='#'>option 2</a>
                            <a href='#'>option 3</a>
                        </div>
                    </div>
                </Card>
                <Card styleName={'PostBody'}>
                    <p>fdrfrfre</p>
                </Card>
                <Card styleName={'PostContent'}>
                    <p>fdrfrfre</p>
                </Card>
            </Card>
        </>
    );
}
