import './AddGroup.css';
import Card from '../../../../components/card/Card';
import { MessagesIcon } from '../../../../components/Icons/Icons';
import { useState } from 'react';
import GetCloseFriends from '../../../profile/CloseFriend';
import { useEffect } from 'react';

export default function AddGroup({setOpenModal,  flag , setFlag}) {
    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    const [closeF, setCloseF] = useState([]);


    useEffect(() => {
        GetCloseFriends().then((response) => {
            setCloseF(response);
        });
    }, []);

    const CreateGroup = async () => {
        let creategroup = await fetch('http://localhost:5070/group', {
            credentials: 'include',
            method: 'POST',
            body: JSON.stringify({ Name: name, Description: description }),
        })
            .then(async (resp) => await resp.text())
            .then((data) => data);
            setOpenModal(false)
            setFlag(!flag)
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
                        <label htmlFor='selectCF'>Close Friends: </label>
                        <select>
                            {closeF && closeF.map((ele) => (
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
