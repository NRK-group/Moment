import Input from '../../../components/Input/Input';
import { CloseIcon } from '../../../components/Icons/Icons';
import MiniUserCard from '../../../components/MiniUserCard/MiniUserCard';
import { useNavigate } from 'react-router-dom';
import { GetFollow } from '../hooks/getFollow';
import { useState } from 'react';
export const NewChatModal = ({ setAddToChatList }) => {
    const navigate = useNavigate();
    let following = GetFollow();
    const [query, setQuery] = useState('');
    let filteredItems = following.filter((item) => {
        let fullName = item.firstName + ' ' + item.lastName;
        return (
            item.name.toLowerCase().includes(query.toLowerCase()) ||
            fullName.toLowerCase().includes(query.toLowerCase())
        );
    });
    const handleNewMessage = (e) => {
        let receiverId = e.currentTarget.getAttribute('data-receiverid');
        fetch('http://localhost:5070/chat', {
            method: 'POST',
            headers: {
                Accept: 'application/json',
            },
            body: JSON.stringify({
                receiverId: receiverId,
            }),
            credentials: 'include',
        })
            .then(async (res) => {
                let data = await res.json();
                console.log(data);
                data ? setAddToChatList(data) : setAddToChatList([]);
                navigate(`/messages/${data.chatId}`, {
                    state: {
                        type: 'privateMessage',
                        user: data.user,
                        details: data.details,
                    },
                });
            })
            .catch((err) => {
                console.log(err);
            });
    };
    return (
        <div
            className='newChatModalContainer'
            onClick={() => {
                navigate(`/messages`);
            }}>
            <div
                className='newChatModal'
                onClick={(e) => {
                    e.stopPropagation();
                }}>
                <div className='newChatModalHeader'>
                    <div className='closeIcon'>
                        <CloseIcon
                            action={() => {
                                navigate(`/messages`);
                            }}
                        />
                    </div>
                    <div className='title'>New Message</div>
                    <div></div>
                </div>
                <div className='newChatModalHeaderSearch'>
                    <div>To:</div>
                    <Input
                        styleName='searchFollowing'
                        placeholder={'Search . . .'}
                        onChange={(e) => {
                            setQuery(e.target.value);
                        }}
                    />
                </div>
                <div className='searchResult scrollbar-hidden'>
                    {filteredItems
                        ? filteredItems.map(
                              ({ id, firstName, lastName, name, img }) => (
                                  <div
                                      key={id}
                                      data-receiverid={id}
                                      onClick={handleNewMessage}>
                                      <MiniUserCard
                                          propsId={id}
                                          name={firstName + ' ' + lastName}
                                          img={img}>
                                          {name}
                                      </MiniUserCard>
                                  </div>
                              )
                          )
                        : null}
                </div>
            </div>
        </div>
    );
};
