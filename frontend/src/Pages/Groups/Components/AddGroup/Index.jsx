import './AddGroup.css';
import Card from '../../../../Components/Card/Card';
import { MessagesIcon } from '../../../../Components/Icons/Icons';
import { useState } from 'react';
import GetFollowers from '../../../Profile/Followers';
import { useEffect, useRef } from 'react';
import { GetCookie } from '../../../Profile/ProfileData';
import { RequestToS } from '../../Hooks/UseGroupshook';
import config from '../../../../../config';

export default function AddGroup({ setOpenModal, flag, setFlag, socket }) {
    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    const [closeF, setCloseF] = useState([]);
    let selectR = useRef(null);

    useEffect(() => {
        GetFollowers().then((response) => {
            setCloseF(response);
        });
    }, []);

    const CreateGroup = async () => {
        let creategroup = await fetch(config.api + '/group', {
            credentials: 'include',
            method: 'POST',
            body: JSON.stringify({ Name: name, Description: description }),
        })
            .then(async (resp) => await resp.text())
            .then((data) => data);
        if (selectR.current.value !== '' && creategroup !== '') {
            console.log('khgiuh');
            RequestToS(
                GetCookie('session_token').split('&')[0],
                selectR.current.value,
                socket,
                'groupInvitationRequest',
                creategroup
            );
        }
        setOpenModal(false);
        setFlag(!flag);
    };

    return (
        <div id='AddGroup' className='AddGroupContainer'>
            <Card styleName='newPostBox'>
                <Card styleName='newPostHeader'>
                    <span className='newPostTitle'>Create a Group</span>
                </Card>

                <Card styleName='NewPostContent'>
                    <Card styleName='NewPostContent'>
                        <br />
                        <br />
                        <br />

                        <div className='NewPostContentInput'>
                            <label htmlFor='InputName'>Group name: </label>
                            <input
                                id='InputName'
                                value={name}
                                onChange={(e) => setName(e.target.value)}
                            />
                        </div>
                        <br />
                        <div className='selectCF'>
                            <label htmlFor='selectCF'>Followers: </label>
                            <select ref={selectR}>
                                <option key={'1'} value={''}>
                                    {'Pick a followers'}
                                </option>
                                {closeF &&
                                    closeF.map((ele) => (
                                        <option key={ele.id} value={ele.id}>
                                            {ele.name}
                                        </option>
                                    ))}
                            </select>
                        </div>

                        <textarea
                            value={description}
                            onChange={(e) => setDescription(e.target.value)}
                            cols='100'
                            rows='7'
                            wrap='hard'
                            className='newPostTextContent'
                            maxLength='280'
                            placeholder={`What it's about ?`}
                        />
                        <button
                            className='NewPostSendBtn'
                            onClick={CreateGroup}>
                            <span className='shareText'>Create</span>
                            <MessagesIcon />
                        </button>
                    </Card>
                </Card>
            </Card>
        </div>
    );
}
